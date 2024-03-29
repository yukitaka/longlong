package usecase

import (
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	mockRepository "github.com/yukitaka/longlong/server/core/mock/repository"
)

func TestNewOrganizationCreator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expect := 1
	organizationRep := mockRepository.NewMockOrganizations(ctrl)
	organizationRep.EXPECT().Create("TestParent", entity.Individual{User: entity.NewUser(1)}).Return(expect, nil)

	memberRep := mockRepository.NewMockOrganizationMembers(ctrl)
	rep := NewOrganizationCreatorRepository(organizationRep, memberRep)
	itr := NewOrganizationCreator(rep)

	id, _ := itr.New("TestParent", entity.Individual{User: entity.NewUser(1)})
	if id != expect {
		t.Errorf("NewOrganizationCreator() = %v", id)
	}
}
