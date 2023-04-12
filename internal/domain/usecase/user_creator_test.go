package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	mockRepository "github.com/yukitaka/longlong/mock/repository"
)

func TestNewUserCreator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expect := int64(1)
	rep := mockRepository.NewMockUsers(ctrl)
	rep.EXPECT().Create("Name").Return(expect, nil)

	individualRep := mockRepository.NewMockIndividuals(ctrl)

	itr := NewUserCreator(rep, individualRep)
	id, _ := itr.Users.Create("Name")
	if id != expect {
		t.Errorf("NewUserCreator() = %v", id)
	}
}
