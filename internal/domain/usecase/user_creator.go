package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/repository"
)

type UserCreator struct {
	repository.Users
	repository.Individuals
}

func NewUserCreator(users repository.Users, individuals repository.Individuals) *UserCreator {
	return &UserCreator{users, individuals}
}

func (it *UserCreator) New(organizationId int64, name string) (int64, error) {
	userId, err := it.Users.Create(organizationId, name)
	if err != nil {
		return -1, err
	}
	id, err := it.Individuals.Create("Default", userId, -1)
	if err != nil {
		return 0, err
	}

	return id, nil
}
