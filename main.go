package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/websocket"
	"log"
	"os"
)

func main() {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
		}
		c.Next()
	})
	app.Use(middleware.Logger())

	setupRoutes(app)
	go handleBroadcasts()

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
	app.Static("/", "./frontend/dist/index.html")
	app.Static("/static", "./frontend/dist/static")

	app.Get("/ws/", websocket.New(func(c *websocket.Conn) {
		wsHandler(c)
	}))

	app.Use(JwtAuthentication)
	app.Post("/api/v1/account/register", createAccount)
	app.Post("/api/v1/account/login", authenticate)
}
