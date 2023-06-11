package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/yukitaka/longlong/internal/domain/entity"
	rep "github.com/yukitaka/longlong/internal/domain/repository"
)

type Habits struct {
	*sqlx.DB
}

func NewHabitsRepository(con *sqlx.DB) rep.Habits {
	return &Habits{
		DB: con,
	}
}

func (h *Habits) Close() {
	err := h.DB.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func (h Habits) Find(id int) (*entity.Habit, error) {
	query := "select * from habits where id = $1"

	panic("implement me " + query)
}
