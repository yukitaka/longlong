package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	mock_repository "github.com/yukitaka/longlong/mock/repository"
)

func TestNewOrganizationFinder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expect := int64(1)
	rep := mock_repository.NewMockOrganizations(ctrl)
	rep.EXPECT().Find(expect).Return(&entity.Organization{ID: expect, Name: "Test"}, nil)

	finder := NewOrganizationFinder(rep)
	o, _ := finder.FindById(expect)
	if o.ID != expect || o.Name != "Test" {
		t.Errorf("NewOrganizationFinder() = %v", o)
	}
}
