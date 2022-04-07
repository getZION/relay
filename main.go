package main

import (
	"log"

	"github.com/getzion/relay/dwn"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		// Require JSON
		app.Use(func(c *fiber.Ctx) error {
			if c.Is("json") {
				return c.Next()
			}
			return c.SendString("JSON required.")
		})

		// Instantiate new Request struct
		r := new(dwn.Request)
		//  Parse body into Request struct
		if err := c.BodyParser(r); err != nil {
			log.Println(err)
			c.Status(400).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
			return nil
		}

		return c.JSON(&fiber.Map{
			"success":   true,
			"requestId": r.RequestId,
			"target":    r.Target,
		})
	})

	app.Listen(":3000")
}
