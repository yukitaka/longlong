package entity

type Organize struct {
	ParentID int64
	ID       int64
	Name     string
}

func NewOrganize(parentId, id int64, name string) *Organize {
	return &Organize{ParentID: parentId, ID: id, Name: name}
}
