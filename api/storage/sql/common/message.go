package common

import (
	"time"

	"github.com/getzion/relay/api"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (c *Connection) GetMessages() ([]api.Message, error) {
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

func (c *Connection) InsertMessage(message *api.Message) error {
	currentTime := time.Now().Unix()
	message.Zid = uuid.NewString()
	message.Created = currentTime
	message.Updated = currentTime

	result := c.db.Create(message)
	if result.Error != nil {
		return result.Error
	}

	logrus.Info("Message inserted with zid %s", message.Zid)
	return nil
}
