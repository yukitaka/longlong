package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_repository "github.com/yukitaka/longlong/mock/repository"
)

func TestNewOrganizationCreator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rep := mock_repository.NewMockOrganizations(ctrl)
	rep.EXPECT().Create("Test").Return(1)

	finder := NewOrganizationCreator(rep)
	id := finder.Create("Test")
	if id != 1 {
		t.Errorf("NewOrganizationCreator() = %v", id)
	}
}
