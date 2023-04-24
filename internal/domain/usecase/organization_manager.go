package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/repository"
)

type OrganizationManager struct {
	id int64
	repository.Organizations
	repository.OrganizationBelongings
	repository.Individuals
}

func NewOrganizationManager(id int64, organizations repository.Organizations, organizationBelongings repository.OrganizationBelongings, individuals repository.Individuals) *OrganizationManager {
	return &OrganizationManager{id, organizations, organizationBelongings, individuals}
}

func (it *OrganizationManager) AssignIndividual(individualId int64) error {
	return it.OrganizationBelongings.Entry(individualId)
}

func (it *OrganizationManager) RejectIndividual(individualId int64, reason string) error {
	return it.OrganizationBelongings.Leave(individualId, reason)
}

func (it *OrganizationManager) Members() (*[]entity.Individual, error) {
	return it.OrganizationBelongings.Members(it.id, it.Individuals)
}
