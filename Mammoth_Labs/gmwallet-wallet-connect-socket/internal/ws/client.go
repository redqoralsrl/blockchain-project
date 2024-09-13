package ws

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func (c *Client) Close() {
	c.closeMutex.Lock()
	defer c.closeMutex.Unlock()

	if !c.closed {
		close(c.Message)
		c.closed = true
	}
}

func (c *Client) SendMessage(msg *Message) {
	c.closeMutex.Lock()
	defer c.closeMutex.Unlock()

	if !c.closed {
		c.Message <- msg
	}
}

func (c *Client) WriteMessage() {
	ticker := time.NewTicker(45 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case message, ok := <-c.Message:
			if !ok {
				return // 메시지 채널이 닫힘
			}
			_ = c.Conn.WriteJSON(message) // 메시지 보내기
		case <-ticker.C:
			if err := c.Conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(10*time.Second)); err != nil {
				log.Println("Write ping error:", err)
				return // 핑 보내기 실패 시 연결 종료
			}
		}
	}
}

func (c *Client) ReadMessage(hub *Hub) error {
	defer func() {
		_ = c.Conn.Close()
		c.Close()
		hub.Unregister <- c
	}()
	if hub.PendingMessage != nil && c.Hash == hub.PendingMessage.Hash {
		hub.Unicast <- hub.PendingMessage
	}

	for {
		if c.closed {
			return errors.New("connection closed")
		}
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			//if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			//	log.Printf("websocket read message error: %v", err)
			//}
			return err
		}

		if len(m) > 0 {
			var msg Message
			err = json.Unmarshal(m, &msg)
			if err != nil {
				// JSON 파싱 에러 처리
				log.Printf("error parsing message: %v", err)
				return errors.New("error parsing message")
			}

			if c.Hash == msg.Hash {
				message := &Message{
					Hash:  c.Hash,
					Topic: msg.Topic,
					Code:  200,
					Data:  msg.Data,
				}

				hub.PendingMessage = message

				hub.Unicast <- message
			}
		}
	}
}
