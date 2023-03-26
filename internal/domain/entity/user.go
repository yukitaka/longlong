package entity

type User struct {
	Id int64
}

func NewUser(id int64) *User {
	return &User{Id: id}
}
