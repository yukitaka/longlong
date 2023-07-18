package entity

import "github.com/yukitaka/longlong/server/core/pkg/domain/value_object"

type UserIdentify struct {
	IndividualId   int
	OrganizationId int
	value_object.Role
}

func NewUserIdentify(individual Individual, organization Organization, role value_object.Role) *UserIdentify {
	return &UserIdentify{individual.Id, organization.Id, role}
}
