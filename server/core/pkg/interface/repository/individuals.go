package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	rep "github.com/yukitaka/longlong/server/core/pkg/domain/repository"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
)

type Individuals struct {
	*datastore.Connection
}

func NewIndividualsRepository(con *datastore.Connection) rep.Individuals {
	return &Individuals{
		Connection: con,
	}
}

func (rep *Individuals) Close() {
	rep.Connection.Close()
}

func (rep *Individuals) Create(name string, userId, profileId int) (int, error) {
	query := "select max(id) from individuals"
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
	query = "insert into individuals (id, name, user_id, profile_id) values ($1, $2, $3, $4)"
	_, err = rep.DB.Exec(query, id, name, userId, profileId)
	if err != nil {
		fmt.Println(id)
		return -1, err
	}
	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (rep *Individuals) Find(id int) (*entity.Individual, error) {
	row := rep.DB.QueryRowx("select name, user_id, profile_id from individuals where id = $1", id)

	var name string
	var userId int
	var profileId int
	err := row.Scan(&name, &userId, &profileId)
	if err != nil {
		return nil, err
	}

	user, err := NewUsersRepository(rep.Connection).Find(userId)
	if err != nil {
		return nil, err
	}

	profile, err := NewProfilesRepository(rep.Connection).Find(profileId)
	if err != nil {
		return nil, err
	}

	return &entity.Individual{
		Id:      id,
		Name:    name,
		User:    user,
		Profile: profile,
	}, nil
}

func (rep *Individuals) FindByUserId(userId int) (*[]entity.Individual, error) {
	r, err := rep.DB.Queryx("select id, name, profile_id from individuals where user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	var individuals []entity.Individual
	for r.Next() {
		var id int
		var name string
		var profileId int

		err := r.Scan(&id, &name, &profileId)
		if err != nil {
			return nil, err
		}

		user, err := NewUsersRepository(rep.Connection).Find(userId)
		if err != nil {
			return nil, err
		}

		profile, err := NewProfilesRepository(rep.Connection).Find(profileId)
		if err != nil {
			return nil, err
		}

		individuals = append(individuals, entity.Individual{
			Id:      id,
			Name:    name,
			User:    user,
			Profile: profile,
		})
	}

	return &individuals, nil
}
