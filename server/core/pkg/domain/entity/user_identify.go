package entity

import "github.com/yukitaka/longlong/server/core/pkg/domain/value_object"

type UserIdentify struct {
	IndividualId      int `json:"individual_id"`
	OrganizationId    int `json:"organization_id"`
	value_object.Role `json:"role"`
}

func NewUserIdentify(individual Individual, organization Organization, role value_object.Role) *UserIdentify {
	return &UserIdentify{individual.Id, organization.Id, role}
}
