package cache

import (
	"github.com/getzion/relay/api"
)

func (c *cacheStorage) GetUsers() ([]api.User, error) {

	return nil, nil
}

func (c *cacheStorage) GetUserByDid(did string) (*api.User, error) {
	// var user v1.UserORM
	// result := s.connection.DB.Model(&user).First(&user, "did = ?", did)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	// return &user, nil
	return nil, nil
}

func (c *cacheStorage) GetUserByUsername(username string) (*api.User, error) {
	return nil, nil
}

func (c *cacheStorage) InsertUser(*api.User) error {
	// currentTime := time.Now().Unix()
	// user := v1.UserORM{
	// 	Name:           model.Name,
	// 	Did:            model.Did,
	// 	Email:          model.Email,
	// 	Username:       model.Username,
	// 	Created:        currentTime,
	// 	Updated:        currentTime,
	// 	Img:            model.Img,
	// 	Bio:            model.Bio,
	// 	PriceToMessage: model.PriceToMessage,
	// }

	// result := s.connection.DB.Create(&user)
	// if result.Error != nil {
	// 	mySqlError := result.Error.(*mysql.MySQLError)
	// 	if mySqlError != nil && mySqlError.Number == 1062 {
	// 		return nil, fmt.Errorf("the specified username already exist: %s", user.Username)
	// 	}
	// 	return nil, result.Error
	// }

	// return &user, nil
	return nil
}
