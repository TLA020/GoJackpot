package main

import (
	"fmt"
	"goprac/models"
	"sync"
)

type Message struct {
	UserName string `json:"userName"`
	Email string `json:"email"`
	Msg string `json:"msg"`
	Avatar string `json:"avatar"`
}

func NewMessage(msg string, userName string, email string) *Message {
	return &Message {
		Msg: msg,
		UserName: userName,
		Email: email,
		Avatar: fmt.Sprint("https://api.adorable.io/avatars/64/%s", email),
	}
}

type Chat struct {
	Mutex       sync.Mutex
	Messages     []*Message `json:"messages"`
	incoming      chan *Message
}

func NewChat() *Chat {
	return &Chat{
		Mutex:     sync.Mutex{},
		Messages:     make([]*Message, 0),
	}
}

func (c *Chat) Incoming() chan *Message {
	return c.incoming
}

func (c *Chat) handleMessages() {
	for msg := range c.Incoming() {
		c.Messages = append(c.Messages, msg)
		msg.broadcast()
	}
}

func (m *Message) broadcast() {
	SendBroadcast(models.NewEvent("chat-message", map[string]interface{}{
		"userName": m.UserName,
		"Email": m.Email,
		"Msg": m.Msg,
		"Avatar": m.Avatar,
	}))
}




