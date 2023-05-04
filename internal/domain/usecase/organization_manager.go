package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/repository"
	"github.com/yukitaka/longlong/internal/domain/value_object"
)

type OrganizationManager struct {
	organization *entity.Organization
	repository.Organizations
	repository.OrganizationMembers
	repository.Individuals
}

func NewOrganizationManager(organization *entity.Organization, organizations repository.Organizations, organizationBelongings repository.OrganizationMembers, individuals repository.Individuals) *OrganizationManager {
	return &OrganizationManager{organization, organizations, organizationBelongings, individuals}
}

func (it *OrganizationManager) AssignIndividual(individualId int64) error {
	return it.OrganizationMembers.Entry(it.organization.Id, individualId, value_object.MEMBER)
}

func (it *OrganizationManager) RejectIndividual(individualId int64, reason string) error {
	return it.OrganizationMembers.Leave(individualId, reason)
}

func (it *OrganizationManager) Members() (*[]entity.OrganizationMember, error) {
	return it.OrganizationMembers.Members(it.organization, it.Individuals)
}
