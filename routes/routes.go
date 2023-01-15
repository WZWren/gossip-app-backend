package routes

import (
	"github.com/WZWren/gossip-app-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Healthcheck)
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
}
