package mysql

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/getzion/relay/api/storage/sql/common"
	_ "github.com/go-sql-driver/mysql"

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

	databaseConnectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?multiStatements=true", params.User, params.Pass, params.Host, params.Name)

	db, err := sql.Open("mysql", databaseConnectionString)
	if err != nil {
		return nil, err
	}
	logrus.Infof("Connected to MySQL database")

	connection := mysqlStorage{
		Connection: common.NewStore(db, sq.StatementBuilder.RunWith(sq.NewStmtCache(db))),
	}

	return &connection, nil
}
