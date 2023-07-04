package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prolifel/kaching/config"
)

func main() {
	app := config.New()

	config.Catch(app.InitProgresql())

	appNew := fiber.New()

	appNew.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	appNew.Listen(":3000")
}
