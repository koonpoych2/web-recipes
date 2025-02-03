package pkg

import "github.com/gofiber/fiber/v2"

func SetupMiddlewares(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Next()
	})
}
