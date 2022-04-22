package storage

import (
	"github.com/getzion/relay/api/storage/sql/mysql"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewStorage(storeType string) (storage *gorm.DB, err error) {
	switch storeType {
	case "mysql":
		storage, err = mysql.NewMySqlStorage()
	default:
		logrus.Panic("Unknown storage type")
	}

	if err != nil {
		return nil, err
	}

	return storage, nil
}
