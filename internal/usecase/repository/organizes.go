//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import "github.com/yukitaka/longlong/internal/entity"

type Organizes interface {
	Create(name string) int64
	Find(int64) (*entity.Organize, error)
}
