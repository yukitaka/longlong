package datastore

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yukitaka/longlong/internal/util"
)

func NewConnectionOpen(driver string, datasource string) (*sql.DB, error) {
	con, err := sql.Open(driver, datasource)
	if err != nil {
		util.CheckErr(err)
	}

	return con, err
}
