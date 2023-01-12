//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import "github.com/yukitaka/longlong/internal/entity"

type Organizes interface {
	Create(name string) int
	Find(int) (*entity.Organize, error)
}
