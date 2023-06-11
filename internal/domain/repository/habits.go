package repository

import "github.com/yukitaka/longlong/internal/domain/entity"

type Habits interface {
	Find(id int) (*entity.Habit, error)
	Close()
}
