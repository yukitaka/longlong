package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
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

func (rep *Authentications) Create(organizationId int, identify, token string) (int, error) {
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

	query = "insert into authentications (id, organization_id, identify, token) values ($1, $2, $3, $4)"
	_, err = rep.DB.Exec(query, id, organizationId, identify, token)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (rep *Authentications) Find(organizationId int, identify string) (*entity.Authentication, error) {
	stmt, err := rep.DB.Preparex("select individual_id, token from authentications where organization_id=$1 and identify=$2")
	if err != nil {
		return nil, err
	}
	defer func(stmt *sqlx.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}(stmt)
	var id int
	var token string
	err = stmt.QueryRowx(organizationId, identify).Scan(&id, &token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(fmt.Sprintf("authentication identify %s is nothing", identify))
		} else {
			return nil, err
		}
	}
	return entity.NewAuthentication(id, organizationId, identify, token), nil
}

func (rep *Authentications) FindToken(organizationId int, identify string) (int, string, error) {
	if organizationId > 0 {
		auth, err := rep.Find(organizationId, identify)
		if err != nil {
			return -1, "", err
		}

		return auth.Id, auth.Token, nil
	} else {
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
			if errors.Is(err, sql.ErrNoRows) {
				return -1, "", errors.New(fmt.Sprintf("authentication identify %s is nothing", identify))
			} else {
				return -1, "", err
			}
		}
		return id, token, nil
	}
}

func (rep *Authentications) UpdateToken(id int, token string) error {
	query := "update authentications set token=$1 where id=$2"
	_, err := rep.DB.Exec(query, token, id)
	if err != nil {
		return err
	}

	return nil
}

func (rep *Authentications) Store(organizationId int, identify, token string) (bool, error) {
	query := "select max(id) from authentications"
	row := rep.DB.QueryRowx(query)
	var nullableId sql.NullInt32
	err := row.Scan(&nullableId)
	if err != nil {
		return false, err
	}
	id := 0
	if nullableId.Valid {
		id = int(nullableId.Int32)
		id++
	}

	query = "insert into authentications (id, organization_id, identify, token) values ($1, $2, $3, $4)"
	_, err = rep.DB.Exec(query, id, organizationId, identify, token)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (rep *Authentications) StoreOAuth2Info(identify, accessToken, refreshToken string, expiry time.Time) (bool, error) {
	query := "insert into oauth_authentications (identify, access_token, refresh_token, expiry) values ($1, $2, $3, $4)"
	_, err := rep.DB.Exec(query, identify, accessToken, refreshToken, expiry)
	if err != nil {
		return false, err
	}
	return true, nil
}
