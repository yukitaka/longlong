package usecase

import (
	"fmt"
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"github.com/yukitaka/longlong/server/core/pkg/domain/repository"
	"github.com/yukitaka/longlong/server/core/pkg/domain/value_object"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
	rep "github.com/yukitaka/longlong/server/core/pkg/interface/repository"
)

type OrganizationManager struct {
	organization *entity.Organization
	repository.Organizations
	repository.OrganizationMembers
	repository.Individuals
}

func NewOrganizationManager(org *entity.Organization, con *datastore.Connection) *OrganizationManager {
	return &OrganizationManager{
		organization:        org,
		Organizations:       rep.NewOrganizationsRepository(con),
		OrganizationMembers: rep.NewOrganizationMembersRepository(con),
		Individuals:         rep.NewIndividualsRepository(con),
	}
}

func (it *OrganizationManager) AssignIndividual(individualId int) error {
	return it.OrganizationMembers.Entry(it.organization.Id, individualId, value_object.MEMBER)
}

func (it *OrganizationManager) QuitIndividual(operator *entity.OrganizationMember, individualId int, reason string) error {
	target, err := it.OrganizationMembers.Find(it.organization.Id, individualId)
	if err != nil {
		return err
	}
	if !operator.CanManage(target) {
		return fmt.Errorf("error: you don't have permission to quit this organization")
	}

	return it.OrganizationMembers.Leave(it.organization.Id, individualId, reason)
}

func (it *OrganizationManager) Members() (*[]entity.OrganizationMember, error) {
	return it.OrganizationMembers.Members(it.organization, it.Individuals)
}
