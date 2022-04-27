package common

import (
	"time"

	"github.com/getzion/relay/api"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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
	c.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(api.User{}).Updates(api.User{})
	c.db.Model(&user).Association("Communities").Append(&community)
	logrus.Info("AddUserToCommunity done maybe?!")
	return nil
}

func (c *Connection) RemoveUserFromCommunity(community *api.Community, user *api.User) error {
	return nil
	// var exist bool
	// err := c.db.QueryRow(fmt.Sprintf(`SELECT EXISTS(SELECT id FROM community_users cu WHERE cu.community_zid = '%s' AND cu.user_did = '%s' AND cu.left_date IS NULL)`, communityZid, userDid)).Scan(&exist)
	// if !exist {
	// 	return fmt.Errorf("user already doesn't member of this community")
	// }

	// tx, err := c.db.Begin()
	// if err != nil {
	// 	return err
	// }

	// sqlBuilder := c.builder.Update("community_users").Set("left_date", time.Now().Unix())

	// if leftReason != "" {
	// 	sqlBuilder = sqlBuilder.Set("left_reason", leftReason)
	// }

	// sqlBuilder = sqlBuilder.Where(sq.And{
	// 	sq.Eq{"community_zid": communityZid},
	// 	sq.Eq{"user_did": userDid},
	// 	sq.Eq{"left_date": nil},
	// })

	// _, err = sqlBuilder.RunWith(tx).Exec()
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	// err = tx.Commit()
	// if err != nil {
	// 	return err
	// }

	// return nil
}
