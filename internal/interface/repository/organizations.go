package repository

import (
	"fmt"
	"github.com/yukitaka/longlong/internal/entity"
	rep "github.com/yukitaka/longlong/internal/usecase/repository"
)

type Organizations struct {
	organizations map[int]*entity.Organization
}

func NewOrganizationsRepository() rep.Organizations {
	return &Organizations{
		organizations: make(map[int]*entity.Organization),
	}
}

func (o *Organizations) Create(name string) int {
	o.organizations[0] = &entity.Organization{
		Name: name,
	}
	fmt.Printf("Call to create Organization name by %s.\n", name)

	return 0
}

func (o *Organizations) Find(id int) (*entity.Organization, error) {
	o.organizations[id] = &entity.Organization{
		Name: "example",
	}

	return o.organizations[id], nil
}
