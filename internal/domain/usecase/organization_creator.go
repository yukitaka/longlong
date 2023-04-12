package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/repository"
)

type OrganizationCreator struct {
	repository.Organizations
	repository.OrganizationBelongings
}

func NewOrganizationCreator(organizations repository.Organizations, belongings repository.OrganizationBelongings) *OrganizationCreator {
	return &OrganizationCreator{organizations, belongings}
}

func (it *OrganizationCreator) New(name string, individual entity.Individual) (int64, error) {
	id, err := it.Organizations.Create(name, individual)
	if err != nil {
		return -1, err
	}

	return id, nil
}
