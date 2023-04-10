package repository

import (
	"database/sql"
	"fmt"
	"github.com/yukitaka/longlong/internal/domain/entity"
	rep "github.com/yukitaka/longlong/internal/domain/repository"
	"github.com/yukitaka/longlong/internal/util"
)

type Avatars struct {
	*sql.DB
}

func NewAvatarsRepository() rep.Avatars {
	con, err := sql.Open("sqlite3", "./longlong.db")
	if err != nil {
		util.CheckErr(err)
	}

	return &Avatars{
		DB: con,
	}
}

func (rep *Avatars) Close() {
	err := rep.DB.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func (rep *Avatars) Create(name string, userId, profileId int64) (int64, error) {
	query := "select max(id) from avatars"
	row := rep.DB.QueryRow(query)
	var nullableId sql.NullInt64
	err := row.Scan(&nullableId)
	if err != nil {
		return -1, err
	}
	id := int64(0)
	if nullableId.Valid {
		id = nullableId.Int64
		id++
	}

	tx, err := rep.DB.Begin()
	if err != nil {
		return -1, err
	}
	query = "insert into avatars (id, name, user_id, profile_id) values (?, ?, ?, ?)"
	_, err = rep.DB.Exec(query, id, name, userId, profileId)
	if err != nil {
		return -1, err
	}
	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (rep *Avatars) Find(id int64) (*entity.Avatar, error) {
	//TODO implement me
	panic("implement me")
}
