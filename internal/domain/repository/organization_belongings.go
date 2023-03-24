//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
)

type OrganizationBelongings interface {
	Entry(organizationId, userId int64) error
	Leave(organizationId, userId int64, reason string) error
	Members(organizationId int64) (*[]entity.User, error)
	Close()
}
