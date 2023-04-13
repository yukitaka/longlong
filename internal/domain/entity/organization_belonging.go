package entity

import "github.com/yukitaka/longlong/internal/domain/value_object"

type OrganizationBelonging struct {
	Organization *Organization
	Individual   *Individual
	Role         value_object.Role
}

func NewOrganizationBelonging(organization *Organization, individual *Individual, role value_object.Role) *OrganizationBelonging {
	return &OrganizationBelonging{Organization: organization, Individual: individual, Role: role}
}
