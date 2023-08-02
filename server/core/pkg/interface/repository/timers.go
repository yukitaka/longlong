package repository

import (
	"database/sql"
	"fmt"
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	rep "github.com/yukitaka/longlong/server/core/pkg/domain/repository"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
	"time"
)

type Timers struct {
	*datastore.Connection
}

func NewTimersRepository(con *datastore.Connection) rep.Timers {
	return &Timers{
		Connection: con,
	}
}

func (t *Timers) InsertTimers(timer *entity.Timer) ([]int, error) {
	var ids []int
	insertTimer := func(durationType string, numbers []int, interval int) error {
		if interval <= 0 || len(numbers) == 0 {
			return nil
		}
		id, err := t.nextId("timers")
		if err != nil {
			return err
		}
		query := "insert into timers (id, duration_type, number, interval, reference_at) values ($1, $2, $3, $4, $5)"
		for _, v := range numbers {
			_, err = t.DB.Exec(query, id, durationType, v, interval, time.Now())
			if err != nil {
				return err
			}
			ids = append(ids, id)
		}
		return nil
	}
	if err := insertTimer("minute", timer.Minute, timer.MinuteInterval); err != nil {
		return nil, err
	}
	if err := insertTimer("hour", timer.Hour, timer.HourInterval); err != nil {
		return nil, err
	}
	if err := insertTimer("day", timer.Day, timer.DayInterval); err != nil {
		return nil, err
	}
	if err := insertTimer("month", timer.Month, timer.MonthInterval); err != nil {
		return nil, err
	}
	if err := insertTimer("weekday", timer.Weekday, timer.WeekdayInterval); err != nil {
		return nil, err
	}

	return ids, nil
}

func (t *Timers) nextId(table string) (int, error) {
	query := fmt.Sprintf("select max(id) from %s", table)
	row := t.DB.QueryRowx(query)
	var nullableId sql.NullInt32
	err := row.Scan(&nullableId)
	if err != nil {
		return -1, err
	}
	id := 0
	if nullableId.Valid {
		id = int(nullableId.Int32)
		id++
	}

	return id, nil
}
