//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
)

type Organizations interface {
	Create(name string, individual entity.Individual) (int, error)
	Find(id int) (*entity.Organization, error)
	FindByName(name string) (*entity.Organization, error)
	FindAll(ids []interface{}) (*[]entity.Organization, error)
	List() (*[]entity.Organization, error)
}
