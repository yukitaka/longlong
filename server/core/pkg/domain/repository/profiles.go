//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import "github.com/yukitaka/longlong/server/core/pkg/domain/entity"

type Profiles interface {
	Create(nickName, fullName, bio string) (int, error)
	Find(id int) (*entity.Profile, error)
}
