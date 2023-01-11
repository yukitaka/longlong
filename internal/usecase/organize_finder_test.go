package usecase

import (
	"github.com/yukitaka/longlong/internal/entity"
	"testing"
)

type mock struct{}

func (rep *mock) Find(id int64) (*entity.Organize, error) {
	return entity.NewOrganize(0, id, "Test"), nil
}

func TestNewOrganizeFinder(t *testing.T) {
	rep := mock{}
	finder := NewOrganizeFinder(&rep)
	o, _ := finder.FindById(1)
	if o.ID != 1 || o.Name != "Test" {
		t.Errorf("NewOrganizeFinder() = %v", o)
	}
}
