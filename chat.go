package main

import (
	"sync"
)

type Message struct {
	UserName string `json:"userName"`
	Email string `json:"email"`
	Msg string `json:"msg"`
	Avatar string `json:"avatar"`
}

func (m Message) GetType() string {
	return "chat-message"
}

type ChatSnapshot struct {
	Messages []Message `json:"messages"`
}

func (c ChatSnapshot) GetType() string {
	return "chat-snapshot"
}

func NewMessage(msg string, username string, avatar string) *Message {
	return &Message {
		Msg: msg,
		UserName: username,
		Avatar: avatar,
	}
}

type Chat struct {
	Mutex         sync.Mutex
	Messages      []Message `json:"messages"`
	incoming      chan *Message
}

func NewChat() *Chat {
	return &Chat{
		Mutex:     	  sync.Mutex{},
		Messages:     make([]Message, 0),
		incoming: 	  make(chan *Message),
	}
}

func (c *Chat) Incoming() chan *Message {
	return c.incoming
}

func handleChatMessages() {
	for msg := range chat.Incoming() {
		chat.Mutex.Lock()

		if len(chat.Messages) > 8 {
			chat.Messages = chat.Messages[1:]
		}
		chat.Messages = append(chat.Messages, *msg)
		chat.Mutex.Unlock()
		msg.broadcast()
	}
}

func (m *Message) broadcast() {
	broadcast <- m
}




