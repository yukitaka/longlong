package usecase

import (
	"fmt"
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"github.com/yukitaka/longlong/server/core/pkg/domain/repository"
	"github.com/yukitaka/longlong/server/core/pkg/domain/value_object"
)

type OrganizationManagerRepository struct {
	repository.Organizations
	repository.OrganizationMembers
	repository.Individuals
}

func NewOrganizationManagerRepository(organizations repository.Organizations, organizationMembers repository.OrganizationMembers, individuals repository.Individuals) *OrganizationManagerRepository {
	return &OrganizationManagerRepository{organizations, organizationMembers, individuals}
}

func (rep *OrganizationManagerRepository) Close() {
	rep.Organizations.Close()
	rep.OrganizationMembers.Close()
	rep.Individuals.Close()
}

type OrganizationManager struct {
	organization *entity.Organization
	repository   *OrganizationManagerRepository
}

func NewOrganizationManager(organization *entity.Organization, repository *OrganizationManagerRepository) *OrganizationManager {
	return &OrganizationManager{organization, repository}
}

func (it *OrganizationManager) AssignIndividual(individualId int) error {
	return it.repository.OrganizationMembers.Entry(it.organization.Id, individualId, value_object.MEMBER)
}

func (it *OrganizationManager) QuitIndividual(operator *entity.OrganizationMember, individualId int, reason string) error {
	target, err := it.repository.OrganizationMembers.Find(it.organization.Id, individualId)
	if err != nil {
		return err
	}
	if !operator.CanManage(target) {
		return fmt.Errorf("error: you don't have permission to quit this organization")
	}

	return it.repository.OrganizationMembers.Leave(it.organization.Id, individualId, reason)
}

func (it *OrganizationManager) Members() (*[]entity.OrganizationMember, error) {
	return it.repository.OrganizationMembers.Members(it.organization, it.repository.Individuals)
}
