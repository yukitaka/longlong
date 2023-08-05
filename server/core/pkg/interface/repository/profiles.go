package repository

import (
	"database/sql"
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	rep "github.com/yukitaka/longlong/server/core/pkg/domain/repository"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
)

type Profiles struct {
	*datastore.Connection
}

func NewProfilesRepository(con *datastore.Connection) rep.Profiles {
	return &Profiles{
		Connection: con,
	}
}

func (rep *Profiles) Create(nickName, fullName, bio string) (int, error) {
	query := "select max(id) from profiles"
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
	query = "insert into profiles (id, nick_name, full_name, biography) values ($1, $2, $3, $4)"
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
	query := "select nick_name, full_name, biography from profiles where id = $1"
	var nickName, fullName, bio string
	err := rep.DB.QueryRowx(query, id).Scan(&nickName, &fullName, &bio)
	if err != nil {
		return nil, err
	}

	return entity.NewProfile(id, nickName, fullName, bio), nil
}
