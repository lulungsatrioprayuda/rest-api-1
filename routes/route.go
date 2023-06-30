package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lulungsatrioprayuda/rest-api-1/controllers/usercontroller"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	user := api.Group("/user")
	user.Get("/", usercontroller.Index)
}
