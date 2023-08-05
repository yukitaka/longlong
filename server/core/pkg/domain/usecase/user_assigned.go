package usecase

import (
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"github.com/yukitaka/longlong/server/core/pkg/domain/repository"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
	rep "github.com/yukitaka/longlong/server/core/pkg/interface/repository"
)

type UserAssigned struct {
	repository.Individuals
	repository.Organizations
	repository.OrganizationMembers
}

func NewUserAssigned(con *datastore.Connection) *UserAssigned {
	return &UserAssigned{
		Individuals:         rep.NewIndividualsRepository(con),
		Organizations:       rep.NewOrganizationsRepository(con),
		OrganizationMembers: rep.NewOrganizationMembersRepository(con),
	}
}

func (it *UserAssigned) OrganizationList(operator *entity.OrganizationMember) (*[]entity.OrganizationMember, error) {
	individuals, err := it.Individuals.FindByUserId(operator.Individual.User.Id)
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
