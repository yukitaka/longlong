package entity

type Individual struct {
	Id        int64
	Name      string
	UserId    int64
	ProfileId int64
}

func NewIndividual(id, userId, profileId int64, name string) *Individual {
	return &Individual{Id: id, UserId: userId, ProfileId: profileId, Name: name}
}
