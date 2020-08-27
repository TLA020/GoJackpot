package controllers

import (
	"github.com/gofiber/fiber"
	"goprac/models"
	"log"
)

var CreateAccount = func(c *fiber.Ctx) {
	account := &models.Account{}

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

var Authenticate = func(c *fiber.Ctx) {
	account := &models.Account{}

	if err := c.BodyParser(account); err != nil {
		log.Fatal(err)
	}

	token, err := models.Login(account.Email, account.Password)
	if err != nil {
		c.Status(401).Send(err)
		return
	}

	c.Send(token)
}
