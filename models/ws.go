package models

import (
	"github.com/gofiber/websocket"
)

type Client struct {
	Conn   *websocket.Conn `json:"conn,omitempty"`
	UserId int
	Email  string
}

type Event interface {
	GetType() string
}

func (c *Client) SendMessage(event Event) error {
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
