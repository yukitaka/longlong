package repository

import (
	"database/sql"
	"fmt"
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	rep "github.com/yukitaka/longlong/server/core/pkg/domain/repository"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
)

type Users struct {
	users map[int]*entity.User
	*datastore.Connection
}

func NewUsersRepository(con *datastore.Connection) rep.Users {
	return &Users{
		users:      make(map[int]*entity.User),
		Connection: con,
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
	row := rep.DB.QueryRowx(query)
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
	query = "insert into users (id) values ($1)"
	_, err = rep.DB.Exec(query, id)
	if err != nil {
		return -1, err
	}
	query = "insert into profiles (id, full_name) values ($1, $2)"
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
