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

	app.Get("/api/thread/get", controllers.GetThreads)
	app.Post("/api/thread/post", controllers.PostThread)

	// get comment is a post request in this case: we need the relevant
	// thread information to get it.
	app.Post("/api/comment/get", controllers.GetComments)
	app.Post("/api/comment/post", controllers.PostComment)
	app.Delete("/api/comment/delete", controllers.DeleteComment)
	app.Patch("/api/comment/update", controllers.UpdateComment)

	app.Post("api/tabs/post", controllers.PostTabs)
	app.Post("api/tabs/get", controllers.GetTabs)
}
