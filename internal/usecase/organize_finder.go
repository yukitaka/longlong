package usecase

import (
	"github.com/yukitaka/longlong/internal/entity"
	"github.com/yukitaka/longlong/internal/usecase/repository"
)

type OrganizeFinder struct {
	repository.Organizes
}

func NewOrganizeFinder(organizes repository.Organizes) *OrganizeFinder {
	return &OrganizeFinder{Organizes: organizes}
}

func (it *OrganizeFinder) FindById(id int) (*entity.Organize, error) {
	return it.Organizes.Find(id)
}
