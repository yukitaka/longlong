//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import "github.com/yukitaka/longlong/server/core/pkg/domain/entity"

type Habits interface {
	Find(id int) (*entity.Habit, error)
	Create(name, timer string) (*entity.Habit, error)
	Close()
}
