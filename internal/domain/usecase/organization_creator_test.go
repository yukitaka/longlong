package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	mockRepository "github.com/yukitaka/longlong/mock/repository"
)

func TestNewOrganizationCreator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expect := int64(1)
	organizationRep := mockRepository.NewMockOrganizations(ctrl)
	organizationRep.EXPECT().Create("TestParent", entity.Individual{UserId: 1}).Return(expect, nil)

	belongingRep := mockRepository.NewMockOrganizationBelongings(ctrl)
	itr := NewOrganizationCreator(organizationRep, belongingRep)

	id, _ := itr.Create("TestParent", entity.Individual{UserId: 1})
	if id != expect {
		t.Errorf("NewOrganizationCreator() = %v", id)
	}
}
