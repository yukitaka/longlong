package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/repository"
)

type ProfileCreator struct {
	repository.Profiles
}

func NewProfileCreator(profiles repository.Profiles) *ProfileCreator {
	return &ProfileCreator{profiles}
}

func (it *ProfileCreator) New(operator *entity.OrganizationMember, nickName, fullName, bio string) (int, error) {
	profileId, err := it.Profiles.Create(nickName, fullName, bio)
	if err != nil {
		return -1, err
	}

	return profileId, nil
}
