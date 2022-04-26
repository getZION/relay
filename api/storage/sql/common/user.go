package common

import (
	"time"

	"github.com/getzion/relay/api"
)

func (c *Connection) GetUsers() ([]api.User, error) {
	var users []api.User
	result := c.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (c *Connection) GetUserByDid(did string) (*api.User, error) {
	var user api.User
	result := c.db.First(&user, did)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (c *Connection) InsertUser(user *api.User) error {
	currentTime := time.Now().Unix()
	user.Created = currentTime
	user.Updated = currentTime

	result := c.db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
