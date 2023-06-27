//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import "time"

type Authentications interface {
	Create(identify, token string) (int, error)
	FindToken(identify string) (int, string, error)
	UpdateToken(id int, token string) error
	StoreOAuth2Info(identify, accessToken, tokenType, refreshToken string, expiry time.Time) (bool, error)
	Close()
}
