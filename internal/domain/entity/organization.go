package entity

type Organization struct {
	ParentId int64
	Id       int64
	Name     string
}

func NewOrganization(parentId, id int64, name string) *Organization {
	return &Organization{ParentId: parentId, Id: id, Name: name}
}
