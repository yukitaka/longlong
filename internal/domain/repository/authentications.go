//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

type Authentications interface {
	Create(identify, token string) (int64, error)
	FindToken(identify string) (int64, string, error)
	Close()
}
