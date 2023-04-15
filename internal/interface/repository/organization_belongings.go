package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/yukitaka/longlong/internal/domain/entity"
	rep "github.com/yukitaka/longlong/internal/domain/repository"
	"github.com/yukitaka/longlong/internal/util"
)

type OrganizationBelongings struct {
	organization *entity.Organization
	*sql.DB
}

func NewOrganizationBelongingsRepository(organizations rep.Organizations, id int64) rep.OrganizationBelongings {
	con, err := sql.Open("sqlite3", "./longlong.db")
	if err != nil {
		util.CheckErr(err)
	}
	var organization *entity.Organization
	if id >= 0 {
		organization, err = organizations.Find(id)
		if err != nil {
			util.CheckErr(err)
		}
	}

	return &OrganizationBelongings{
		organization: organization,
		DB:           con,
	}
}

func (o OrganizationBelongings) Entry(individualId int64) error {
	//TODO implement me
	panic("implement me")
}

func (o OrganizationBelongings) Leave(individualId int64, reason string) error {
	//TODO implement me
	panic("implement me")
}

func (o OrganizationBelongings) Members() (*[]entity.Individual, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrganizationBelongings) IndividualsAssigned(individuals *[]entity.Individual) (*[]entity.OrganizationBelonging, error) {
	ids := make([]int64, len(*individuals))
	for i, individual := range *individuals {
		ids[i] = individual.Id
	}

	stmt, err := o.DB.Prepare("select t1.id, t1.parent_id, t1.name, t.individual_id from organization_belongings t join organizations t1 on t.organization_id=t1.id where individual_id in ?")
	if err != nil {
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}(stmt)
	rows, err := stmt.Query(ids)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(fmt.Sprintf("individual ids %d are nothing", ids))
		} else {
			return nil, err
		}
	}

	belongings := make([]entity.OrganizationBelonging, len(*individuals))
	for rows.Next() {
		var id int64
		var parentId int64
		var name string
		var individualId int64
		err = rows.Scan(&id, &parentId, &name, &individualId)
		if err != nil {
			return nil, err
		}
		organization := entity.NewOrganization(id, parentId, name)
		for i, individual := range *individuals {
			if individual.Id == individualId {
				belongings[i] = entity.OrganizationBelonging{
					Individual:   &individual,
					Organization: organization,
				}
			}
		}
	}

	return &belongings, nil
}

func (o OrganizationBelongings) Close() {
	err := o.DB.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
