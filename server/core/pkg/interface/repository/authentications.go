package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	rep "github.com/yukitaka/longlong/server/core/pkg/domain/repository"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
	"time"
)

type Authentications struct {
	*datastore.Connection
}

func NewAuthenticationsRepository(con *datastore.Connection) rep.Authentications {
	return &Authentications{
		Connection: con,
	}
}

func (rep *Authentications) Close() {
	rep.Connection.Close()
}

func (rep *Authentications) Create(identify, token string) (int, error) {
	query := "select max(id) from authentications"
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

	query = "insert into authentications (id, identify, token) values ($1, $2, $3)"
	_, err = rep.DB.Exec(query, id, identify, token)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (rep *Authentications) FindToken(identify string) (int, string, error) {
	stmt, err := rep.DB.Preparex("select individual_id, token from authentications where identify=$1")
	if err != nil {
		return -1, "", err
	}
	defer func(stmt *sqlx.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}(stmt)
	var id int
	var token string
	err = stmt.QueryRowx(identify).Scan(&id, &token)
	if err != nil {
		if err == sql.ErrNoRows {
			return -1, "", errors.New(fmt.Sprintf("authentication identify %s is nothing", identify))
		} else {
			return -1, "", err
		}
	}

	return id, token, nil
}

func (rep *Authentications) UpdateToken(id int, token string) error {
	query := "update authentications set token=$1 where id=$2"
	_, err := rep.DB.Exec(query, token, id)
	if err != nil {
		return err
	}

	return nil
}

func (rep *Authentications) StoreOAuth2Info(identify, accessToken, refreshToken string, expiry time.Time) (bool, error) {
	query := "insert into oauth_authentications (identify, access_token, refresh_token, expiry) values ($1, $2, $3, $4)"
	_, err := rep.DB.Exec(query, identify, accessToken, refreshToken, expiry)
	if err != nil {
		return false, err
	}
	return true, nil
}
