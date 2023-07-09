package usecase

import (
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"github.com/yukitaka/longlong/server/core/pkg/domain/repository"
)

type UserAssignedRepository struct {
	repository.Individuals
	repository.Organizations
	repository.OrganizationMembers
}

func NewUserAssignedRepository(individuals repository.Individuals, organizations repository.Organizations, members repository.OrganizationMembers) *UserAssignedRepository {
	return &UserAssignedRepository{individuals, organizations, members}
}

func (rep *UserAssignedRepository) Close() {
	rep.Individuals.Close()
	rep.Organizations.Close()
	rep.OrganizationMembers.Close()
}

type UserAssigned struct {
	repository *UserAssignedRepository
}

func NewUserAssigned(repository *UserAssignedRepository) *UserAssigned {
	return &UserAssigned{repository}
}

func (it *UserAssigned) OrganizationList(operator *entity.OrganizationMember) (*[]entity.OrganizationMember, error) {
	individuals, err := it.repository.Individuals.FindByUserId(operator.Individual.User.Id)
	if err != nil {
		return nil, err
	}

	assigned, err := it.repository.OrganizationMembers.IndividualsAssigned(individuals)
	if err != nil {
		return nil, err
	}
	organizationIds := make([]interface{}, len(*assigned))
	for i, v := range *assigned {
		organizationIds[i] = v.Organization.Id
	}
	organizations, err := it.repository.Organizations.FindAll(organizationIds)
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
