//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
)

type OrganizationBelongings interface {
	Entry(organizationId, avatarId int64) error
	Leave(organizationId, avatarId int64, reason string) error
	Members(organizationId int64) (*[]entity.Avatar, error)
	Close()
}
