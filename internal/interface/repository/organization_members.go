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

type OrganizationMembers struct {
	*sql.DB
}

func NewOrganizationMembersRepository() rep.OrganizationMembers {
	con, err := sql.Open("sqlite3", "./longlong.db")
	if err != nil {
		util.CheckErr(err)
	}

	return &OrganizationMembers{
		DB: con,
	}
}

func (o OrganizationMembers) Find(organizationId, individualId int64) (*entity.OrganizationMember, error) {
	query := "select role from organization_members where organization_id=$1 and individual_id=$2"
	row := o.DB.QueryRow(query, organizationId, individualId)

	var role int
	if err := row.Scan(&role); err != nil {
		return nil, err
	}
	roleType := value_object.Role(role)

	var parentId int64
	var organizationName string
	row = o.DB.QueryRow("select parent_id, name from organizations where id=$1", organizationId)
	if err := row.Scan(&parentId, &organizationName); err != nil {
		return nil, err
	}
	organization := entity.NewOrganization(parentId, organizationId, organizationName)

	var userId int64
	var profileId int64
	var individualName string
	row = o.DB.QueryRow("select user_id, profile_id, name from individuals where id=$1", individualId)
	if err := row.Scan(&userId, &profileId, &individualName); err != nil {
		return nil, err
	}
	individual := entity.NewIndividual(individualId, userId, profileId, individualName)

	return entity.NewOrganizationMember(organization, individual, roleType), nil
}

func (o OrganizationMembers) Entry(organizationId, individualId int64, role value_object.Role) error {
	query := "insert into organization_members (organization_id, individual_id, role) values (?, ?, ?)"
	_, err := o.DB.Exec(query, organizationId, individualId, role)

	return err
}

func (o OrganizationMembers) Leave(individualId int64, reason string) error {
	//TODO implement me
	panic("implement me")
}

func (o OrganizationMembers) Members(organization *entity.Organization, individualRepository rep.Individuals) (*[]entity.OrganizationMember, error) {
	stmt := "select organization_id, individual_id, role from organization_members where organization_id=?"
	ret, err := o.DB.Query(stmt, organization.Id)
	if err != nil {
		return nil, err
	}

	var members []entity.OrganizationMember
	for ret.Next() {
		var oid int64
		var iid int64
		var role int
		err := ret.Scan(&oid, &iid, &role)
		if err != nil {
			return nil, err
		}
		individual, err := individualRepository.Find(iid)
		if err != nil {
			return nil, err
		}
		members = append(members, *entity.NewOrganizationMember(organization, individual, value_object.Role(role)))
	}

	return &members, nil
}

func (o OrganizationMembers) IndividualsAssigned(individuals *[]entity.Individual) (*[]entity.OrganizationMember, error) {
	ids := make([]interface{}, len(*individuals))
	for i, individual := range *individuals {
		ids[i] = individual.Id
	}

	stmt := "select t1.id, t1.parent_id, t1.name, t.individual_id from organization_members t join organizations t1 on t.organization_id=t1.id where individual_id in (?" + strings.Repeat(",?", len(ids)-1) + ")"
	rows, err := o.DB.Query(stmt, ids...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(fmt.Sprintf("individual ids %d are nothing", ids))
		} else {
			return nil, err
		}
	}

	members := make([]entity.OrganizationMember, len(*individuals))
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
				members[i] = entity.OrganizationMember{
					Individual:   &individual,
					Organization: organization,
				}
			}
		}
	}

	return &members, nil
}

func (o OrganizationMembers) Close() {
	err := o.DB.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
