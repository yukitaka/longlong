package usecase

import (
	"fmt"
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"github.com/yukitaka/longlong/server/core/pkg/domain/repository"
	"github.com/yukitaka/longlong/server/core/pkg/domain/value_object"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
	rep "github.com/yukitaka/longlong/server/core/pkg/interface/repository"
	"strings"
)

type UserCreator struct {
	repository.Users
	repository.Individuals
	repository.OrganizationMembers
}

func NewUserCreator(con *datastore.Connection) *UserCreator {
	return &UserCreator{
		Users:               rep.NewUsersRepository(con),
		Individuals:         rep.NewIndividualsRepository(con),
		OrganizationMembers: rep.NewOrganizationMembersRepository(con),
	}
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
