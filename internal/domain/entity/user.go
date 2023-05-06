package entity

type User struct {
	Id int
}

func NewUser(id int) *User {
	return &User{Id: id}
}
