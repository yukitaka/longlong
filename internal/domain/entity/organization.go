package entity

type Organization struct {
	ParentId int
	Id       int
	Name     string
}

func NewOrganization(parentId, id int, name string) *Organization {
	return &Organization{ParentId: parentId, Id: id, Name: name}
}
