package handler

import (
	"encoding/json"
	"fmt"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/validator"
)

type UserHandler struct {
	storage api.Storage
}

func InitUserHandler(storage api.Storage) *UserHandler {
	return &UserHandler{
		storage: storage,
	}
}

func (h *UserHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {
	case constants.COLLECTIONS_QUERY:
		return h.storage.GetUsers()

	case constants.COLLECTIONS_WRITE:
		var user api.User
		err := json.Unmarshal(data, &user)
		if err != nil {
			return nil, err
		}

		err = validator.Struct(&user)
		if err != nil {
			return nil, err
		}

		err = h.storage.InsertUser(&user)
		if err != nil {
			return nil, err
		}

		return user, nil

	default:
		return nil, fmt.Errorf("unimplemented user method: %s", method)
	}
}
