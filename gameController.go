package main

import (
	"github.com/gofiber/fiber"
	"math/rand"
)

var testGame = func(c *fiber.Ctx) {
	go launchTestGame()
	c.Send("[TEST] started")
}

func launchTestGame() {
	player1Id := 2
	player2Id := 4
	fakeBetByUserId(player1Id)
	fakeBetByUserId(player2Id)
}

// faking users/bets to speed up test process.
func fakeBetByUserId(userId int) {
	fakeBet := (rand.Float64() * 20) + 5
	fakePlayer := NewPlayer(userId, "")
	gameManager.GetCurrentGame().PlaceBet(fakePlayer, fakeBet)
}