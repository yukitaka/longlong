package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yukitaka/longlong/internal/domain/entity"
	rep "github.com/yukitaka/longlong/internal/domain/repository"
	"github.com/yukitaka/longlong/internal/domain/value_object"
	"strings"
)

type Organizations struct {
	organizations map[int]*entity.Organization
	*sqlx.DB
}

func NewOrganizationsRepository(con *sqlx.DB) rep.Organizations {
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

func (o *Organizations) Create(name string, individual entity.Individual) (int, error) {
	query := "select max(id) from organizations"
	row := o.DB.QueryRowx(query)
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

	query = "insert into organizations (id, name) values (?, ?)"
	_, err = o.DB.Exec(query, id, name)
	if err != nil {
		return -1, err
	}
	query = "insert into organization_members (organization_id, individual_id, role) values (?, ?, ?)"
	_, err = o.DB.Exec(query, id, name, individual.Id, value_object.OWNER)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (o *Organizations) Find(id int) (*entity.Organization, error) {
	stmt, err := o.DB.Preparex("select name from organizations where id=?")
	if err != nil {
		return nil, err
	}
	defer func(stmt *sqlx.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}(stmt)
	var name string
	err = stmt.QueryRowx(id).Scan(&name)
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
	stmt, err := o.DB.Preparex("select parent_id, id, name from organizations where id in ($1" + strings.Repeat(",$2", len(ids)-1) + ")")
	if err != nil {
		return nil, err
	}
	defer func(stmt *sqlx.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}(stmt)
	res, err := stmt.Queryx(ids...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(fmt.Sprintf("organization ids %d are nothing", ids))
		} else {
			return nil, err
		}
	}
	var organizations []entity.Organization
	for res.Next() {
		var parentId int
		var id int
		var name string
		err = res.Scan(&parentId, &id, &name)
		if err != nil {
			return nil, err
		}
		organizations = append(organizations, *entity.NewOrganization(parentId, id, name))
	}

	return &organizations, nil
}

func (o *Organizations) List() (*[]entity.Organization, error) {
	rows, err := o.DB.Queryx("select id, name from organizations")
	if err != nil {
		return nil, err
	}
	defer func(rows *sqlx.Rows) {
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
