package ws

import (
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type Hub struct {
	Rooms          map[string]*Room
	Register       chan *Client
	Unregister     chan *Client
	Unicast        chan *Message
	PendingMessage *Message // ios reconnect 문제
}

type Room struct {
	Hash    string             `json:"hash"`
	Clients map[string]*Client `json:"clients"`
	Expiry  time.Time          `json:"expiry"`
}

type Client struct {
	UUID    string `json:"uuid"`
	Conn    *websocket.Conn
	Hash    string `json:"hash"`
	Message chan *Message

	closeMutex sync.Mutex // 채널 상태를 동기화하기 위한 뮤텍스
	closed     bool       // 채널이 닫혔는지 여부
}

type Topic string

var (
	JwtMissing         Topic = "JwtMissing"
	JwtInvalid         Topic = "JwtInvalid"
	DAppConnected      Topic = "DAppConnected"
	GmWalletConnected  Topic = "GmWalletConnected"
	Connect            Topic = "connect"
	EthSendTransaction Topic = "eth_sendTransaction"
	PersonalSign       Topic = "personal_sign"
	EthSign            Topic = "eth_sign"
	EthSignTypedDataV4 Topic = "eth_signTypedDataV4"
	RoomClosed         Topic = "RoomClosed"
)

type Message struct {
	Hash  string `json:"hash,omitempty"`
	Topic Topic  `json:"topic"`
	Code  int    `json:"code"`
	Data  string `json:"data,omitempty"`
}
