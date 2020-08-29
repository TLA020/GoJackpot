package main

import (
	"encoding/json"
	"github.com/gofiber/websocket"
	m "goprac/models"
	"log"
	"sync"
)

var connectionsMutex = &sync.Mutex{}
var connections = make([]*m.Connection, 0)

var broadcastChan = make(chan m.Message)

func wsHandler(c *websocket.Conn) {
	connectionsMutex.Lock()

	var newConnection = &m.Connection{
		Conn:     c,
		ReadChan: make(chan m.Message),
	}

	connections = append(connections, newConnection)
	connectionsMutex.Unlock()

	listener(newConnection)
}

func listener(conn *m.Connection) {
	for {
		_, message, err := conn.Conn.ReadMessage()
		if err != nil {
			// Close readchan
			close(conn.ReadChan)

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

		log.Printf("..::SOCKETS:: Message: %s", message)
	}
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
