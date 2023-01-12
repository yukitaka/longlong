package usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/yukitaka/longlong/internal/entity"
	mock_repository "github.com/yukitaka/longlong/mock/repository"
	"testing"
)

func TestNewOrganizeFinder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rep := mock_repository.NewMockOrganizes(ctrl)
	rep.EXPECT().Find(1).Return(&entity.Organize{ID: 1, Name: "Test"}, nil)

	finder := NewOrganizeFinder(rep)
	o, _ := finder.FindById(1)
	if o.ID != 1 || o.Name != "Test" {
		t.Errorf("NewOrganizeFinder() = %v", o)
	}
}
