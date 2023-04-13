package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/repository"
)

type UserAssigned struct {
	UserId int64
	repository.Organizations
	repository.OrganizationBelongings
}

func NewUserAssigned(userId int64, organizations repository.Organizations, belongings repository.OrganizationBelongings) *UserAssigned {
	return &UserAssigned{userId, organizations, belongings}
}

func (it *UserAssigned) OrganizationList() (*[]entity.Organization, error) {
	assigned, err := it.OrganizationBelongings.UserAssigned(it.UserId)
	if err != nil {
		return nil, err
	}
	organizationIds := make([]int64, len(*assigned))
	for i, v := range *assigned {
		organizationIds[i] = v.OrganizationId
	}
	organizations, err := it.Organizations.FindAll(organizationIds)
	if err != nil {
		return nil, err
	}
	return organizations, nil
}
