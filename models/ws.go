package models

import (
	"github.com/gofiber/websocket"
)

type Client struct {
	Conn   *websocket.Conn `json:"conn,omitempty"`
	UserId int
	Email  string
}

func (c *Client) SendMessage(msg interface{}) error {
	return c.Conn.WriteJSON(msg)
}

type Message struct {
	Event string                 `json:"event"`
	Data  map[string]interface{} `json:"data"`
}

type WsAuthEvent struct {
	Name string  `json:"name"`
	Data Account `json:"data"`
}

func NewMessage(event string, data map[string]interface{}) Message {
	return Message{
		Event: event,
		Data:  data,
	}
}
