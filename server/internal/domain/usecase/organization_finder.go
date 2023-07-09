package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/repository"
)

type OrganizationFinder struct {
	repository.Organizations
}

func NewOrganizationFinder(organizations repository.Organizations) *OrganizationFinder {
	return &OrganizationFinder{Organizations: organizations}
}

func (it *OrganizationFinder) FindById(id int) (*entity.Organization, error) {
	return it.Organizations.Find(id)
}

func (it *OrganizationFinder) All() (*[]entity.Organization, error) {
	return it.Organizations.List()
}
