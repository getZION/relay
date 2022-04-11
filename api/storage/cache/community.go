package cache

import (
	"github.com/getzion/relay/api"
)

func (c *cacheStorage) GetCommunities() ([]api.Community, error) {

	iterator, _ := c.db.Txn(false).Get("communities", "id")

	var communities []api.Community
	for obj := iterator.Next(); obj != nil; obj = iterator.Next() {
		c := obj.(*api.Community)
		communities = append(communities, *c)
	}

	return communities, nil
}

func (c *cacheStorage) GetCommunityByZid(string) (*api.Community, error) {
	// var community v1.CommunityORM
	// result := s.connection.DB.Model(&community).First(&community, "zid = ?", zid)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	// return &community, nil
	return nil, nil
}

func (c *cacheStorage) InsertCommunity(community *api.Community) error {

	tx := c.db.Txn(true)
	err := tx.Insert("communities", community)
	tx.Commit()

	return err
}

func (c *cacheStorage) AddUserToCommunity(communityZid, userDid string) error {
	// association := s.connection.DB.Model(community).Omit("users").Association("users").Append(user)
	// return association.Error
	return nil
}

func (c *cacheStorage) RemoveUserToCommunity(communityZid, userDid, leftReason string) error {
	// association := s.connection.DB.Model(community).Omit("users").Association("users").Delete(user)
	// return association.Error
	return nil
}
