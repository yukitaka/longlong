//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
)

type Individuals interface {
	Create(name string, userId, profileId int64) (int64, error)
	Find(id int64) (*entity.Individual, error)
	Close()
}
