package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	mockRepository "github.com/yukitaka/longlong/mock/repository"
)

func TestNewOrganizationFinder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expect := int64(1)
	organizationRep := mockRepository.NewMockOrganizations(ctrl)
	organizationRep.EXPECT().Find(expect).Return(&entity.Organization{Id: expect, Name: "Test"}, nil)

	itr := NewOrganizationFinder(organizationRep)
	o, _ := itr.FindById(expect)
	if o.Id != expect || o.Name != "Test" {
		t.Errorf("NewOrganizationFinder() = %v", o)
	}
}
