package entity

type OrganizationBelonging struct {
	OrganizationId int64
	UserId         int64
}

func NewOrganizationBelonging(organizationId, userId int64) *OrganizationBelonging {
	return &OrganizationBelonging{OrganizationId: organizationId, UserId: userId}
}
