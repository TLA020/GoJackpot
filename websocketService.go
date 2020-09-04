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
var connections = make([]*m.Client, 0)

var broadcastChan = make(chan m.Message)

func wsHandler(c *websocket.Conn) {
	log.Printf("[WS] New connection")
	connectionsMutex.Lock()

	var newConnection = &m.Client{
		Conn: c,
	}

	connections = append(connections, newConnection)
	connectionsMutex.Unlock()
	listener(newConnection)
}

func listener(conn *m.Client) {
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
					break
				}
			}
			gameManager.events <- CurrentUsersEvent{
				connections,
			}
			return
		}

		msg := m.Message{}
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Print(err)
			continue
		}

		log.Printf("Incoming websocket message: %s", msg.Event)

		switch msg.Event {
		case "auth":
			onAuthorizeWsClient(msg, conn)
		case "place-bet":
			onBetPlaced(msg, conn)
		default:
			log.Printf("default")
		}

	}
}

func onBetPlaced(msg m.Message, conn *m.Client) {
	amount := msg.Data["amount"].(float64)
	gambler := &m.Gambler{Conn: conn}
	gameManager.GetCurrentGame().PlaceBet(gambler, Bet{Amount: amount})
}


func onAuthorizeWsClient(msg m.Message, client *m.Client) {
	acc := &m.Account{}
	log.Print("[WS] client wants to authorize by token")

	if err := mapstructure.Decode(msg.Data, &acc); err != nil {
		log.Print(err)
		return
	}

	claims, err := m.ValidateToken(acc.Token)
	if err != nil {
		log.Print(err)
	}

	client.UserId = int(claims["sub"].(float64))
	client.Email = claims["email"].(string)

	// broadcast to change "online players" in front-end
	gameManager.events <- CurrentUsersEvent{
		connections,
	}

	log.Printf("[WS] Client now belongs to userId: %v", client.UserId)
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
