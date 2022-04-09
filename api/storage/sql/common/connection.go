package common

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type Connection struct {
	builder sq.StatementBuilderType
	db      *sql.DB
}

func NewStore(db *sql.DB, builder sq.StatementBuilderType) *Connection {
	return &Connection{
		db:      db,
		builder: builder,
	}
}
