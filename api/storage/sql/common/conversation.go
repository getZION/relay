package common

import (
	"time"

	"github.com/getzion/relay/api"
	"github.com/google/uuid"
)

func (c *Connection) GetConversations() ([]api.Conversation, error) {
	var conversations []api.Conversation
	result := c.db.Find(&conversations)
	if result.Error != nil {
		return nil, result.Error
	}
	return conversations, nil
}

func (c *Connection) InsertConversation(conversation *api.Conversation) error {
	currentTime := time.Now().Unix()
	conversation.Zid = uuid.NewString()
	conversation.Created = currentTime
	conversation.Updated = currentTime

	result := c.db.Create(conversation)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
