package usecase

import (
	"errors"
	"fmt"
	"github.com/yukitaka/longlong/internal/usecase/repository"
)

type OrganizeCreator struct {
	repository.Organizes
}

func NewOrganizeCreator(organizes repository.Organizes) *OrganizeCreator {
	return &OrganizeCreator{organizes}
}

func (it *OrganizeCreator) New(name string) (int64, error) {
	if id := it.Organizes.Create(name); id > 0 {
		return id, nil
	}

	return 0, errors.New(fmt.Sprintf("Error: Create %s", name))
}
