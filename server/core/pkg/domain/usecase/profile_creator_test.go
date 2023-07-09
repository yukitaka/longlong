package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	mockRepository "github.com/yukitaka/longlong/server/core/mock/repository"
)

func TestNewProfileCreator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	profileRep := mockRepository.NewMockProfiles(ctrl)
	profileRep.EXPECT().Create("Test", "Test", "Bio").Return(1, nil)

	itr := NewProfileCreator(profileRep)
	id, _ := itr.New(nil, "Test", "Test", "Bio")
	if id != 1 {
		t.Errorf("NewProfileCreator() = %v", id)
	}
}
