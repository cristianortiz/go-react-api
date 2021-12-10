package main

import (
	"go-react-next/src/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello, world")
	})
	app.Listen(":8000")
}
