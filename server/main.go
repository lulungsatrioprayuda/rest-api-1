package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lulungsatrioprayuda/rest-api-1/database/connection"
	"github.com/lulungsatrioprayuda/rest-api-1/database/migrations"
	"github.com/lulungsatrioprayuda/rest-api-1/routes"
)

func main() {

	connection.DatabseConnection()

	migrations.Migrate()

	app := fiber.New()
	routes.SetupRoutes(app)

	app.Listen(":9999")
}
