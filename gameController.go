package main

import (
	"github.com/gofiber/fiber"
	"github.com/goombaio/namegenerator"
	"math/rand"
	"time"
)

var testGame = func(c *fiber.Ctx) {
	go launchTestGame()
	c.Send("[TEST] started")
}

func launchTestGame() {
	for i := 0; i < 3; i++ {
		fakeBetByUserId(rand.Int() *20 +5)
	}
}

// faking users/bets to speed up test process.
func fakeBetByUserId(userId int) {
	fakeBet := (rand.Float64() * 100) + 5
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	fakePlayer := NewPlayer(userId, nameGenerator.Generate())
	gameManager.GetCurrentGame().PlaceBet(fakePlayer, fakeBet)
}
