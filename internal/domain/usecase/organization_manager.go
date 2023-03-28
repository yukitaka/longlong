package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/repository"
)

type OrganizationManager struct {
	id int64
	repository.Organizations
	repository.OrganizationBelongings
	repository.Users
}

func NewOrganizationManager(id int64, organizations repository.Organizations, organizationBelongings repository.OrganizationBelongings, users repository.Users) *OrganizationManager {
	return &OrganizationManager{id, organizations, organizationBelongings, users}
}

func (it *OrganizationManager) AssignUser(userId int64) error {
	return it.OrganizationBelongings.Entry(it.id, userId)
}
