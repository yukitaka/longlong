package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/repository"
	"github.com/yukitaka/longlong/internal/domain/value_object"
)

type OrganizationManager struct {
	organization *entity.Organization
	repository.Organizations
	repository.OrganizationBelongings
	repository.Individuals
}

func NewOrganizationManager(organization *entity.Organization, organizations repository.Organizations, organizationBelongings repository.OrganizationBelongings, individuals repository.Individuals) *OrganizationManager {
	return &OrganizationManager{organization, organizations, organizationBelongings, individuals}
}

func (it *OrganizationManager) AssignIndividual(individualId int64) error {
	return it.OrganizationBelongings.Entry(it.organization.Id, individualId, value_object.MEMBER)
}

func (it *OrganizationManager) RejectIndividual(individualId int64, reason string) error {
	return it.OrganizationBelongings.Leave(individualId, reason)
}

func (it *OrganizationManager) Members() (*[]entity.OrganizationBelonging, error) {
	return it.OrganizationBelongings.Members(it.organization, it.Individuals)
}
