package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type Authentication struct {
	repository.Authentications
	repository.OrganizationBelongings
}

func NewAuthentication(authentications repository.Authentications, organizationBelongings repository.OrganizationBelongings) *Authentication {
	return &Authentication{authentications, organizationBelongings}
}

func (it *Authentication) Auth(identify, password string) (int64, error) {
	id, token, err := it.Authentications.FindToken(identify)
	if err != nil {
		return -1, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(token), []byte(password))
	if err != nil {
		return -1, err
	}

	return id, nil
}
