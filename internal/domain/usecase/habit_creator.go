package usecase

import "github.com/yukitaka/longlong/internal/domain/repository"

type HabitCreator struct {
	repository.Habits
}

func NewHabitCreator(habits repository.Habits) *HabitCreator {
	return &HabitCreator{habits}
}
