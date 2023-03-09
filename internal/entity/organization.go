package entity

type Organization struct {
	ParentID int64
	ID       int64
	Name     string
}

func NewOrganization(parentId, id int64, name string) *Organization {
	return &Organization{ParentID: parentId, ID: id, Name: name}
}
