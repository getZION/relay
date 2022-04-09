package common

import (
	"time"

	"github.com/getzion/relay/api"
	"github.com/google/uuid"
)

func (c *Connection) GetComments() ([]api.Comment, error) {

	rows, err := c.builder.Select("c.id, c.zid, c.conversation_zid, c.user_did, c.text, c.link, c.created, c.updated").From("comments c").Query()
	if err != nil {
		return nil, err
	}

	var comments []api.Comment
	for rows.Next() {
		var tempComment api.Comment
		rows.Scan(&tempComment.Id, &tempComment.Zid, &tempComment.ConversationZid, &tempComment.UserDid, &tempComment.Text, &tempComment.Link, &tempComment.Created, &tempComment.Updated)

		comments = append(comments, tempComment)
	}

	return comments, nil
}

func (c *Connection) InsertComment(comment *api.Comment) error {
	currentTime := time.Now().Unix()
	comment.Zid = uuid.NewString()
	comment.Created = currentTime
	comment.Updated = currentTime

	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	result, err := c.builder.Insert("comments").
		Columns("zid", "conversation_zid", "user_did", "text", "link", "created", "updated").
		Values(comment.Zid, comment.ConversationZid, comment.UserDid, comment.Text, comment.Link, comment.Created, comment.Updated).
		RunWith(tx).Exec()
	if err != nil {
		tx.Rollback()
		return err
	}

	commentId, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	comment.Id = commentId

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
