package models

import (
	"github.com/gofiber/websocket"
	"sync"
)

type Client struct {
	Conn   *websocket.Conn `json:"conn,omitempty"`
	UserId int
	Username string
	Avatar string
	Email  string
	Mutex sync.Mutex
}

type Event interface {
	GetType() string
}

func (c *Client) SendMessage(event Event) error {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	msg := map[string]interface{}{
		"type": event.GetType(),
		"data": event,
	}

	return c.Conn.WriteJSON(msg)
}

type IncomingMessage struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type WsAuthRequest struct {
	Name string  `json:"name"`
	Data Account `json:"data"`
}
