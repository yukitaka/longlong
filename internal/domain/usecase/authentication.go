package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/repository"
)

type Authentication struct {
	repository.Authentications
}

func NewAuthentication(authentications repository.Authentications) *Authentication {
	return &Authentication{authentications}
}

func (it *Authentication) Auth(name, token string) (int64, error) {
	id, err := it.Authentications.Auth(name, token)
	if err != nil {
		return -1, err
	}

	return id, nil
}
