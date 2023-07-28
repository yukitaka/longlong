package datastore

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yukitaka/longlong/server/core/pkg/util"
)

type Connection struct {
	*sqlx.DB
}

func NewConnectionOpen(driver string, datasource string) (*Connection, error) {
	con, err := sqlx.Open(driver, datasource)
	if err != nil {
		util.CheckErr(err)
	}

	return &Connection{con}, err
}
