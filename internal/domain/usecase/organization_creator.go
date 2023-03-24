package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/repository"
)

type OrganizationCreator struct {
	repository.Organizations
}

func NewOrganizationCreator(organizations repository.Organizations) *OrganizationCreator {
	return &OrganizationCreator{organizations}
}

func (it *OrganizationCreator) New(name string) (int64, error) {
	id, err := it.Organizations.Create(name)
	if err != nil {
		return -1, err
	}

	return id, nil
}
