package entity

type Authentication struct {
	Id             int
	OrganizationId int
	Identify       string
	Token          string
}

func NewAuthentication(id int, organizationId int, identify string, token string) *Authentication {
	return &Authentication{Id: id, OrganizationId: organizationId, Identify: identify, Token: token}
}
