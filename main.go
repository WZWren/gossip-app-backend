package main

import (
	"github.com/WZWren/gossip-app-backend/database"
	"github.com/WZWren/gossip-app-backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	// this is needed as our backend and frontend runs on diff ports
	// AllowCred allows HTTP only cookies to be sent.
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
