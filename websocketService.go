package main

import (
	"encoding/json"
	"github.com/gofiber/websocket"
	"github.com/mitchellh/mapstructure"
	m "goprac/models"
	"goprac/utils"
	"log"
	"sync"
)

var clientsMutex = &sync.Mutex{}

var clients = make([]*m.Client, 0)
var register = make(chan *m.Client)
var broadcast = make(chan m.Event)
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
	var client = &m.Client{Conn: c, Mutex: sync.Mutex{}}

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
		log.Print("incoming ws message")
		if err != nil {
			log.Println("[WS] Read error:", err)
			return // runs deferred function
		}

		handleMessages(client, message)
	}
}

func onClientsUpdate() {
	log.Print("clients update send broadcast")
	gameManager.events <- CurrentUsersEvent{
		clients,
	}
}

func handleMessages(client *m.Client, message []byte) {
	msg := m.IncomingMessage{}
	if err := json.Unmarshal(message, &msg); err != nil {
		log.Print(err)
		return
	}

	// temp move to event handler.
	switch msg.Type {
	case "auth":
		onAuthorizeWsClient(msg, client)
	case "place-bet":
		onBetPlaced(msg, client)
	case "chat-message":
		handleChatMsg(msg, client)
	default:
		log.Printf("default")
	}
}

func handleChatMsg(e m.IncomingMessage, c *m.Client) {
	msg := NewMessage(e.Data["message"].(string), c.Username, c.Avatar)
	chat.incoming <- msg
}

func onBetPlaced(msg m.IncomingMessage, conn *m.Client) {
	amount := msg.Data["amount"].(float64)
	player := NewPlayer(conn.UserId, conn.Username, conn.Avatar)
	gameManager.GetCurrentGame().PlaceBet(player, amount)
}

func onAuthorizeWsClient(msg m.IncomingMessage, client *m.Client) {
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

	if !utils.KeyExists(claims, "username") || !utils.KeyExists(claims, "avatar") {
		// temp
		gameManager.events <- GameEvent{
			Type: "invalid-user",
		}
		return
	}

	client.Username = claims["username"].(string)
	client.Avatar = claims["avatar"].(string)
	client.UserId = int(claims["sub"].(float64))
	client.Email = claims["email"].(string)

	// game snapshot
	gameManager.events <- GameEvent{
		Type: "current-game",
		Game: *gameManager.GetCurrentGame(),
	}

	// chat snapshot
	chatSnapshot := ChatSnapshot{Messages: chat.Messages}
	_ = client.SendMessage(chatSnapshot)

	log.Printf("[WS] Client now belongs to: %v", client.Username)
}

func SendBroadcast(event m.Event) {
	broadcast <- event
}
