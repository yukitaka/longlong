package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/yukitaka/longlong/internal/domain/entity"
	rep "github.com/yukitaka/longlong/internal/domain/usecase/repository"
	"github.com/yukitaka/longlong/internal/util"
)

type Organizations struct {
	organizations map[int]*entity.Organization
	sql.DB
}

func NewOrganizationsRepository() rep.Organizations {
	con, err := sql.Open("sqlite3", "./longlong.db")
	if err != nil {
		util.CheckErr(err)
	}

	return &Organizations{
		organizations: make(map[int]*entity.Organization),
		DB:            *con,
	}
}

func (o *Organizations) Close() {
	o.DB.Close()
}

func (o *Organizations) Create(name string) int {
	query := "select max(id) from organizations"
	row := o.DB.QueryRow(query)
	var nullableId sql.NullInt64
	err := row.Scan(&nullableId)
	if err != nil {
		util.CheckErr(err)
		return -1
	}
	id := 0
	if nullableId.Valid {
		id = int(nullableId.Int64)
		id++
	}

	query = "insert into organizations (id, name) values (?, ?)"
	_, err = o.DB.Exec(query, id, name)
	if err != nil {
		util.CheckErr(err)
		return -1
	}
	fmt.Printf("Create Organization %s id %d.\n", name, id)

	return id
}

func (o *Organizations) Find(id int) (*entity.Organization, error) {
	if organization, ok := o.organizations[id]; ok {
		return organization, nil
	}
	fmt.Printf("Call to find Organization id by %d.\n", id)

	return nil, errors.New(fmt.Sprintf("%d is not found.", id))
}
