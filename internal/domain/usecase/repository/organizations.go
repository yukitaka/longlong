//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
)

type Organizations interface {
	Create(name string) int64
	Find(id int64) (*entity.Organization, error)
	Close()
}
