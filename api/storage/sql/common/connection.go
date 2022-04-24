package common

import (
	"gorm.io/gorm"
)

type Connection struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Connection {
	return &Connection{
		db: db,
	}
}
