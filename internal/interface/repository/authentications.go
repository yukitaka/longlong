package repository

import (
	"database/sql"
	"errors"
	"fmt"
	rep "github.com/yukitaka/longlong/internal/domain/repository"
	"github.com/yukitaka/longlong/internal/util"
)

type Authentications struct {
	*sql.DB
}

func NewAuthenticationsRepository() rep.Authentications {
	con, err := sql.Open("sqlite3", "./longlong.db")
	if err != nil {
		util.CheckErr(err)
	}

	return &Authentications{
		DB: con,
	}
}

func (rep *Authentications) Close() {
	err := rep.DB.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func (rep *Authentications) Create(identify, token string) (int, error) {
	query := "select max(id) from authentications"
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

	query = "insert into authentications (id, identify, token) values (?, ?)"
	_, err = rep.DB.Exec(query, id, identify, token)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (rep *Authentications) FindToken(identify string) (int, string, error) {
	stmt, err := rep.DB.Prepare("select individual_id, token from authentications where identify=?")
	if err != nil {
		return -1, "", err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}(stmt)
	var id int
	var token string
	err = stmt.QueryRow(identify).Scan(&id, &token)
	if err != nil {
		if err == sql.ErrNoRows {
			return -1, "", errors.New(fmt.Sprintf("authentication identify %s is nothing", identify))
		} else {
			return -1, "", err
		}
	}

	return id, token, nil
}
