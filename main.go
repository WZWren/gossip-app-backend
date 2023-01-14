package main

import (
	"github.com/WZWren/gossip-app-backend/database"
	"github.com/WZWren/gossip-app-backend/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
