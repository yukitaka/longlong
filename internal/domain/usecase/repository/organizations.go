//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
)

type Organizations interface {
	Create(name string) int
	Find(id int) (*entity.Organization, error)
}
