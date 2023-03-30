//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
)

type OrganizationBelongings interface {
	Entry(avatarId int64) error
	Leave(avatarId int64, reason string) error
	Members() (*[]entity.Avatar, error)
	Close()
}
