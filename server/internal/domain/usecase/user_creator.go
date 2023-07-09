package usecase

import (
	"fmt"
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/repository"
	"github.com/yukitaka/longlong/internal/domain/value_object"
	"strings"
)

type UserCreatorRepository struct {
	repository.Users
	repository.Individuals
	repository.OrganizationMembers
}

func NewUserCreatorRepository(users repository.Users, individuals repository.Individuals, members repository.OrganizationMembers) *UserCreatorRepository {
	return &UserCreatorRepository{users, individuals, members}
}

func (rep *UserCreatorRepository) Close() {
	rep.Users.Close()
	rep.Individuals.Close()
	rep.OrganizationMembers.Close()
}

type UserCreator struct {
	repository *UserCreatorRepository
}

func NewUserCreator(repository *UserCreatorRepository) *UserCreator {
	return &UserCreator{repository}
}

func (it *UserCreator) New(operator *entity.OrganizationMember, name string, role string) (int, error) {
	roleType, err := value_object.ParseRole(strings.ToUpper(role))
	if err != nil {
		return 0, err
	}
	if operator.Role.IsBelow(roleType) {
		return -1, fmt.Errorf("New user role isn't permitted.\n")
	}

	userId, err := it.repository.Users.Create(name)
	if err != nil {
		return -1, err
	}
	id, err := it.repository.Individuals.Create(name, userId, -1)
	if err != nil {
		return 0, err
	}
	err = it.repository.OrganizationMembers.Entry(operator.Organization.Id, id, roleType)

	return id, nil
}
