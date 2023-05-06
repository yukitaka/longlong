package entity

type Individual struct {
	Id        int
	Name      string
	UserId    int
	ProfileId int
}

func NewIndividual(id, userId, profileId int, name string) *Individual {
	return &Individual{Id: id, UserId: userId, ProfileId: profileId, Name: name}
}
