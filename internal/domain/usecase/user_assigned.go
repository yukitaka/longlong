package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/repository"
)

type UserAssigned struct {
	UserId int64
	repository.Individuals
	repository.Organizations
	repository.OrganizationBelongings
}

func NewUserAssigned(userId int64, individuals repository.Individuals, organizations repository.Organizations, belongings repository.OrganizationBelongings) *UserAssigned {
	return &UserAssigned{userId, individuals, organizations, belongings}
}

func (it *UserAssigned) OrganizationList() (*[]entity.Organization, error) {
	individuals, err := it.Individuals.FindByUserId(it.UserId)
	if err != nil {
		return nil, err
	}

	assigned, err := it.OrganizationBelongings.IndividualsAssigned(individuals)
	if err != nil {
		return nil, err
	}
	organizationIds := make([]interface{}, len(*assigned))
	for i, v := range *assigned {
		organizationIds[i] = v.Organization.Id
	}
	organizations, err := it.Organizations.FindAll(organizationIds)
	if err != nil {
		return nil, err
	}
	return organizations, nil
}
