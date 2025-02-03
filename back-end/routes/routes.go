package routes

import (
	"your-project/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User routes
	api.Get("/users", handlers.GetUsers)
	api.Post("/users", handlers.CreateUser)

	// Product routes
	api.Get("/products", handlers.GetProducts)
	api.Post("/products", handlers.CreateProduct)
}
