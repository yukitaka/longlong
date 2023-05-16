//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

type Profiles interface {
	Create(nickName, fullName, bio string) (int, error)
	Close()
}
