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
	return nil, nil
	// var user api.User
	// err := c.builder.Select("u.id, u.did, u.username, u.email, u.name, u.bio, u.img, u.price_to_message, u.created, u.updated").
	// 	From("users u").
	// 	Where(sq.Eq{"did": did}).
	// 	QueryRow().
	// 	Scan(&user.Id, &user.Did, &user.Username, &user.Email, &user.Name, &user.Bio, &user.Img, &user.PriceToMessage, &user.Created, &user.Updated)
	// if err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		return nil, fmt.Errorf("user not found: %s", did)
	// 	}

	// 	return nil, err
	// }

	// return &user, nil
}

func (c *Connection) GetUserByUsername(username string) (*api.User, error) {
	return nil, nil
	// var user api.User
	// err := c.builder.Select("u.id, u.did, u.username, u.email, u.name, u.bio, u.img, u.price_to_message, u.created, u.updated").
	// 	From("users u").
	// 	Where(sq.Eq{"username": username}).
	// 	QueryRow().
	// 	Scan(&user.Id, &user.Did, &user.Username, &user.Email, &user.Name, &user.Bio, &user.Img, &user.PriceToMessage, &user.Created, &user.Updated)
	// if err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		return nil, fmt.Errorf("user not found: %s", username)
	// 	}

	// 	return nil, err
	// }

	// return &user, nil
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
