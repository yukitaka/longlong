package usecase

import (
	"fmt"
	"github.com/yukitaka/longlong/internal/domain/repository"
	"github.com/yukitaka/longlong/internal/domain/value_object"
	"strings"
)

type UserCreator struct {
	repository.Users
	repository.Individuals
	repository.OrganizationMembers
}

func NewUserCreator(users repository.Users, individuals repository.Individuals, belongings repository.OrganizationMembers) *UserCreator {
	return &UserCreator{users, individuals, belongings}
}

func (it *UserCreator) New(operatorId, organizationId int64, name string, role string) (int64, error) {
	operator, err := it.OrganizationMembers.Find(organizationId, operatorId)
	roleType, err := value_object.ParseRole(strings.ToUpper(role))
	if err != nil {
		return 0, err
	}
	if operator.Role > roleType {
		return -1, fmt.Errorf("New user role isn't permitted.\n")
	}

	userId, err := it.Users.Create(name)
	if err != nil {
		return -1, err
	}
	id, err := it.Individuals.Create(name, userId, -1)
	if err != nil {
		return 0, err
	}
	err = it.OrganizationMembers.Entry(organizationId, id, roleType)

	return id, nil
}
