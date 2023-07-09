package usecase

import (
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"github.com/yukitaka/longlong/server/core/pkg/domain/value_object"
	"testing"

	"github.com/golang/mock/gomock"
	mockRepository "github.com/yukitaka/longlong/server/core/mock/repository"
)

func TestNewUserCreator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expect := 1
	userRep := mockRepository.NewMockUsers(ctrl)
	userRep.EXPECT().Create("Name").Return(expect, nil)

	individualRep := mockRepository.NewMockIndividuals(ctrl)
	individualRep.EXPECT().Create("Name", 1, -1).Return(expect, nil)
	memberRep := mockRepository.NewMockOrganizationMembers(ctrl)
	memberRep.EXPECT().Entry(1, 1, value_object.MEMBER).Return(nil)

	rep := NewUserCreatorRepository(userRep, individualRep, memberRep)
	itr := NewUserCreator(rep)
	operator := entity.NewOrganizationMember(
		entity.NewOrganization(0, 1, "Test"),
		entity.NewIndividual(1, entity.NewUser(1), entity.NewProfile(1, "", "", ""), "Test"),
		value_object.ADMIN,
	)
	id, _ := itr.New(operator, "Name", "MEMBER")
	if id != expect {
		t.Errorf("NewUserCreator() = %v", id)
	}
}
