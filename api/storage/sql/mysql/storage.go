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

	type CommunityORM struct {
		// Conversations   []*ConversationORM `gorm:"foreignkey:community_zid;association_foreignkey:Zid"`
		Created         int64  `gorm:"not null"`
		Deleted         bool   `gorm:"default:false"`
		Description     string `gorm:"size:250;not null"`
		EscrowAmount    int64  `gorm:"not null"`
		Id              int64  `gorm:"primary_key;unique"`
		Img             string
		LastActive      int64
		Name            string `gorm:"size:150;unique;not null"`
		OwnerDid        string `gorm:"not null"`
		OwnerUsername   string `gorm:"not null"`
		PricePerMessage int64  `gorm:"not null"`
		PriceToJoin     int64  `gorm:"not null"`
		Public          bool   `gorm:"default:true"`
		// Tags            []*TagORM `gorm:"foreignkey:Id;association_foreignkey:Id;many2many:community_tags;jointable_foreignkey:CommunityId;association_jointable_foreignkey:TagId"`
		Updated int64
		// Users           []*UserORM `gorm:"foreignkey:Zid;association_foreignkey:Did;many2many:community_users;jointable_foreignkey:CommunityZid;association_jointable_foreignkey:UserDid"`
		Zid string `gorm:"unique;not null"`
	}

	db.AutoMigrate(
		&CommunityORM{},
	// &api.Community{},
	// &api.Conversation{},
	// &api.User{},
	)
	logrus.Info("Migrations successful.")

	return db, nil
}
