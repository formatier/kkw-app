package main

import (
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/healtz", func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status": "healthy",
		})
	})

	app.Listen(":3000")
}
