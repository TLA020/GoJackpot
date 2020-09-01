package main

import (
	"encoding/json"
	"github.com/gofiber/websocket"
	"github.com/mitchellh/mapstructure"
	m "goprac/models"
	"log"
	"sync"
)

var connectionsMutex = &sync.Mutex{}
var connections = make([]*m.Connection, 0)

var broadcastChan = make(chan m.Message)

func wsHandler(c *websocket.Conn) {
	log.Printf("wsHandler")
	connectionsMutex.Lock()

	var newConnection = &m.Connection{
		Conn:     c,
	}

	connections = append(connections, newConnection)
	connectionsMutex.Unlock()
	_ = newConnection.Conn.WriteMessage(websocket.TextMessage, []byte("welcome"))
	listener(newConnection)
}

func listener(conn *m.Connection) {
	for {
		_, message, err := conn.Conn.ReadMessage()
		if err != nil {
		//	close(conn.ReadChan)
			log.Printf("Error: %s", err)
			// user disconnected or something else went wrong, delete from connections.
			for i, c := range connections {
				if c == conn {
					connectionsMutex.Lock()
					connections = append(connections[:i], connections[i+1:]...)
					connectionsMutex.Unlock()
					return
				}
			}
		}

		msg := m.Message{}
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Print(err)
			continue
		}

		log.Printf("Incoming websocket message: %s", msg.Name)

		switch msg.Name {
		case "auth":
			onAuthorizeWsClient(msg, conn)
		default:
			log.Printf("default")
		}


	}
}

func onAuthorizeWsClient(msg m.Message, conn *m.Connection) {
	acc := &m.Account{}
	log.Print("Websocket client wants to authorize by token")

	if err := mapstructure.Decode(msg.Data, &acc); err != nil {
		log.Print(err)
		return
	}

	 claims, err := m.ValidateToken(acc.Token)
	 if err != nil {
		    log.Print(err)
	 }

	 conn.UserId = int(claims["sub"].(float64))
	 log.Printf("Connection now belongs to userId: %v", conn.UserId)
}

func handleBroadcasts() {
	for msg := range broadcastChan {
		// broadcast event to all connections
		for _, c := range connections {
			err := c.SendMessage(msg)
			if err != nil {
				log.Print(err)
			}
		}
	}
}

func SendBroadcast(msg m.Message) {
	broadcastChan <- msg
}