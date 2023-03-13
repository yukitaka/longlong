package repository

import (
	"errors"
	"fmt"
	"github.com/yukitaka/longlong/internal/domain/entity"
	rep "github.com/yukitaka/longlong/internal/domain/usecase/repository"
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
	id := 0
	for key := range o.organizations {
		fmt.Println(key)
		if key > id {
			id = key
		}
	}
	id++

	o.organizations[id] = &entity.Organization{
		Name: name,
	}
	fmt.Printf("Call to create Organization name by %s %d.\n", name, id)

	return id
}

func (o *Organizations) Find(id int) (*entity.Organization, error) {
	if organization, ok := o.organizations[id]; ok {
		return organization, nil
	}
	fmt.Printf("Call to find Organization id by %d.\n", id)

	return nil, errors.New(fmt.Sprintf("%d is not found.", id))
}
