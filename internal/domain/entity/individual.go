package entity

import "github.com/yukitaka/longlong/internal/domain/repository"

type Individual struct {
	Id        int64
	Name      string
	UserId    int64
	ProfileId int64
}

func NewIndividual(id, userId, profileId int64, name string) *Individual {
	return &Individual{Id: id, UserId: userId, ProfileId: profileId, Name: name}
}

func (it *Individual) Organizations(organizationsRep repository.Organizations, organization_belongingsRep repository.OrganizationBelongings) (*[]Organization, error) {
	assigned, err := organization_belongingsRep.UserAssigned(it.UserId)
	if err != nil {
		return nil, err
	}
	organizationIds := make([]int64, len(*assigned))
	for i, v := range *assigned {
		organizationIds[i] = v.OrganizationId
	}
	organizations, err := organizationsRep.FindAll(organizationIds)
	if err != nil {
		return nil, err
	}
	return organizations, nil
}
