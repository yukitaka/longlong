//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
)

type Individuals interface {
	Create(name string, userId, profileId int) (int, error)
	Find(id int) (*entity.Individual, error)
	FindByUserId(userId int) (*[]entity.Individual, error)
}
