package entity

type Avatar struct {
	Id        int64
	Name      string
	UserId    int64
	ProfileId int64
}

func NewAvatar(id, userId, profileId int64, name string) *Avatar {
	return &Avatar{Id: id, UserId: userId, ProfileId: profileId, Name: name}
}
