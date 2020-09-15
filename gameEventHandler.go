package main

import (
	m "goprac/models"
)


type GameEvent struct {
	Type string
	Game Game
	Player *Player
	Amount float64
}

type CurrentUsersEvent struct {
	Clients []*m.Client
}

type CountDownEvent struct {
	TimeLeft float64
}

func handleGameEvents() {
	for event := range gameManager.Events() {
		switch e := event.(type) {
		case GameEvent:
			genericGameEventHandler(e)
		case CurrentUsersEvent:
			currentUsersHandler(e.Clients)
		case CountDownEvent:
			countDownHandler(e.TimeLeft)
		}
	}
}

func genericGameEventHandler(event GameEvent) {
	SendBroadcast(m.NewMessage(event.Type, map[string]interface{} {
		"type": event.Type,
		"game": event.Game,
		"player": event.Player,
		"amount": event.Amount,
	}))
}

func currentUsersHandler(clients []*m.Client) {
	SendBroadcast(m.NewMessage("current-users", map[string]interface{}{
		"users": clients,
	}))
}

func countDownHandler(timeLeft float64) {
	SendBroadcast(m.NewMessage("time-left", map[string]interface{}{
		"timeLeft": timeLeft,
	}))
}
