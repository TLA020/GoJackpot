package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	jwtware "github.com/gofiber/jwt"
	"github.com/gofiber/websocket"
	"log"
	"os"
)

var stopChan = make(chan bool)
var gameManager = NewGameManager()

//var chat = NewChat()

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
		}
		c.Next()
	})

	app.Use(middleware.Logger())

	setupRoutes(app)
	go runHub()

	listenPort := os.Getenv("HTTP_LISTEN_PORT")
	if len(listenPort) < 1 {
		listenPort = "5001"
	}

	log.Printf("Starting HTTP-server on port %s", listenPort)
	go handleGameEvents()
	//go chat.handleMessages()
	gameManager.NewGame()

	err := app.Listen(listenPort)
	if err != nil {
		log.Fatal(err)
	}

	stopChan <- true
}

func setupRoutes(app *fiber.App) {
	app.Static("/", "./frontend/dist/index.html")
	app.Static("/static", "./frontend/dist/static")

	app.Get("/ws/", websocket.New(func(c *websocket.Conn) {
		wsHandler(c)
	}))

	app.Post("/api/v1/accounts/register", signUp)
	app.Post("/api/v1/accounts/login", signIn)
	app.Post("/api/v1/game/test/bets/random", testGame)
	// after this middleware only authorized routes.
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	}))

	app.Post("api/v1/accounts/avatar", uploadAvatar)

	app.Get("/restricted", func(c *fiber.Ctx) {
		// test endpoint
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		_ = c.JSON(claims)
	})
}
