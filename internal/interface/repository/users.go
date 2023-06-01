package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yukitaka/longlong/internal/domain/entity"
	rep "github.com/yukitaka/longlong/internal/domain/repository"
)

type Users struct {
	users map[int]*entity.User
	*sql.DB
}

func NewUsersRepository(con *sql.DB) rep.Users {
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

func (rep *Users) Create(name string) (int, error) {
	query := "select max(id) from users"
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
	query = "insert into users (id) values (?)"
	_, err = rep.DB.Exec(query, id, name)
	if err != nil {
		return -1, err
	}
	query = "insert into profiles (id, full_name) values (?, ?)"
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

func (rep *Users) Find(id int) (*entity.User, error) {
	return entity.NewUser(id), nil
}
