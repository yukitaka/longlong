package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/repository"
)

type UserCreator struct {
	repository.Users
}

func NewUserCreator(users repository.Users) *UserCreator {
	return &UserCreator{users}
}

func (it *UserCreator) New(name string) (int64, error) {
	id, err := it.Users.Create(name)
	if err != nil {
		return -1, err
	}

	return id, nil
}
