package main

import (
	"log"

	"back-end/database"
	"back-end/routes"

	"back-end/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load config
	config.LoadEnv()

	// Initialize Fiber app
	app := fiber.New()

	// Connect to the database
	database.ConnectDB()

	// Register routes
	routes.SetupRoutes(app)

	// Start the server
	port := config.GetPort()
	log.Fatal(app.Listen(":" + port))
}
