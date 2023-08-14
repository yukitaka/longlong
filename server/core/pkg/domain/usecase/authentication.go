package usecase

import (
	"fmt"
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"github.com/yukitaka/longlong/server/core/pkg/domain/repository"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
	rep "github.com/yukitaka/longlong/server/core/pkg/interface/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Authentication struct {
	repository.Authentications
	repository.Organizations
	repository.OrganizationMembers
}

func NewAuthentication(con *datastore.Connection) *Authentication {
	return &Authentication{
		Authentications:     rep.NewAuthenticationsRepository(con),
		Organizations:       rep.NewOrganizationsRepository(con),
		OrganizationMembers: rep.NewOrganizationMembersRepository(con),
	}
}

func (it *Authentication) FindById(organizationId int, identify string) (*entity.Authentication, error) {
	return it.Authentications.Find(organizationId, identify)
}

func (it *Authentication) Store(organizationId int, identify, token string) (bool, error) {
	return it.Authentications.Store(organizationId, identify, token)
}

func (it *Authentication) StoreOAuth2Info(identify, accessToken, refreshToken string, expiry time.Time) (bool, error) {
	return it.Authentications.StoreOAuth2Info(identify, accessToken, refreshToken, expiry)
}

func (it *Authentication) AuthOAuth(identify, token string) (int, error) {
	id, dbToken, err := it.Authentications.FindToken(-1, identify)
	if err != nil {
		return -1, err
	}
	if token != dbToken {
		err = it.Authentications.UpdateToken(id, token)
		if err != nil {
			return id, err
		}
	}

	return id, nil
}

func (it *Authentication) Auth(organization, identify, password string) (int, int, error) {
	org, err := it.Organizations.FindByName(organization)
	if err != nil {
		return -1, -1, err
	}
	id, token, err := it.Authentications.FindToken(org.Id, identify)
	if err != nil {
		return -1, -1, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(token), []byte(password))
	if err != nil {
		return -1, -1, err
	}

	organizationMembers, err := it.OrganizationMembers.IndividualsAssigned(&[]entity.Individual{*entity.NewIndividual(id, &entity.User{}, &entity.Profile{}, identify)})
	for _, ob := range *organizationMembers {
		o, _ := it.Organizations.Find(ob.Organization.Id)
		if o.Name == organization {
			return id, o.Id, nil
		}
	}

	return -1, -1, fmt.Errorf("Error: organization %s not allowed.", organization)
}
