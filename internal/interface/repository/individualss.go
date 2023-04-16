package repository

import (
	"database/sql"
	"fmt"
	"github.com/yukitaka/longlong/internal/domain/entity"
	rep "github.com/yukitaka/longlong/internal/domain/repository"
	"github.com/yukitaka/longlong/internal/util"
)

type Individuals struct {
	*sql.DB
}

func NewIndividualsRepository() rep.Individuals {
	con, err := sql.Open("sqlite3", "./longlong.db")
	if err != nil {
		util.CheckErr(err)
	}

	return &Individuals{
		DB: con,
	}
}

func (rep *Individuals) Close() {
	err := rep.DB.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func (rep *Individuals) Create(name string, userId, profileId int64) (int64, error) {
	query := "select max(id) from individuals"
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
	query = "insert into individuals (id, name, user_id, profile_id) values (?, ?, ?, ?)"
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

func (rep *Individuals) Find(id int64) (*entity.Individual, error) {
	//TODO implement me
	panic("implement me")
}

func (rep *Individuals) FindByUserId(userId int64) (*[]entity.Individual, error) {
	r, err := rep.DB.Query("select id, name, profile_id from individuals where user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	var individuals []entity.Individual
	for r.Next() {
		var id int64
		var name string
		var profileId int64

		err := r.Scan(&id, &name, &profileId)
		if err != nil {
			return nil, err
		}
		individuals = append(individuals, entity.Individual{
			Id:        id,
			Name:      name,
			UserId:    userId,
			ProfileId: profileId,
		})
	}

	return &individuals, nil
}
