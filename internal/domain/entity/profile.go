package entity

type Profile struct {
	Id       int64
	Name     string
	FullName string
}

func NewProfile(id int64, name string) *Profile {
	return &Profile{Id: id, Name: name}
}
