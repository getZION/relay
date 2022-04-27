package common

import (
	"time"

	"github.com/getzion/relay/api"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (c *Connection) GetCommunities() ([]api.Community, error) {
	var communities []api.Community
	result := c.db.Preload("Users").Find(&communities)
	if result.Error != nil {
		return nil, result.Error
	}
	return communities, nil
}

func (c *Connection) GetCommunityByZid(zid string) (*api.Community, error) {
	logrus.Infof("GetCommunityByZid: Looking for %s", zid)
	var community api.Community
	result := c.db.First(&community, "zid = ?", zid)
	if result.Error != nil {
		return nil, result.Error
	}
	logrus.Infof("GetCommunityByZid: Returning %s", community.Name)
	return &community, nil
}

func (c *Connection) InsertCommunity(community *api.Community) error {
	currentTime := time.Now().Unix()
	community.Zid = uuid.NewString()
	community.Created = currentTime
	community.Updated = currentTime
	community.LastActive = currentTime

	result := c.db.Create(community)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *Connection) AddUserToCommunity(community *api.Community, user *api.User) error {
	logrus.Infof("[AddUserToCommunity] community: %s", community.Name)
	logrus.Infof("[AddUserToCommunity] user: %s", user.Name)

	// c.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)

	association := c.db.Model(community).Omit("Users").Association("Users").Append(user)
	logrus.Info("AddUserToCommunity done maybe?!!!!!")
	logrus.Info("AddUserToCommunity error?", association.Error())
	return association

	// c.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(api.User{}).Updates(api.User{})
	// c.db.Model(&user).Association("Communities").Append(&community)
}

func (c *Connection) RemoveUserFromCommunity(community *api.Community, user *api.User) error {
	association := c.db.Model(community).Omit("users").Association("users").Append(user)
	return association
}
