package handler

import (
	"encoding/json"
	"fmt"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/constants"
	"github.com/getzion/relay/api/validator"
)

type CommunityHandler struct {
	storage api.Storage
}

func InitCommunityHandler(storage api.Storage) *CommunityHandler {
	return &CommunityHandler{
		storage: storage,
	}
}

func (h *CommunityHandler) Execute(data []byte, method string) (interface{}, error) {
	switch method {

	case constants.COLLECTIONS_QUERY:
		return h.storage.GetCommunities()

	case constants.COLLECTIONS_WRITE:
		var community api.Community
		err := json.Unmarshal(data, &community)
		if err != nil {
			return nil, err
		}

		err = validator.Struct(&community)
		if err != nil {
			return nil, err
		}

		err = h.storage.InsertCommunity(&community)
		if err != nil {
			return nil, err
		}

		return community, nil

	default:
		return nil, fmt.Errorf("unimplemented organization method: %s", method)
	}
}
