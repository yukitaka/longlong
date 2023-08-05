package usecase

import (
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"github.com/yukitaka/longlong/server/core/pkg/domain/repository"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
	rep "github.com/yukitaka/longlong/server/core/pkg/interface/repository"
)

type OrganizationCreator struct {
	repository.Organizations
	repository.OrganizationMembers
}

func NewOrganizationCreator(con *datastore.Connection) *OrganizationCreator {
	return &OrganizationCreator{
		Organizations:       rep.NewOrganizationsRepository(con),
		OrganizationMembers: rep.NewOrganizationMembersRepository(con),
	}
}

func (it *OrganizationCreator) New(name string, individual entity.Individual) (int, error) {
	id, err := it.Organizations.Create(name, individual)
	if err != nil {
		return -1, err
	}

	return id, nil
}
