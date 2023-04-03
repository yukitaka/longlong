package entity

type Authentication struct {
	Id       int64
	Identify string
	Token    string
}

func NewAuthentication(id int64, identify string, token string) *Authentication {
	return &Authentication{Id: id, Identify: identify, Token: token}
}
