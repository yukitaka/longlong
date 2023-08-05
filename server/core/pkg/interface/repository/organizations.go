package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	rep "github.com/yukitaka/longlong/server/core/pkg/domain/repository"
	"github.com/yukitaka/longlong/server/core/pkg/domain/value_object"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
	"strings"
)

type Organizations struct {
	organizations map[int]*entity.Organization
	*datastore.Connection
}

func NewOrganizationsRepository(con *datastore.Connection) rep.Organizations {
	return &Organizations{
		organizations: make(map[int]*entity.Organization),
		Connection:    con,
	}
}

func (rep *Organizations) Create(name string, individual entity.Individual) (int, error) {
	query := "select max(id) from organizations"
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

	query = "insert into organizations (id, name) values ($1, $2)"
	_, err = rep.DB.Exec(query, id, name)
	if err != nil {
		return -1, err
	}
	query = "insert into organization_members (organization_id, individual_id, role) values ($1, $2, $3)"
	_, err = rep.DB.Exec(query, id, individual.Id, value_object.OWNER)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (rep *Organizations) Find(id int) (*entity.Organization, error) {
	stmt, err := rep.DB.Preparex("select name from organizations where id=$1")
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(fmt.Sprintf("organization id %d is nothing", id))
		} else {
			return nil, err
		}
	}

	return entity.NewOrganization(0, id, name), nil
}

func (rep *Organizations) FindByName(name string) (*entity.Organization, error) {
	stmt, err := rep.DB.Preparex("select id from organizations where name=$1")
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
	err = stmt.QueryRowx(name).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(fmt.Sprintf("organization name %s is nothing", name))
		} else {
			return nil, err
		}
	}

	return entity.NewOrganization(0, id, name), nil
}

func (rep *Organizations) FindAll(ids []interface{}) (*[]entity.Organization, error) {
	stmt, err := rep.DB.Preparex("select parent_id, id, name from organizations where id in ($1" + strings.Repeat(",$2", len(ids)-1) + ")")
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

func (rep *Organizations) List() (*[]entity.Organization, error) {
	rows, err := rep.DB.Queryx("select id, name from organizations")
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
