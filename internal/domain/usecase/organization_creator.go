package usecase

import (
	"errors"
	"fmt"
	"github.com/yukitaka/longlong/internal/domain/usecase/repository"
)

type OrganizationCreator struct {
	repository.Organizations
}

func NewOrganizationCreator(organizations repository.Organizations) *OrganizationCreator {
	return &OrganizationCreator{organizations}
}

func (it *OrganizationCreator) New(name string) (int, error) {
	if id := it.Organizations.Create(name); id > 0 {
		return id, nil
	}

	return 0, errors.New(fmt.Sprintf("Error: Create %s", name))
}
