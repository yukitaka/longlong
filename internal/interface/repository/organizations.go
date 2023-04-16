package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/yukitaka/longlong/internal/domain/entity"
	rep "github.com/yukitaka/longlong/internal/domain/repository"
	"github.com/yukitaka/longlong/internal/domain/value_object"
	"github.com/yukitaka/longlong/internal/util"
	"strings"
)

type Organizations struct {
	organizations map[int]*entity.Organization
	*sql.DB
}

func NewOrganizationsRepository() rep.Organizations {
	con, err := sql.Open("sqlite3", "./longlong.db")
	if err != nil {
		util.CheckErr(err)
	}

	return &Organizations{
		organizations: make(map[int]*entity.Organization),
		DB:            con,
	}
}

func (o *Organizations) Close() {
	err := o.DB.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func (o *Organizations) Create(name string, individual entity.Individual) (int64, error) {
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
	query = "insert into organization_belongings (organization_id, individual_id, role) values (?, ?, ?)"
	_, err = o.DB.Exec(query, id, name, individual.Id, value_object.OWNER)
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
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}(stmt)
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

func (o *Organizations) FindAll(ids []interface{}) (*[]entity.Organization, error) {
	stmt, err := o.DB.Prepare("select parent_id, name from organizations where id in (?" + strings.Repeat(",?", len(ids)-1) + ")")
	if err != nil {
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}(stmt)
	var parentId int64
	var name string
	err = stmt.QueryRow(ids...).Scan(&parentId, &name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(fmt.Sprintf("organization ids %d are nothing", ids))
		} else {
			return nil, err
		}
	}

	return nil, nil
}

func (o *Organizations) List() (*[]entity.Organization, error) {
	rows, err := o.DB.Query("select id, name from organizations")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}(rows)

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
