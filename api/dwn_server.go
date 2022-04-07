package api

import (
	"github.com/gofiber/fiber/v2"
)

type DWNServer struct {
	app *fiber.App
}

func InitDWNServer() *DWNServer {
	dwnServer := &DWNServer{}
	app := fiber.New(fiber.Config{})
	// app.Post("/", dwnServer.Process)
	dwnServer.app = app
	return dwnServer
}

func (dwnServer *DWNServer) Listen(addr string) error {
	return dwnServer.app.Listen(addr)
}
