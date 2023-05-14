package entity

type Profile struct {
	Id        int
	NickName  string
	FullName  string
	Biography string
}

func NewProfile(id int, nickName string, fullName string, bio string) *Profile {
	return &Profile{Id: id, NickName: nickName, FullName: fullName, Biography: bio}
}
