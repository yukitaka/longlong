package repository

import (
	"database/sql"
	"fmt"
	"github.com/yukitaka/longlong/internal/domain/entity"
	rep "github.com/yukitaka/longlong/internal/domain/repository"
	"github.com/yukitaka/longlong/internal/util"
)

type Users struct {
	users map[int]*entity.User
	*sql.DB
}

func NewUsersRepository() rep.Users {
	con, err := sql.Open("sqlite3", "./longlong.db")
	if err != nil {
		util.CheckErr(err)
	}

	return &Users{
		users: make(map[int]*entity.User),
		DB:    con,
	}
}

func (rep *Users) Close() {
	err := rep.DB.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func (rep *Users) Create(name string) (int64, error) {
	query := "select max(id) from users"
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
	query = "insert into users (id) values (?)"
	_, err = rep.DB.Exec(query, id, name)
	if err != nil {
		return -1, err
	}
	query = "insert into profiles (id, name) values (?, ?)"
	_, err = rep.DB.Exec(query, id, name)
	if err != nil {
		return -1, err
	}
	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (rep *Users) Find(id int64) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}
