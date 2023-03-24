package entity

type User struct {
	OrganizationId int64
	Id             int64
	Name           string
	FullName       string
}

func NewUser(organizationId, id int64, name string) *User {
	return &User{OrganizationId: organizationId, Id: id, Name: name}
}
