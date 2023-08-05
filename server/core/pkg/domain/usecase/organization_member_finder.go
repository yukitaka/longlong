package usecase

import (
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"github.com/yukitaka/longlong/server/core/pkg/domain/repository"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
	rep "github.com/yukitaka/longlong/server/core/pkg/interface/repository"
)

type OrganizationMemberFinder struct {
	repository.OrganizationMembers
}

func NewOrganizationMemberFinder(con *datastore.Connection) *OrganizationMemberFinder {
	return &OrganizationMemberFinder{OrganizationMembers: rep.NewOrganizationMembersRepository(con)}
}

func (it *OrganizationMemberFinder) FindById(organizationId, individualId int) (*entity.OrganizationMember, error) {
	return it.OrganizationMembers.Find(organizationId, individualId)
}
