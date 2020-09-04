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
}

type CurrentGame struct {
	Game Game
}

type CurrentUsersEvent struct {
	Clients []*m.Client
}

type CountDownEvent struct {
	TimeLeft float64
}

type WinnerPickedEvent struct {
	Player Player
	Amount float64
}

func newGameHandler(game Game) {
	SendBroadcast(m.NewMessage("new-game", map[string]interface{}{
		"game": game,
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
		"game": game,
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

func winnerPickedHandler(player Player, amount float64) {
	SendBroadcast(m.NewMessage("winner-picked", map[string]interface{}{
		"winner": player,
		"amount": amount,
	}))
}

func sendCurrentGame(game Game) {
	SendBroadcast(m.NewMessage("current-game", map[string]interface{}{
		"game": game,
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
		case CountDownEvent:
			countDownHandler(e.TimeLeft)
		case WinnerPickedEvent:
			winnerPickedHandler(e.Player, e.Amount)
		case CurrentGame:
			sendCurrentGame(e.Game)
		}
	}
}
