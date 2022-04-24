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
	// //todo: add owner_did? check creator user for permission?
	currentTime := time.Now().Unix()
	conversation.Zid = uuid.NewString()
	conversation.Created = currentTime
	conversation.Updated = currentTime

	result := c.db.Create(conversation)
	if result.Error != nil {
		return result.Error
	}

	logrus.Info("Done with conversation insert?")
	return nil

	// tx, err := c.db.Begin()
	// if err != nil {
	// 	return err
	// }

	// result, err := c.builder.Insert("conversations").
	// 	Columns("zid", "community_zid", "user_did", "text", "link", "img", "video", "public", "public_price", "created", "updated").
	// 	Values(conversation.Zid, conversation.CommunityZid, conversation.UserDid, conversation.Text, conversation.Link, conversation.Img, conversation.Video, conversation.Public, conversation.PublicPrice, conversation.Created, conversation.Updated).
	// 	RunWith(tx).Exec()
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	// conversationId, err := result.LastInsertId()
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	// conversation.Id = conversationId

	// err = tx.Commit()
	// if err != nil {
	// 	return err
	// }

	// return nil
}
