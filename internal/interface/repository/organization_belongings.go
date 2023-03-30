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

func NewOrganizationBelongingsRepository(id int64, organizations rep.Organizations) rep.OrganizationBelongings {
	con, err := sql.Open("sqlite3", "./longlong.db")
	if err != nil {
		util.CheckErr(err)
	}
	organization, err := organizations.Find(id)
	if err != nil {
		util.CheckErr(err)
	}

	return &OrganizationBelongings{
		organization: organization,
		DB:           con,
	}
}

func (o OrganizationBelongings) Entry(avatarId int64) error {
	//TODO implement me
	panic("implement me")
}

func (o OrganizationBelongings) Leave(avatarId int64, reason string) error {
	//TODO implement me
	panic("implement me")
}

func (o OrganizationBelongings) Members() (*[]entity.Avatar, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrganizationBelongings) Close() {
	err := o.DB.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
