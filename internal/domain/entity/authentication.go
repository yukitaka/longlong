package entity

type Authentication struct {
	Id    int64
	Name  string
	Token string
}

func NewAuthentication(id int64, name string, token string) *Authentication {
	return &Authentication{Id: id, Name: name, Token: token}
}
