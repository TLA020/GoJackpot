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
	for i := 0; i < 20; i++ {
		fakeBetByUserId(rand.Int() *20 +5)
	}

	gameManager.GetCurrentGame().CalculateShares()
}

// faking users/bets to speed up test process.
func fakeBetByUserId(userId int) {
	fakeBet := (rand.Float64() * 100) + 5
	fakePlayer := NewPlayer(userId, "")
	gameManager.GetCurrentGame().PlaceBet(fakePlayer, fakeBet)
}
