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

	rep := mock_repository.NewMockOrganizations(ctrl)
	rep.EXPECT().Find(1).Return(&entity.Organization{ID: 1, Name: "Test"}, nil)

	finder := NewOrganizationFinder(rep)
	o, _ := finder.FindById(1)
	if o.ID != 1 || o.Name != "Test" {
		t.Errorf("NewOrganizationFinder() = %v", o)
	}
}
