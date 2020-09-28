package main

import (
	m "goprac/models"
)

// generic struct, can be used for multiple game-events, like: new-game
type GameEvent struct {
	Type    string  `json:"type"`
	Game    Game	`json:"game"`
	Player  *Player	`json:"player"`
	Amount  float64	`json:"amount"`
}

func (e GameEvent) GetType() string {
	return e.Type
}

// raises when winner picked
type WinnerPickedEvent struct {
	Player 		*Player `json:"player"`
	Amount 		float64 `json:"amount"`
	Ticket 		int 	`json:"ticket"`
	Percentage  float64 `json:"percentage"`
}

func (e WinnerPickedEvent) GetType() string {
	return "winner-picked"
}

// raised when new bet placed by user
type BetPlacedEvent struct {
	Player 		*Player	`json:"player"`
	Amount 		float64	`json:"amount"`
}

func (e BetPlacedEvent) GetType() string {
	return "bet-placed"
}

// raised when current connections / users changed
type CurrentUsersEvent struct {
	Clients []*m.Client
}

func (e CurrentUsersEvent) GetType() string {
	return "current-users"
}

// raised when game started to emit time-left.
type CountDownEvent struct {
	TimeLeft float64 `json:"timeLeft"`
}

func (e CountDownEvent) GetType() string {
	return "time-left"
}

func handleGameEvents() {
	for event := range gameManager.Events() {
		SendBroadcast(event)
	}
}


