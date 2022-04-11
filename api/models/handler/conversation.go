package handler

import (
	"encoding/json"
	"fmt"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/validator"
)

type ConversationHandler struct {
	storage api.Storage
}

func InitConversationHandler(storage api.Storage) *ConversationHandler {
	return &ConversationHandler{
		storage: storage,
	}
}

func (h *ConversationHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {
	case constants.COLLECTIONS_QUERY:
		return h.storage.GetConversations()

	case constants.COLLECTIONS_WRITE:
		var conversation api.Conversation
		err := json.Unmarshal(data, &conversation)
		if err != nil {
			return nil, err
		}

		err = validator.Struct(&conversation)
		if err != nil {
			return nil, err
		}

		err = h.storage.InsertConversation(&conversation)
		if err != nil {
			return nil, err
		}

		return conversation, nil

	default:
		return nil, fmt.Errorf("unimplemented conversation method: %s", method)
	}
}
