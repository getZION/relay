package storage

import (
	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/storage/sql/mysql"
	"github.com/sirupsen/logrus"
)

func NewStorage(storeType string) (storage api.Storage, err error) {
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
