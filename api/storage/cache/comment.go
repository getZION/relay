package cache

import "github.com/getzion/relay/api"

func (c *cacheStorage) GetComments() ([]api.Comment, error) {
	// var comments []v1.CommentORM
	// result := s.connection.DB.Find(&comments)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	// return comments, nil
	return nil, nil

}

func (c *cacheStorage) InsertComment(*api.Comment) error {
	// currentTime := time.Now().Unix()
	// comment := v1.CommentORM{
	// 	Zid:             uuid.NewString(),
	// 	ConversationZid: model.ConversationZid,
	// 	UserDid:         model.UserDid,
	// 	Text:            model.Text,
	// 	Link:            model.Link,
	// 	Created:         currentTime,
	// 	Updated:         currentTime,
	// }

	// result := s.connection.DB.Create(&comment)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	// return &comment, nil
	return nil
}
