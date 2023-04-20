package usecase

import (
	"fmt"
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type Authentication struct {
	repository.Authentications
	repository.Organizations
	repository.OrganizationBelongings
}

func NewAuthentication(authentications repository.Authentications, organizations repository.Organizations, organizationBelongings repository.OrganizationBelongings) *Authentication {
	return &Authentication{authentications, organizations, organizationBelongings}
}

func (it *Authentication) Auth(organization, identify, password string) (int64, error) {
	id, token, err := it.Authentications.FindToken(identify)
	if err != nil {
		return -1, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(token), []byte(password))
	if err != nil {
		return -1, err
	}

	organizationBelongings, err := it.OrganizationBelongings.IndividualsAssigned(&[]entity.Individual{*entity.NewIndividual(id, 0, 0, identify)})
	for _, ob := range *organizationBelongings {
		o, _ := it.Organizations.Find(ob.Organization.Id)
		if o.Name == organization {
			return id, nil
		}
	}

	return -1, fmt.Errorf("Error: organization %s not allowed", organization)
}
