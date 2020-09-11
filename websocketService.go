package main

import (
	"encoding/json"
	"github.com/gofiber/websocket"
	"github.com/mitchellh/mapstructure"
	m "goprac/models"
	"log"
	"sync"
)

var clientsMutex = &sync.Mutex{}

var clients = make([]*m.Client, 0)
var register = make(chan *m.Client)
var broadcast = make(chan m.Message)
var unregister = make(chan *m.Client)

func runHub() {
	for {
		select {
		case client := <-register:
			// Register new connection
			clientsMutex.Lock()
			clients = append(clients, client)
			clientsMutex.Unlock()
			onClientsUpdate()

		case message := <-broadcast:
			// Broadcast msg to all clients
			for _, c := range clients {
				err := c.SendMessage(message)
				if err != nil {
					unregister <- c
					log.Print(err)
				}
			}

		case connection := <-unregister:
			// Remove the client from the hub
			for i, c := range clients {
				if c == connection {
					clientsMutex.Lock()
					clients = append(clients[:i], clients[i+1:]...)
					clientsMutex.Unlock()
					onClientsUpdate()
				}
			}
		}
	}
}

func wsHandler(c *websocket.Conn) {
	log.Printf("[WS] New connection")
	var client = &m.Client{Conn: c}

	register <- client
	listener(client)
}

func listener(client *m.Client) {
	defer func() {
		unregister <- client
		if err := client.Conn.Close(); err != nil {
			log.Println("[WS] Close error:", err)
		}
	}()

	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println("[WS] Read error:", err)
			return // runs deferred function
		}

		msg := m.Message{}
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Print(err)
			continue
		}
		handleMessage(client, msg)
	}
}

func onClientsUpdate() {
	log.Print("clients update send broadcast")
	gameManager.events <- CurrentUsersEvent{
		clients,
	}
}

func handleMessage(client *m.Client, msg m.Message) {
	// temp move to event handler.
	switch msg.Event {
	case "auth":
		onAuthorizeWsClient(msg, client)
	case "place-bet":
		onBetPlaced(msg, client)
	default:
		log.Printf("default")
	}
}

func onBetPlaced(msg m.Message, conn *m.Client) {
	amount := msg.Data["amount"].(float64)
	player := NewPlayer(conn.UserId, conn.Email)
	gameManager.GetCurrentGame().PlaceBet(player, amount)
}

func onAuthorizeWsClient(msg m.Message, client *m.Client) {
	defer onClientsUpdate()

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

	gameManager.events <- CurrentGame{
		gameManager.currentGame,
	}

	log.Printf("[WS] Client now belongs to userId: %v", client.UserId)
}

func SendBroadcast(msg m.Message) {
	broadcast <- msg
}
