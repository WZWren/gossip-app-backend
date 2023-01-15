package routes

import (
	"github.com/WZWren/gossip-app-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Healthcheck)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Get("/api/getthreads", controllers.GetThreads)
	app.Post("/api/postthread", controllers.PostThread)

	app.Post("/api/getcomments", controllers.GetComments)
	app.Post("/api/postcomment", controllers.PostComment)
	app.Delete("/api/deletecomment", controllers.DeleteComment)
	app.Patch("/api/updatecomment", controllers.UpdateComment)
}
