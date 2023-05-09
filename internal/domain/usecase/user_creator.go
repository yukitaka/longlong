package usecase

import (
	"fmt"
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/repository"
	"github.com/yukitaka/longlong/internal/domain/value_object"
	"strings"
)

type UserCreator struct {
	repository.Users
	repository.Individuals
	repository.OrganizationMembers
}

func NewUserCreator(users repository.Users, individuals repository.Individuals, members repository.OrganizationMembers) *UserCreator {
	return &UserCreator{users, individuals, members}
}

func (it *UserCreator) New(operator *entity.OrganizationMember, name string, role string) (int, error) {
	roleType, err := value_object.ParseRole(strings.ToUpper(role))
	if err != nil {
		return 0, err
	}
	if operator.Role.IsBelow(roleType) {
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
	err = it.OrganizationMembers.Entry(operator.Organization.Id, id, roleType)

	return id, nil
}
