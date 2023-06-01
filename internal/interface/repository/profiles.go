package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yukitaka/longlong/internal/domain/entity"
	rep "github.com/yukitaka/longlong/internal/domain/repository"
)

type Profiles struct {
	*sql.DB
}

func NewProfilesRepository(con *sql.DB) rep.Profiles {
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

func (rep *Profiles) Find(id int) (*entity.Profile, error) {
	query := "select nick_name, full_name, biography from profiles where id = ?"
	row := rep.DB.QueryRow(query, id)
	var nickName, fullName, bio string
	err := row.Scan(&nickName, &fullName, &bio)
	if err != nil {
		return nil, err
	}

	return entity.NewProfile(id, nickName, fullName, bio), nil
}
