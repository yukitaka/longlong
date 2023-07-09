package usecase

import (
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"github.com/yukitaka/longlong/server/core/pkg/domain/repository"
)

type HabitCreator struct {
	repository.Habits
}

func NewHabitCreator(habits repository.Habits) *HabitCreator {
	return &HabitCreator{habits}
}

func (it *HabitCreator) New(name, timer string) (*entity.Habit, error) {
	habit, err := it.Habits.Create(name, timer)
	if err != nil {
		return nil, err
	}

	return habit, nil
}
