package entity

type Authentication struct {
	Id       int
	Identify string
	Token    string
}

func NewAuthentication(id int, identify string, token string) *Authentication {
	return &Authentication{Id: id, Identify: identify, Token: token}
}
