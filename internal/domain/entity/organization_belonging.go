package entity

type OrganizationBelonging struct {
	OrganizationId int64
	AvatarId       int64
}

func NewOrganizationBelonging(organizationId, avatarId int64) *OrganizationBelonging {
	return &OrganizationBelonging{OrganizationId: organizationId, AvatarId: avatarId}
}
