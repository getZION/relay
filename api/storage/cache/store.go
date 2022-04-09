package cache

import (
	"github.com/hashicorp/go-memdb"
	"github.com/sirupsen/logrus"
)

type cacheStorage struct {
	db *memdb.MemDB
}

// NewStore creates a new *CacheStore
func NewStorage() (*cacheStorage, error) {

	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"users": {
				Name: "users",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "Id"},
					},
					"did": {
						Name:    "did",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Did"},
					},
					"username": {
						Name:    "username",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Username"},
					},
					"name": {
						Name:    "name",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
				},
			},
			"communities": {
				Name: "communities",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "Id"},
					},
					"zid": {
						Name:    "zid",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Zid"},
					},
					"owner_did": {
						Name:    "owner_did",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "OwnerDid"},
					},
					"owner_username": {
						Name:    "owner_username",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "OwnerUsername"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		logrus.Panic(err)
	}

	return &cacheStorage{
		db: db,
	}, nil
}
