//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import "github.com/yukitaka/longlong/internal/entity"

type Organizations interface {
	Create(name string) int
	Find(int) (*entity.Organization, error)
}
