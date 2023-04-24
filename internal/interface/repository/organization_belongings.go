package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/yukitaka/longlong/internal/domain/entity"
	rep "github.com/yukitaka/longlong/internal/domain/repository"
	"github.com/yukitaka/longlong/internal/util"
	"strings"
)

type OrganizationBelongings struct {
	*sql.DB
}

func NewOrganizationBelongingsRepository() rep.OrganizationBelongings {
	con, err := sql.Open("sqlite3", "./longlong.db")
	if err != nil {
		util.CheckErr(err)
	}

	return &OrganizationBelongings{
		DB: con,
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

func (o OrganizationBelongings) Members(organizationId int64) (*[]entity.Individual, error) {
	stmt := "select organization_id, individual_id from organization_belongings where organization_id=?"
	ret, err := o.DB.Query(stmt, organizationId)
	if err != nil {
		return nil, err
	}

	var individuals []entity.Individual
	for ret.Next() {
		var oid int64
		var iid int64
		err := ret.Scan(&oid, &iid)
		if err != nil {
			return nil, err
		}
		individuals = append(individuals, *entity.NewIndividual(iid, -1, -1, ""))
	}

	return &individuals, nil
}

func (o OrganizationBelongings) IndividualsAssigned(individuals *[]entity.Individual) (*[]entity.OrganizationBelonging, error) {
	ids := make([]interface{}, len(*individuals))
	for i, individual := range *individuals {
		ids[i] = individual.Id
	}

	stmt := "select t1.id, t1.parent_id, t1.name, t.individual_id from organization_belongings t join organizations t1 on t.organization_id=t1.id where individual_id in (?" + strings.Repeat(",?", len(ids)-1) + ")"
	rows, err := o.DB.Query(stmt, ids...)
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
		organization := entity.NewOrganization(parentId, id, name)
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
