package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type Authentication struct {
	repository.Authentications
}

func NewAuthentication(authentications repository.Authentications) *Authentication {
	return &Authentication{authentications}
}

func (it *Authentication) Auth(name, password string) (int64, error) {
	id, token, err := it.Authentications.FindToken(name)
	if err != nil {
		return -1, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	if token != string(hash) {
		return -1, nil
	}

	return id, nil
}
