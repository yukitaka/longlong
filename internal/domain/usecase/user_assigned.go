package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/repository"
)

type UserAssigned struct {
	repository.Individuals
	repository.Organizations
	repository.OrganizationMembers
}

func NewUserAssigned(individuals repository.Individuals, organizations repository.Organizations, members repository.OrganizationMembers) *UserAssigned {
	return &UserAssigned{individuals, organizations, members}
}

func (it *UserAssigned) OrganizationList(operator *entity.OrganizationMember) (*[]entity.OrganizationMember, error) {
	individuals, err := it.Individuals.FindByUserId(operator.Individual.UserId)
	if err != nil {
		return nil, err
	}

	assigned, err := it.OrganizationMembers.IndividualsAssigned(individuals)
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
	for _, o := range *organizations {
		for i, v := range *assigned {
			if o.Id == v.Organization.Id {
				(*assigned)[i].Organization = &o
			}
		}
	}

	return assigned, nil
}
