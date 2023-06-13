//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import "github.com/yukitaka/longlong/internal/domain/entity"

type Habits interface {
	Find(id int) (*entity.Habit, error)
	Close()
}