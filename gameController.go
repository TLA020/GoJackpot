package main

import (
	"github.com/gofiber/fiber"
	m "goprac/models"
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

	//time.Sleep(time.Second * 5)
	//fakeBetByUserId(player2Id)
	//
	//time.Sleep(time.Second * 10)
	//fakeBetByUserId(player1Id)
}

// faking users/bets to speed up test process.
func fakeBetByUserId(userId int) {
	fakeBet := Bet{Amount: (rand.Float64() * 20) + 5}
	fakeWsConn := &m.Connection{Conn: nil, UserId: userId}
	fakeGambler := &m.Gambler{Conn: fakeWsConn}

	gameManager.GetCurrentGame().PlaceBet(fakeGambler, fakeBet)
}