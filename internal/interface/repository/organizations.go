package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/yukitaka/longlong/internal/domain/entity"
	rep "github.com/yukitaka/longlong/internal/domain/repository"
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

func (o *Organizations) Create(name string) (int64, error) {
	query := "select max(id) from organizations"
	row := o.DB.QueryRow(query)
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

	query = "insert into organizations (id, name) values (?, ?)"
	_, err = o.DB.Exec(query, id, name)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (o *Organizations) Find(id int64) (*entity.Organization, error) {
	stmt, err := o.DB.Prepare("select name from organizations where id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow(id).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(fmt.Sprintf("organization id %d is nothing", id))
		} else {
			return nil, err
		}
	}

	return entity.NewOrganization(0, id, name), nil
}

func (o *Organizations) List() (*[]entity.Organization, error) {
	rows, err := o.DB.Query("select id, name from organizations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var organizations []entity.Organization
	for rows.Next() {
		var organization entity.Organization
		err = rows.Scan(&organization.Id, &organization.Name)
		if err != nil {
			return nil, err
		}
		organizations = append(organizations, organization)
	}

	return &organizations, nil
}
