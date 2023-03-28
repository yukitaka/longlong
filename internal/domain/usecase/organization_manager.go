package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
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

func (it *OrganizationManager) AssignAvatar(avatarId int64) error {
	return it.OrganizationBelongings.Entry(it.id, avatarId)
}

func (it *OrganizationManager) RejectAvatar(avatarId int64, reason string) error {
	return it.OrganizationBelongings.Leave(it.id, avatarId, reason)
}

func (it *OrganizationManager) Members() (*[]entity.Avatar, error) {
	return it.OrganizationBelongings.Members(it.id)
}
