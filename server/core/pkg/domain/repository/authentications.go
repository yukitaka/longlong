//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import "time"

type Authentications interface {
	Create(organizationId int, identify, token string) (int, error)
	FindToken(organizationId int, identify string) (int, string, error)
	UpdateToken(id int, token string) error
	StoreOAuth2Info(identify, accessToken, refreshToken string, expiry time.Time) (bool, error)
	Close()
}
