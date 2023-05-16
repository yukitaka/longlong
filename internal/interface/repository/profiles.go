package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	rep "github.com/yukitaka/longlong/internal/domain/repository"
	"github.com/yukitaka/longlong/internal/util"
)

type Profiles struct {
	*sql.DB
}

func NewProfilesRepository() rep.Profiles {
	con, err := sql.Open("sqlite3", "./longlong.db")
	if err != nil {
		util.CheckErr(err)
	}

	return &Profiles{
		DB: con,
	}
}

func (rep *Profiles) Close() {
	err := rep.DB.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func (rep *Profiles) Create(nickName, fullName, bio string) (int, error) {
	query := "select max(id) from profiles"
	row := rep.DB.QueryRow(query)
	var nullableId sql.NullInt32
	err := row.Scan(&nullableId)
	if err != nil {
		return -1, err
	}
	id := 0
	if nullableId.Valid {
		id = int(nullableId.Int32)
		id++
	}

	tx, err := rep.DB.Begin()
	if err != nil {
		return -1, err
	}
	query = "insert into profiles (id, nick_name, full_name, biography) values (?, ?, ?, ?)"
	_, err = rep.DB.Exec(query, id, nickName, fullName, bio)
	if err != nil {
		return -1, err
	}
	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return id, nil
}
