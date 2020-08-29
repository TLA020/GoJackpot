package models

import (
	"github.com/gofiber/websocket"
)

type Connection struct {
	Conn     *websocket.Conn
	ReadChan chan Message
}

func (c *Connection) SendMessage(msg interface{}) error {
	return c.Conn.WriteJSON(msg)
}

type Message struct {
	Name string                 `json:"name"`
	Data map[string]interface{} `json:"data"`
}

func newMessage(name string, data map[string]interface{}) Message {
	return Message{
		Name: name,
		Data: data,
	}
}
