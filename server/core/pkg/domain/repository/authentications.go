//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"time"
)

type Authentications interface {
	Create(organizationId int, identify, token string) (int, error)
	Find(organizationId int, identify string) (*entity.Authentication, error)
	FindToken(organizationId int, identify string) (int, string, error)
	UpdateToken(id int, token string) error
	Store(organizationId int, identify, token string) (bool, error)
	StoreOAuth2Info(identify, accessToken, refreshToken string, expiry time.Time) (bool, error)
}
