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

func (it *UserCreator) New(organizationId int64, name string, role string) (int64, error) {
	var roleType value_object.Role
	if role == "owner" {
		roleType = value_object.OWNER
	} else if role == "admin" {
		roleType = value_object.ADMIN
	} else {
		roleType = value_object.MEMBER
	}

	userId, err := it.Users.Create(name)
	if err != nil {
		return -1, err
	}
	id, err := it.Individuals.Create(name, userId, -1)
	if err != nil {
		return 0, err
	}
	err = it.OrganizationBelongings.Entry(organizationId, id, roleType)

	return id, nil
}
