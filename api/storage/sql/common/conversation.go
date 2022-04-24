package common

import (
	"time"

	"github.com/getzion/relay/api"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (c *Connection) GetConversations() ([]api.Conversation, error) {
	return nil, nil
	// row := c.db.QueryRow("CALL get_conversations")

	// var conversations []api.Conversation
	// var jsonConversations string

	// if err := row.Scan(&jsonConversations); err != nil {
	// 	return nil, err
	// }

	// err := json.Unmarshal([]byte(jsonConversations), &conversations)
	// if err != nil {
	// 	return nil, err
	// }

	// return conversations, nil
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

	logrus.Info("Conversation inserted with zid %s", conversation.Zid)
	return nil
}
