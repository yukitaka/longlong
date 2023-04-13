package repository

import (
	"database/sql"
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

func (o OrganizationBelongings) UserAssigned(userId int64) (*[]entity.OrganizationBelonging, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrganizationBelongings) Close() {
	err := o.DB.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
