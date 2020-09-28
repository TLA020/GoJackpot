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

type Event struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type WsAuthEvent struct {
	Name string  `json:"name"`
	Data Account `json:"data"`
}

func NewEvent(event string, data map[string]interface{}) Event {
	return Event{
		Type: event,
		Data: data,
	}
}
