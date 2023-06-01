package datastore

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yukitaka/longlong/internal/util"
)

func NewSqliteOpen() (*sql.DB, error) {
	con, err := sql.Open("sqlite3", "./longlong.db")
	if err != nil {
		util.CheckErr(err)
	}

	return con, err
}
