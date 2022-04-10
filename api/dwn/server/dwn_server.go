package server

import (
	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/dwn"
	"github.com/getzion/relay/api/dwn/errors"
	"github.com/getzion/relay/api/dwn/handler"
	"github.com/getzion/relay/api/dwn/handler/collections"
	"github.com/getzion/relay/api/models"
	"github.com/sirupsen/logrus"

	. "github.com/getzion/relay/utils"
	"github.com/gofiber/fiber/v2"
)

type interfaceMethodHandler func(handler *handler.RequestContext) ([]string, *errors.MessageLevelError)

type DWNServer struct {
	app                      *fiber.App
	modelManager             *models.ModelManager
	storage                  api.Storage
	validHubInterfaceMethods map[string]interfaceMethodHandler
}

var validHubInterfaceMethods = map[string]interfaceMethodHandler{
	constants.COLLECTIONS_QUERY: collections.CollectionsQuery,
	constants.COLLECTIONS_WRITE: collections.CollectionsWrite,
}

func InitDWNServer(modelManager *models.ModelManager, storage api.Storage) *DWNServer {
	dwnServer := &DWNServer{
		modelManager:             modelManager,
		storage:                  storage,
		validHubInterfaceMethods: validHubInterfaceMethods,
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

	var method interfaceMethodHandler
	var ok bool

	for _, message := range request.Messages {
		reply := &dwn.Reply{
			Status: &dwn.Status{},
		}

		context := handler.RequestContext{
			ModelManager: dwnServer.modelManager,
			Message:      message,
		}

		pubKey, _ := context.GetPublicKey()
		verified, _ := context.VerifyRequest(pubKey)

		reply.MessageId = "Placeholder1"

		if message.Descriptor == nil || message.Descriptor.Method == "" {
			reply.Status.Code = 400
			reply.Status.Message = errors.ImproperlyConstructedErrorMessage
			response.Replies = append(response.Replies, reply)
			logrus.Info("request message descriptor or method cannot be null or empty")
			continue
		} else if method, ok = validHubInterfaceMethods[message.Descriptor.Method]; !ok {
			reply.Status.Code = 501
			reply.Status.Message = errors.InterfaceNotImplementedErrorMessage
			response.Replies = append(response.Replies, reply)
			logrus.Infof("interface method is not implemented: %s", message.Descriptor.Method)
			continue
		}

		entry, mErr := method(&context)
		if mErr != nil {
			reply.Status.Code = mErr.Code
			reply.Status.Message = mErr.Message
			response.Replies = append(response.Replies, reply)
			logrus.Infof("identity hub method execution failed: %v", mErr)
			continue
		}

		if entry != nil {
			reply.Entries = entry
		}
		Log.Info().Bool("verified", verified).Msg("Processed valid request")
		reply.Status.Code = 200
		reply.Status.Message = errors.MessageSuccessfulMessage
		response.Replies = append(response.Replies, reply)
	}

	ctx.Response().Header.Set("Content-Type", "application/json")
	return ctx.Status(fiber.StatusOK).JSON(response)
}
