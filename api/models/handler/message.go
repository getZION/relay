package handler

import (
	"encoding/json"
	"fmt"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/validator"
)

type MessageHandler struct {
	storage api.Storage
}

func InitMessageHandler(storage api.Storage) *MessageHandler {
	return &MessageHandler{
		storage: storage,
	}
}

func (h *MessageHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {
	case constants.COLLECTIONS_QUERY:
		return h.storage.GetMessages()

	case constants.COLLECTIONS_WRITE:
		var message api.Message
		err := json.Unmarshal(data, &message)
		if err != nil {
			return nil, err
		}

		err = validator.Struct(&message)
		if err != nil {
			return nil, err
		}

		err = h.storage.InsertMessage(&message)
		if err != nil {
			return nil, err
		}

		return message, nil

	default:
		return nil, fmt.Errorf("unimplemented message method: %s", method)
	}
}
