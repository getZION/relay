package server

import (
	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/dwn"
	"github.com/getzion/relay/api/dwn/handler"
	"github.com/getzion/relay/api/models"

	. "github.com/getzion/relay/utils"
	"github.com/gofiber/fiber/v2"
)

type DWNServer struct {
	app          *fiber.App
	modelManager *models.ModelManager
	storage      api.Storage
}

func InitDWNServer(modelManager *models.ModelManager, storage api.Storage) *DWNServer {
	dwnServer := &DWNServer{
		modelManager: modelManager,
		storage:      storage,
	}
	app := fiber.New(fiber.Config{})
	app.Post("/", dwnServer.Process)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Relay is live!")
	})
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

	for _, message := range request.Messages {
		context := handler.RequestContext{
			// ModelManager: identityHub.modelManager,
			// ModelManager: dwn
			Message: message,
		}

		pubKey, _ := context.GetPublicKey()
		verified, _ := context.VerifyRequest(pubKey)

		Log.Info().Bool("verified", verified).Msg("Processed valid request")
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
