package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_repository "github.com/yukitaka/longlong/mock/repository"
)

func TestNewUserCreator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expect := int64(1)
	rep := mock_repository.NewMockUsers(ctrl)
	rep.EXPECT().Create("Name").Return(expect, nil)

	itr := NewUserCreator(rep)
	id, _ := itr.Create("Name")
	if id != expect {
		t.Errorf("NewUserCreator() = %v", id)
	}
}
