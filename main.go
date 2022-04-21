package main

import (
	. "github.com/getzion/relay/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	Log.Debug().Msg("Loading...")

	app := fiber.New(fiber.Config{})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Relay is live!")
	})
	app.Listen(":8080")
}
