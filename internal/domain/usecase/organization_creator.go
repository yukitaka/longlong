package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/repository"
)

type OrganizationCreator struct {
	repository.Organizations
	repository.OrganizationMembers
}

func NewOrganizationCreator(organizations repository.Organizations, members repository.OrganizationMembers) *OrganizationCreator {
	return &OrganizationCreator{organizations, members}
}

func (it *OrganizationCreator) New(name string, individual entity.Individual) (int, error) {
	id, err := it.Organizations.Create(name, individual)
	if err != nil {
		return -1, err
	}

	return id, nil
}
