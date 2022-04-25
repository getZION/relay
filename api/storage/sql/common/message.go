package common

import (
	"time"

	"github.com/getzion/relay/api"
	"github.com/google/uuid"
)

func (c *Connection) GetMessages() ([]api.Message, error) {
	var messages []api.Message
	result := c.db.Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}
	return messages, nil
}

func (c *Connection) InsertMessage(message *api.Message) error {
	currentTime := time.Now().Unix()
	message.Zid = uuid.NewString()
	message.Created = currentTime
	message.Updated = currentTime

	result := c.db.Create(message)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
