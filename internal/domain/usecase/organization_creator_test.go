package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	mockRepository "github.com/yukitaka/longlong/mock/repository"
)

func TestNewOrganizationCreator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expect := int64(1)
	rep := mockRepository.NewMockOrganizations(ctrl)
	rep.EXPECT().Create("TestParent").Return(expect, nil)

	itr := NewOrganizationCreator(rep)
	id, _ := itr.Create("TestParent")
	if id != expect {
		t.Errorf("NewOrganizationCreator() = %v", id)
	}
}
