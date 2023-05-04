package entity

import "github.com/yukitaka/longlong/internal/domain/value_object"

type OrganizationMember struct {
	Organization *Organization
	Individual   *Individual
	Role         value_object.Role
}

func NewOrganizationMember(organization *Organization, individual *Individual, role value_object.Role) *OrganizationMember {
	return &OrganizationMember{Organization: organization, Individual: individual, Role: role}
}
