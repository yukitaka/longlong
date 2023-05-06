package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/repository"
)

type OrganizationMemberFinder struct {
	repository.OrganizationMembers
}

func NewOrganizationMemberFinder(members repository.OrganizationMembers) *OrganizationMemberFinder {
	return &OrganizationMemberFinder{OrganizationMembers: members}
}

func (it *OrganizationMemberFinder) FindById(organizationId, individualId int) (*entity.OrganizationMember, error) {
	return it.OrganizationMembers.Find(organizationId, individualId)
}
