package usecase

import (
	"fmt"
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationRepository struct {
	repository.Authentications
	repository.Organizations
	repository.OrganizationMembers
}

func NewAuthenticationRepository(authentications repository.Authentications, organizations repository.Organizations, organizationMembers repository.OrganizationMembers) *AuthenticationRepository {
	return &AuthenticationRepository{authentications, organizations, organizationMembers}
}

func (rep *AuthenticationRepository) Close() {
	rep.Authentications.Close()
	rep.Organizations.Close()
	rep.OrganizationMembers.Close()
}

type Authentication struct {
	repository *AuthenticationRepository
}

func NewAuthentication(repository *AuthenticationRepository) *Authentication {
	return &Authentication{repository}
}

func (it *Authentication) Auth(organization, identify, password string) (int, error) {
	id, token, err := it.repository.Authentications.FindToken(identify)
	if err != nil {
		return -1, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(token), []byte(password))
	if err != nil {
		return -1, err
	}

	organizationMembers, err := it.repository.OrganizationMembers.IndividualsAssigned(&[]entity.Individual{*entity.NewIndividual(id, &entity.User{}, &entity.Profile{}, identify)})
	for _, ob := range *organizationMembers {
		o, _ := it.repository.Organizations.Find(ob.Organization.Id)
		if o.Name == organization {
			return id, nil
		}
	}

	return -1, fmt.Errorf("Error: organization %s not allowed", organization)
}
