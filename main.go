package main

import (
	"github.com/getzion/relay/api/dwn/server"
	"github.com/getzion/relay/api/models"
	"github.com/getzion/relay/api/storage"
	"github.com/sirupsen/logrus"
)

func main() {
	storage, err := storage.NewStorage("mysql")
	if err != nil {
		logrus.Panic(err)
	}

	modelManager := models.NewModelManager(storage)
	server := server.InitDWNServer(modelManager, storage)

	// app := fiber.New(fiber.Config{})
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Relay is live!")
	// })
	// app.Listen(":8080")

	logrus.Fatal(server.Listen(":8080"))
}
