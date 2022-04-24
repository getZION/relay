package common

import (
	"github.com/getzion/relay/api"
)

func (c *Connection) GetUsers() ([]api.User, error) {
	return nil, nil
	// row := c.db.QueryRow("CALL get_users")

	// var users []api.User
	// var jsonUsers string

	// if err := row.Scan(&jsonUsers); err != nil {
	// 	return nil, err
	// }

	// err := json.Unmarshal([]byte(jsonUsers), &users)
	// if err != nil {
	// 	return nil, err
	// }

	// return users, nil
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
	return nil
	// currentTime := time.Now().Unix()
	// user.Created = currentTime
	// user.Updated = currentTime

	// tx, err := c.db.Begin()
	// if err != nil {
	// 	return err
	// }

	// result, err := c.builder.Insert("users").
	// 	Columns("did", "username", "email", "name", "bio", "img", "price_to_message", "created", "updated").
	// 	Values(user.Did, user.Username, user.Email, user.Name, user.Bio, user.Img, user.PriceToMessage, user.Created, user.Updated).
	// 	RunWith(tx).Exec()
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	// userId, err := result.LastInsertId()
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	// user.Id = userId

	// err = tx.Commit()
	// if err != nil {
	// 	return err
	// }

	// return nil
}
