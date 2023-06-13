package repository

import (
	"github.com/jmoiron/sqlx"
	rep "github.com/yukitaka/longlong/internal/domain/repository"
)

type Timers struct {
	*sqlx.DB
}

func NewTimersRepository(con *sqlx.DB) rep.Timers {
	return &Timers{
		DB: con,
	}
}
