//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
)

type Users interface {
	Create(name string) (int, error)
	Find(id int) (*entity.User, error)
	Close()
}
