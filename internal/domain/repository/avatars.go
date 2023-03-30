//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
)

type Avatars interface {
	Create(name string) (int64, error)
	Find(id int64) (*entity.Avatar, error)
	Close()
}
