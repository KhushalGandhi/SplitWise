package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"splitwise/db"
	"splitwise/migrations"
	"splitwise/routes"
)

//"splitwise/handlers/migrations"

func main() {
	// Initialize Database
	database.Connect()
	migrations.Migrate(database.DB)

	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
