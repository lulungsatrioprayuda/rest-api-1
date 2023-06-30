package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lulungsatrioprayuda/rest-api-1/controllers/usercontroller"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	user := api.Group("/users")
	user.Get("/", usercontroller.Index)
	user.Post("/", usercontroller.CreateUser)
	user.Get("/:id", usercontroller.ShowUser)
	user.Put("/:id", usercontroller.UpdateUser)
	user.Delete("/:id", usercontroller.DeleteUser)
}
