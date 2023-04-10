package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/repository"
)

type UserCreator struct {
	repository.Users
	repository.Avatars
}

func NewUserCreator(users repository.Users, avatars repository.Avatars) *UserCreator {
	return &UserCreator{users, avatars}
}

func (it *UserCreator) New(name string) (int64, error) {
	userId, err := it.Users.Create(name)
	if err != nil {
		return -1, err
	}
	id, err := it.Avatars.Create("Default", userId, -1)
	if err != nil {
		return 0, err
	}

	return id, nil
}
