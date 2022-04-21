package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type mysqlConnectionParams struct {
	Host string `envconfig:"DB_HOST"`
	Name string `envconfig:"DB_NAME"`
	User string `envconfig:"DB_USER"`
	Pass string `envconfig:"DB_PASS"`
}

func NewMySqlStorage() (*gorm.DB, error) {
	var err error
	var params mysqlConnectionParams
	envconfig.Process("", &params)

	databaseConnectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?multiStatements=true", params.User, params.Pass, params.Host, params.Name)

	db, err := gorm.Open(mysql.Open(databaseConnectionString), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	logrus.Infof("Connected to MySQL database")

	return db, nil
}
