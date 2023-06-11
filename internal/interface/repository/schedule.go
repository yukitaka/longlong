package repository

import (
	"github.com/jmoiron/sqlx"
	rep "github.com/yukitaka/longlong/internal/domain/repository"
)

type Schedules struct {
	*sqlx.DB
}

func NewSchedulesRepository(con *sqlx.DB) rep.Schedule {
	return &Schedules{
		DB: con,
	}
}
