package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"goprac/controllers"
	"goprac/services"
	"log"
	"os"
)

func main() {
	app := fiber.New()
	app.Use(services.JwtAuthentication)
	app.Use(middleware.Logger())
	setupRoutes(app)

	listenPort := os.Getenv("HTTP_LISTEN_PORT")
	if len(listenPort) < 1 {
		listenPort = "5001"
	}

	log.Printf("Starting HTTP-server on port %s", listenPort)

	err := app.Listen(listenPort)
	if err != nil {
		log.Fatal(err)
	}
}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) {
		msg := fmt.Sprintf("Hello 👋!")
		c.Send(msg)
	})
	app.Post("/api/v1/account/register", controllers.CreateAccount)
	app.Post("/api/v1/account/login", controllers.Authenticate)
}
