package collections

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/gabriel-vasile/mimetype"
	"github.com/getzion/relay/api/dwn/errors"
	"github.com/getzion/relay/api/dwn/handler"
	"github.com/google/uuid"
)

func CollectionsQuery(context *handler.RequestContext) ([]string, *errors.MessageLevelError) {

	var err error

	if _, err = uuid.Parse(context.Message.Descriptor.ObjectId); err != nil {
		return nil, errors.NewMessageLevelError(400, fmt.Sprintf("invalid objectId: %s", context.Message.Descriptor.ObjectId), err)
	}

	if context.Message.Descriptor.DataFormat != "" {
		if dataFormat := mimetype.Lookup(context.Message.Descriptor.DataFormat); dataFormat == nil {
			err = fmt.Errorf("invalid dataFormat: %s", context.Message.Descriptor.DataFormat)
			return nil, errors.NewMessageLevelError(400, err.Error(), err)
		}
	}

	if context.Message.Descriptor.DateSort != "" && (context.Message.Descriptor.DateSort != "createdAscending" && context.Message.Descriptor.DateSort != "createdDescending" &&
		context.Message.Descriptor.DateSort != "publishedAscending" && context.Message.Descriptor.DateSort != "publishedDescending") {
		err = fmt.Errorf("invalid dateSort: %s", context.Message.Descriptor.DateSort)
		return nil, errors.NewMessageLevelError(400, err.Error(), err)
	}

	//todo: check data & dataFormat only for application/json or do we need provide other formats?

	messageHandler, err := context.ModelManager.GetModelHandler(context.Message.Descriptor.Schema)
	if err != nil {
		return nil, errors.NewMessageLevelError(400, err.Error(), err)
	}

	data, err := messageHandler.Execute([]byte(context.Message.Data), context.Message.Descriptor.Method)
	if err != nil {
		return nil, errors.NewMessageLevelError(500, err.Error(), err)
	}

	var entries []string
	v := reflect.ValueOf(data)
	for i := 0; i < v.Len(); i++ {
		val := v.Index(i).Interface()
		result, err := json.Marshal(&val)
		if err != nil {
			return nil, errors.NewMessageLevelError(500, err.Error(), err)
		}
		entries = append(entries, string(result))
	}

	return entries, nil
}
