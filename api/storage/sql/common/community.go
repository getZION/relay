package common

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/getzion/relay/api"
	"github.com/google/uuid"
)

func (c *Connection) GetCommunities() ([]api.Community, error) {

	row := c.db.QueryRow("CALL get_communities")

	var communities []api.Community
	var jsonCommunities string

	if err := row.Scan(&jsonCommunities); err != nil {
		return nil, err
	}

	err := json.Unmarshal([]byte(jsonCommunities), &communities)
	if err != nil {
		return nil, err
	}

	return communities, nil
}

func (c *Connection) GetCommunityByZid(zid string) (*api.Community, error) {

	var community api.Community
	err := c.builder.Select("c.id, c.zid").From("communities c").Where(sq.Eq{"zid": zid}).QueryRow().Scan(&community.Id, &community.Zid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("community not found %s", zid)
		}

		return nil, err
	}

	tagRows, err := c.builder.Select("ct.tag").From("community_tags ct").Where(sq.Eq{"community_zid": zid}).Query()
	if err != nil {
		return nil, err
	}

	var tags []string
	for tagRows.Next() {
		var tag string
		tagRows.Scan(&tag)
		tags = append(tags, tag)
	}

	tagRows.Close()
	community.Tags = tags

	userRows, err := c.builder.Select("cu.id, cu.community_zid, cu.user_did, cu.joined_date, cu.left_date").From("community_users cu").Where(sq.Eq{"community_zid": zid}).Query()
	if err != nil {
		return nil, err
	}

	var users []api.UserCommunity
	for userRows.Next() {
		var user api.UserCommunity
		userRows.Scan(&user.Id, &user.CommunityZid, &user.UserDid, &user.JoinedDate, &user.LeftDate)
		users = append(users, user)
	}
	userRows.Close()
	community.Users = users

	return &community, nil
}

func (c *Connection) InsertCommunity(community *api.Community) error {

	currentTime := time.Now().Unix()
	community.Zid = uuid.NewString()
	community.Created = currentTime
	community.Updated = currentTime
	community.LastActive = currentTime

	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	result, err := c.builder.Insert("communities").
		Columns("zid", "name", "owner_did", "owner_username", "description", "escrow_amount", "img", "last_active", "price_per_message", "price_to_join", "public", "created", "updated", "deleted").
		Values(community.Zid, community.Name, community.OwnerDid, community.OwnerUsername, community.Description, community.EscrowAmount, community.Img, community.LastActive, community.PricePerMessage, community.PriceToJoin, community.Public, community.Created, community.Updated, community.Deleted).
		RunWith(tx).Exec()
	if err != nil {
		tx.Rollback()
		return err
	}

	communityId, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	community.Id = communityId

	for _, t := range community.Tags {
		_, err := tx.Exec(fmt.Sprintf(`INSERT INTO tags (tag) SELECT * FROM (SELECT '%s' as t) AS tmp WHERE NOT EXISTS (SELECT tag FROM tags WHERE tag = tmp.t) LIMIT 1`, t))
		if err != nil {
			tx.Rollback()
			return err
		}

		_, err = tx.Exec(fmt.Sprintf(`INSERT INTO community_tags (community_zid, tag) VALUES ('%s', '%s')`, community.Zid, t))
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) AddUserToCommunity(communityZid, userDid string) error {

	var exist bool
	err := c.db.QueryRow(fmt.Sprintf(`SELECT EXISTS(SELECT id FROM relay3.community_users cu WHERE cu.community_zid = '%s' AND cu.user_did = '%s' AND cu.left_date IS NULL)`, communityZid, userDid)).Scan(&exist)
	if exist {
		return fmt.Errorf("user already member of this community")
	}

	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(fmt.Sprintf(`INSERT INTO community_users (community_zid, user_did, joined_date) VALUES ('%s', '%s', %d)`, communityZid, userDid, time.Now().Unix()))
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) RemoveUserToCommunity(communityZid, userDid, leftReason string) error {

	var exist bool
	err := c.db.QueryRow(fmt.Sprintf(`SELECT EXISTS(SELECT id FROM relay3.community_users cu WHERE cu.community_zid = '%s' AND cu.user_did = '%s' AND cu.left_date IS NULL)`, communityZid, userDid)).Scan(&exist)
	if !exist {
		return fmt.Errorf("user already doesn't member of this community")
	}

	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	sqlBuilder := c.builder.Update("community_users").Set("left_date", time.Now().Unix())

	if leftReason != "" {
		sqlBuilder = sqlBuilder.Set("left_reason", leftReason)
	}

	sqlBuilder = sqlBuilder.Where(sq.And{
		sq.Eq{"community_zid": communityZid},
		sq.Eq{"user_did": userDid},
		sq.Eq{"left_date": nil},
	})

	_, err = sqlBuilder.RunWith(tx).Exec()
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
