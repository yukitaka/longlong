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
	rep := mockRepository.NewMockOrganizations(ctrl)
	rep.EXPECT().Create("TestParent", entity.Avatar{UserId: 1}).Return(expect, nil)

	belongingRep := mockRepository.NewMockOrganizationBelongings(ctrl)
	itr := NewOrganizationCreator(rep, belongingRep)

	id, _ := itr.Create("TestParent", entity.Avatar{UserId: 1})
	if id != expect {
		t.Errorf("NewOrganizationCreator() = %v", id)
	}
}
