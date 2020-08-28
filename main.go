package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/template/html"
	"goprac/controllers"
	"goprac/services"
	"log"
	"os"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(&fiber.Settings{
		Views: engine,
	})

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
		_ = c.Render("hello-world", fiber.Map{})
	})

	app.Post("/api/v1/account/register", controllers.CreateAccount)
	app.Post("/api/v1/account/login", controllers.Authenticate)
}
