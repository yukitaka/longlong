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

func NewOrganizationManager(organization *entity.Organization, organizations repository.Organizations, organizationMembers repository.OrganizationMembers, individuals repository.Individuals) *OrganizationManager {
	return &OrganizationManager{organization, organizations, organizationMembers, individuals}
}

func (it *OrganizationManager) AssignIndividual(individualId int) error {
	return it.OrganizationMembers.Entry(it.organization.Id, individualId, value_object.MEMBER)
}

func (it *OrganizationManager) RejectIndividual(individualId int, reason string) error {
	return it.OrganizationMembers.Leave(it.organization.Id, individualId, reason)
}

func (it *OrganizationManager) Members() (*[]entity.OrganizationMember, error) {
	return it.OrganizationMembers.Members(it.organization, it.Individuals)
}
