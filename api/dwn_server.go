package api

import (
	"github.com/getzion/relay/api/dwn"
	"github.com/gofiber/fiber/v2"
)

type DWNServer struct {
	app *fiber.App
}

func InitDWNServer() *DWNServer {
	dwnServer := &DWNServer{}
	app := fiber.New(fiber.Config{})
	app.Post("/", dwnServer.Process)
	dwnServer.app = app
	return dwnServer
}

func (dwnServer *DWNServer) Listen(addr string) error {
	return dwnServer.app.Listen(addr)
}

func (dwnServer *DWNServer) Process(ctx *fiber.Ctx) error {
	request := dwn.Request{}

	if err := ctx.BodyParser(&request); err != nil {
		ctx.SendStatus(fiber.StatusBadRequest)
	}

	response := &dwn.Response{
		RequestId: request.RequestId,
		Status: &dwn.Status{
			Code: 200,
		},
		Replies: []*dwn.Reply{},
	}

	ctx.Response().Header.Set("Content-Type", "application/json")
	return ctx.Status(fiber.StatusOK).JSON(response)
}
