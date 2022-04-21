package storage

import (
	"github.com/getzion/relay/api/storage/sql/mysql"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// NewStorage should use config options to return a connection to the requested database
func NewStorage(storeType string) (storage *gorm.DB, err error) {

	switch storeType {
	case "mysql":
		storage, err = mysql.NewMySqlStorage()
	default:
		logrus.Panic("No")
	}
	// case "cache":
	// 	storage, err = cache.NewStorage()
	// default:
	// 	logrus.Infof("unknown storage database: %s, cache storage activated", storeType)
	// 	storage, err = cache.NewStorage()
	// }

	if err != nil {
		return nil, err
	}

	return storage, nil
}
