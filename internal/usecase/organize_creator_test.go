package usecase

import (
	"github.com/golang/mock/gomock"
	mock_repository "github.com/yukitaka/longlong/mock/repository"
	"testing"
)

func TestNewOrganizeCreator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rep := mock_repository.NewMockOrganizes(ctrl)
	rep.EXPECT().Create("Test").Return(int64(1))

	finder := NewOrganizeCreator(rep)
	id := finder.Create("Test")
	if id != 1 {
		t.Errorf("NewOrganizeCreator() = %v", id)
	}
}
