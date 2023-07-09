package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	rep "github.com/yukitaka/longlong/server/core/pkg/domain/repository"
	"github.com/yukitaka/longlong/server/core/pkg/domain/value_object"
	"strings"
)

type OrganizationMembers struct {
	*sqlx.DB
}

func NewOrganizationMembersRepository(con *sqlx.DB) rep.OrganizationMembers {
	return &OrganizationMembers{
		DB: con,
	}
}

func (rep OrganizationMembers) Find(organizationId, individualId int) (*entity.OrganizationMember, error) {
	query := "select role from organization_members where organization_id=$1 and individual_id=$2"
	row := rep.DB.QueryRowx(query, organizationId, individualId)

	var role int
	if err := row.Scan(&role); err != nil {
		return nil, err
	}
	roleType := value_object.Role(role)

	var parentId int
	var organizationName string
	row = rep.DB.QueryRowx("select parent_id, name from organizations where id=$1", organizationId)
	if err := row.Scan(&parentId, &organizationName); err != nil {
		return nil, err
	}
	organization := entity.NewOrganization(parentId, organizationId, organizationName)

	var userId int
	var profileId int
	var individualName string
	row = rep.DB.QueryRowx("select user_id, profile_id, name from individuals where id=$1", individualId)
	if err := row.Scan(&userId, &profileId, &individualName); err != nil {
		return nil, err
	}
	user, err := NewUsersRepository(rep.DB).Find(userId)
	if err != nil {
		return nil, err
	}
	profile, err := NewProfilesRepository(rep.DB).Find(profileId)
	if err != nil {
		return nil, err
	}

	individual := entity.NewIndividual(individualId, user, profile, individualName)

	return entity.NewOrganizationMember(organization, individual, roleType), nil
}

func (rep OrganizationMembers) Entry(organizationId, individualId int, role value_object.Role) error {
	query := "insert into organization_members (organization_id, individual_id, role) values ($1, $2, $3)"
	_, err := rep.DB.Exec(query, organizationId, individualId, role)

	return err
}

func (rep OrganizationMembers) Leave(organizationId, individualId int, reason string) error {
	stmt := "delete from organization_members where organization_id=$1 and individual_id=$2"
	_, err := rep.DB.Exec(stmt, organizationId, individualId)

	return err
}

func (rep OrganizationMembers) Members(organization *entity.Organization, individualRepository rep.Individuals) (*[]entity.OrganizationMember, error) {
	stmt := "select organization_id, individual_id, role from organization_members where organization_id=$1"
	ret, err := rep.DB.Queryx(stmt, organization.Id)
	if err != nil {
		return nil, err
	}

	var members []entity.OrganizationMember
	for ret.Next() {
		var oid int
		var iid int
		var role int
		err := ret.Scan(&oid, &iid, &role)
		if err != nil {
			return nil, err
		}
		individual, err := individualRepository.Find(iid)
		if err != nil {
			continue
		}
		members = append(members, *entity.NewOrganizationMember(organization, individual, value_object.Role(role)))
	}

	return &members, nil
}

func (rep OrganizationMembers) IndividualsAssigned(individuals *[]entity.Individual) (*[]entity.OrganizationMember, error) {
	ids := make([]interface{}, len(*individuals))
	for i, individual := range *individuals {
		ids[i] = individual.Id
	}

	stmt := "select t1.id, t1.parent_id, t1.name, t.individual_id from organization_members t join organizations t1 on t.organization_id=t1.id where individual_id in ($1" + strings.Repeat(",$2", len(ids)-1) + ")"
	rows, err := rep.DB.Queryx(stmt, ids...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(fmt.Sprintf("individual ids %d are nothing", ids))
		} else {
			return nil, err
		}
	}

	members := make([]entity.OrganizationMember, len(*individuals))
	for rows.Next() {
		var id int
		var parentId int
		var name string
		var individualId int
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

func (rep OrganizationMembers) Close() {
	err := rep.DB.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
