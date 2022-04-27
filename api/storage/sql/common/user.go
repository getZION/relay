package common

import (
	"fmt"
	"time"

	"github.com/getzion/relay/api"
	"github.com/sirupsen/logrus"
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
	logrus.Infof("GetUserByDid: Looking for %s", did)
	var user api.User
	result := c.db.First(&user, "did = ?", did)
	if result.Error != nil {
		return nil, result.Error
	}
	logrus.Infof("GetUserByDid: Returning %s", user.Username)
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

	logrus.Infof("Created user. Attempting to join Lounge community...")
	loungeZidCypher := "016cf037-6247-4b69-90e0-e83e2aad98fb"
	community, err := c.GetCommunityByZid(loungeZidCypher)
	if err != nil {
		return fmt.Errorf("community not found: %s", loungeZidCypher)
	}
	err = c.AddUserToCommunity(community, user)
	if err != nil {
		return err
	}
	logrus.Infof("Success? %s", user.Did)

	return nil
}
