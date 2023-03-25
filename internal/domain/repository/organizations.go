//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
)

type Organizations interface {
	Create(name string) (int64, error)
	Find(id int64) (*entity.Organization, error)
	List() (*[]entity.Organization, error)
	Close()
}