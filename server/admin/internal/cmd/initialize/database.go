package initialize

import (
	"github.com/jmoiron/sqlx"
)

type Database struct {
	*sqlx.DB
}

func NewDatabase(db *sqlx.DB) (*Database, error) {
	return &Database{db}, nil
}
