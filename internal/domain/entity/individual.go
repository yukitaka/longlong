package entity

type Individual struct {
	Id   int
	Name string
	User
	Profile
}

func NewIndividual(id int, user User, profile Profile, name string) *Individual {
	return &Individual{Id: id, User: user, Profile: profile, Name: name}
}
