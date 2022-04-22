package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type mysqlConnectionParams struct {
	Host        string `envconfig:"DB_HOST"`
	Name        string `envconfig:"DB_NAME"`
	User        string `envconfig:"DB_USER"`
	Pass        string `envconfig:"DB_PASS"`
	DatabaseUrl string `envconfig:"DATABASE_URL"`
}

func NewMySqlStorage() (*gorm.DB, error) {
	var err error
	var params mysqlConnectionParams
	envconfig.Process("", &params)

	databaseConnectionString := fmt.Sprintf("%s:%s@tcp(%s:25060)/%s?&ssl-mode=REQUIRED", params.User, params.Pass, params.Host, params.Name) // multiStatements=true
	databaseConnectionString2 := params.DatabaseUrl
	logrus.Infof("Comparing 1: %s", databaseConnectionString)
	logrus.Infof("Comparing 2: %s", databaseConnectionString2)

	db, err := gorm.Open(mysql.Open(databaseConnectionString), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	logrus.Infof("Connected to MySQL database.")

	return db, nil
}
