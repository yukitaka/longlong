package entity

type Profile struct {
	Id       int
	Name     string
	FullName string
}

func NewProfile(id int, name string) *Profile {
	return &Profile{Id: id, Name: name}
}
