package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/storage/sql/common"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type mysqlConnectionParams struct {
	Host string `envconfig:"DB_HOST"`
	Name string `envconfig:"DB_NAME"`
	User string `envconfig:"DB_USER"`
	Pass string `envconfig:"DB_PASS"`
}

type mysqlStorage struct {
	*common.Connection
}

func NewMySqlStorage() (*mysqlStorage, error) {
	var err error
	var params mysqlConnectionParams
	envconfig.Process("", &params)

	databaseConnectionString := fmt.Sprintf("%s:%s@tcp(%s:25060)/%s?multiStatements=true", params.User, params.Pass, params.Host, params.Name)

	db, err := gorm.Open(mysql.Open(databaseConnectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}
	logrus.Info("Connected to MySQL database.")

	db.AutoMigrate(
		&api.Community{},
		&api.Conversation{},
		&api.Message{},
		&api.Payment{},
		&api.User{},
	)
	logrus.Info("Migrations successful.")

	connection := mysqlStorage{
		Connection: common.NewStore(db),
	}

	return &connection, nil
}
