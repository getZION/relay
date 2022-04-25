package models

import (
	"fmt"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/models/handler"
)

type ModelManager struct {
	handlers map[string]ModelHandler
}

type ModelHandler interface {
	Execute(data []byte, method string) (interface{}, error)
}

func NewModelManager(storage api.Storage) *ModelManager {
	methodHandler := &ModelManager{
		handlers: map[string]ModelHandler{
			constants.ZION_COMMUNITY:       handler.InitCommunityHandler(storage),
			constants.ZION_CONVERSATION:    handler.InitConversationHandler(storage),
			constants.ZION_MESSAGE:         handler.InitMessageHandler(storage),
			constants.ZION_JOIN_COMMUNITY:  handler.InitCommunityJoinHandler(storage),
			constants.ZION_LEAVE_COMMUNITY: handler.InitCommunityLeaveHandler(storage),
			constants.ZION_USER:            handler.InitUserHandler(storage),

			// constants.PAYMENT:                  handler.InitPaymentHandler(storage),
			// constants.PERSON:

			// constants.ZION_COMMENT:             handler.InitCommentHandler(storage),
			// constants.ZION_COMMUNITY_KICK_USER: handler.InitCommunityKickUserHandler(storage),
		},
	}

	return methodHandler
}

func (sm *ModelManager) GetModelHandler(model string) (ModelHandler, error) {
	if handler, ok := sm.handlers[model]; ok {
		return handler, nil
	}

	return nil, fmt.Errorf("unknown model: %s", model)
}
