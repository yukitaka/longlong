package usecase

import (
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"github.com/yukitaka/longlong/server/core/pkg/domain/repository"
)

type OrganizationCreatorRepository struct {
	repository.Organizations
	repository.OrganizationMembers
}

func NewOrganizationCreatorRepository(organizations repository.Organizations, members repository.OrganizationMembers) *OrganizationCreatorRepository {
	return &OrganizationCreatorRepository{organizations, members}
}

func (rep *OrganizationCreatorRepository) Close() {
	rep.Organizations.Close()
	rep.OrganizationMembers.Close()
}

type OrganizationCreator struct {
	repository *OrganizationCreatorRepository
}

func NewOrganizationCreator(repository *OrganizationCreatorRepository) *OrganizationCreator {
	return &OrganizationCreator{repository}
}

func (it *OrganizationCreator) New(name string, individual entity.Individual) (int, error) {
	id, err := it.repository.Organizations.Create(name, individual)
	if err != nil {
		return -1, err
	}

	return id, nil
}
