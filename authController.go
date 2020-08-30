package main

import (
	"github.com/gofiber/fiber"
	m "goprac/models"
	"log"
)

var createAccount = func(c *fiber.Ctx) {
	account := &m.Account{}

	if err := c.BodyParser(account); err != nil {
		log.Fatal(err)
	}

	err := account.Create()
	if err != nil {
		c.Status(401).Send(err)
		return
	}

	_ = c.JSON(account)
}

var authenticate = func(c *fiber.Ctx) {
	account := &m.Account{}

	if err := c.BodyParser(account); err != nil {
		log.Fatal(err)
	}

    err := account.Login()
	if err != nil {
		c.Status(401).Send(err)
		return
	}

	_ = c.JSON(account)
}
