package ws

import (
	"time"
)

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Unicast:    make(chan *Message, 5),
	}
}

func (h *Hub) Run() {
	go h.DestroyExpiredRooms()

	for {
		select {
		case cl := <-h.Register:
			if _, ok := h.Rooms[cl.Hash]; ok {
				r := h.Rooms[cl.Hash]

				if _, ok := r.Clients[cl.UUID]; !ok {
					r.Clients[cl.UUID] = cl
				}
			}
		case cl := <-h.Unregister:
			if room, ok := h.Rooms[cl.Hash]; ok {
				if cl.UUID != "GM_WALLET" {
					if _, ok := room.Clients[cl.UUID]; ok {
						delete(room.Clients, cl.UUID)
					}
					if len(room.Clients) == 0 {
						delete(h.Rooms, cl.Hash)
					}
				}

			}
		case m := <-h.Unicast:
			if _, ok := h.Rooms[m.Hash]; ok {
				for _, cl := range h.Rooms[m.Hash].Clients {
					//cl.Message <- m
					cl.SendMessage(m)
				}
			}
		}
	}
}

func (h *Hub) DestroyExpiredRooms() {
	ticker := time.Tick(5 * time.Minute) // 5분마다 ticker 생성

	for range ticker {
		now := time.Now()

		// 만약 디비 연결 시에는 이 부분을 mutex로 Lock해야함 하지만 없으니 패스
		for hash, room := range h.Rooms {
			if room.Expiry.Before(now) {
				delete(h.Rooms, hash)
				for _, client := range room.Clients {
					h.Unregister <- client
				}
			}
		}
	}
}
