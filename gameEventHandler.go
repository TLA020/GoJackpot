package main

import (
	m "goprac/models"
)

type NewGameEvent struct {
	Game Game
}

type StartGameEvent struct {
	Game Game
}

type EndGameEvent struct {
	Game Game
}

type NewBetEvent struct {
	Game Game
	Bet  Bet
}

type CurrentUsersEvent struct {
	Clients []*m.Client
}

func newGameHandler(game Game) {
	SendBroadcast(m.NewMessage("new-game", map[string]interface{}{
		"Data": game,
	}))
}

func startGameHandler(game Game) {
	SendBroadcast(m.NewMessage("start-game", map[string]interface{}{
		"Data": game,
	}))
}

func endGameHandler(game Game) {
	SendBroadcast(m.NewMessage("end-game", map[string]interface{}{
		"Data": game,
	}))
}

func betPlacedHandler(game Game) {
	SendBroadcast(m.NewMessage("bet-placed", map[string]interface{}{
		"Data": game,
	}))
}

func currentUsersHandler(clients []*m.Client) {
	SendBroadcast(m.NewMessage("current-users", map[string]interface{}{
		"users": clients,
	}))
}

func handleGameEvents() {
	for event := range gameManager.Events() {
		switch e := event.(type) {
		case NewGameEvent:
			newGameHandler(e.Game)
		case StartGameEvent:
			startGameHandler(e.Game)
		case EndGameEvent:
			endGameHandler(e.Game)
		case NewBetEvent:
			betPlacedHandler(e.Game)
		case CurrentUsersEvent:
			currentUsersHandler(e.Clients)
		}
	}
}
