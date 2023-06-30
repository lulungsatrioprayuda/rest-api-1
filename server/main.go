package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/lulungsatrioprayuda/rest-api-1/database/connection"
	"github.com/lulungsatrioprayuda/rest-api-1/database/migrations"
	"github.com/lulungsatrioprayuda/rest-api-1/routes"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	connection.DatabseConnection()
	migrations.Migrate()

	routes.SetupRoutes(app)

	app.Listen(":9999")
}
