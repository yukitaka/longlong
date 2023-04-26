package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/repository"
	"github.com/yukitaka/longlong/internal/domain/value_object"
)

type UserCreator struct {
	repository.Users
	repository.Individuals
	repository.OrganizationBelongings
}

func NewUserCreator(users repository.Users, individuals repository.Individuals, belongings repository.OrganizationBelongings) *UserCreator {
	return &UserCreator{users, individuals, belongings}
}

func (it *UserCreator) New(organizationId int64, name string) (int64, error) {
	userId, err := it.Users.Create(name)
	if err != nil {
		return -1, err
	}
	id, err := it.Individuals.Create("Default", userId, -1)
	if err != nil {
		return 0, err
	}
	err = it.OrganizationBelongings.Entry(organizationId, id, value_object.MEMBER)

	return id, nil
}
