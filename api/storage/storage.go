package storage

import (
	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/storage/cache"
	"github.com/getzion/relay/api/storage/sql/mysql"
	"github.com/sirupsen/logrus"
)

// NewStorage should use config options to return a connection to the requested database
func NewStorage(storeType string) (storage api.Storage, err error) {

	switch storeType {
	case "mysql":
		storage, err = mysql.NewMySqlStorage()
	case "cache":
		storage, err = cache.NewStorage()
	default:
		logrus.Infof("unknown storage database: %s, cache storage activated", storeType)
		storage, err = cache.NewStorage()
	}

	if err != nil {
		return nil, err
	}

	return storage, nil
}
