package entity

import "github.com/yukitaka/longlong/internal/domain/value_object"

type OrganizationBelonging struct {
	OrganizationId int64
	IndividualId   int64
	Role           value_object.Role
}

func NewOrganizationBelonging(organizationId, individualId int64, role value_object.Role) *OrganizationBelonging {
	return &OrganizationBelonging{OrganizationId: organizationId, IndividualId: individualId, Role: role}
}
