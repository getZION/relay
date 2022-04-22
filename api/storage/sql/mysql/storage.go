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

	databaseConnectionString := fmt.Sprintf("%s:%s@tcp(%s:25060)/%s?multiStatements=true", params.User, params.Pass, params.Host, params.Name)

	db, err := gorm.Open(mysql.Open(databaseConnectionString), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	logrus.Info("Connected to MySQL database.")

	type Community struct {
		Created         int64  `gorm:"not null"`
		Deleted         bool   `gorm:"default:false"`
		Description     string `gorm:"size:250;not null"`
		EscrowAmount    int64  `gorm:"not null"`
		Img             string
		LastActive      int64
		Name            string `gorm:"size:150;unique;not null"`
		OwnerDid        string `gorm:"not null"`
		OwnerUsername   string `gorm:"not null"`
		PricePerMessage int64  `gorm:"not null"`
		PriceToJoin     int64  `gorm:"not null"`
		Updated         int64
		Zid             string `gorm:"primary_key;unique;not null"`
	}

	db.AutoMigrate(
		&Community{},
	// &api.Community{},
	// &api.Conversation{},
	// &api.User{},
	)
	logrus.Info("Migrations successful.")

	return db, nil
}
