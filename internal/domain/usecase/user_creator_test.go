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
	userRep := mockRepository.NewMockUsers(ctrl)
	userRep.EXPECT().Create("Name").Return(expect, nil)

	individualRep := mockRepository.NewMockIndividuals(ctrl)
	belongingRep := mockRepository.NewMockOrganizationBelongings(ctrl)

	itr := NewUserCreator(userRep, individualRep, belongingRep)
	id, _ := itr.Users.Create("Name")
	if id != expect {
		t.Errorf("NewUserCreator() = %v", id)
	}
}
