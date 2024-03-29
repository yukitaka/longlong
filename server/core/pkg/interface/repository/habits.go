package repository

import (
	"database/sql"
	"fmt"
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	rep "github.com/yukitaka/longlong/server/core/pkg/domain/repository"
	"github.com/yukitaka/longlong/server/core/pkg/interface/datastore"
	"time"
)

type Habits struct {
	*datastore.Connection
}

func NewHabitsRepository(con *datastore.Connection) rep.Habits {
	return &Habits{
		Connection: con,
	}
}

func (rep *Habits) Find(id int) (*entity.Habit, error) {
	query := "select id, name, exp from habits where id=$1"
	habit := entity.Habit{}
	err := rep.DB.Get(&habit, query, id)
	if err != nil {
		return nil, err
	}
	t, err := rep.timer(habit.Id)
	if err != nil {
		return nil, err
	}
	habit.Timer = *t

	return &habit, nil
}

func (rep *Habits) Create(name, timer string) (*entity.Habit, error) {
	id, err := rep.nextId("habits")
	if err != nil {
		return nil, err
	}

	t, err := entity.NewTimerByCronSyntax(timer)
	if err != nil {
		return nil, err
	}
	tx, err := rep.DB.Begin()
	if err != nil {
		return nil, err
	}
	query := "insert into habits (id, name, exp) values ($1, $2, $3)"
	_, err = rep.DB.Exec(query, id, name, 0)
	if err != nil {
		return nil, err
	}
	timerIds, err := (&Timers{rep.Connection}).InsertTimers(t)
	for _, v := range timerIds {
		query = "insert into habits_timers (habit_id, timer_id) values ($1, $2)"
		_, err = rep.DB.Exec(query, id, v)
		if err != nil {
			return nil, err
		}
	}
	err = tx.Commit()

	return &entity.Habit{Id: id, Name: name, Timer: *t}, nil
}

func (rep *Habits) timer(habitId int) (*entity.Timer, error) {
	query := "select t.id, duration_type, number, interval, reference_at from timers t join habits_timers t1 on t.id=t1.timer_id where t.id=$1"
	type s struct {
		Id           int       `db:"id"`
		DurationType string    `db:"duration_type"`
		Number       int       `db:"number"`
		Interval     int       `db:"interval"`
		ReferenceAt  time.Time `db:"reference_at"`
	}
	var ss []s

	err := rep.DB.Select(&ss, query, habitId)
	if err != nil {
		return nil, err
	}
	t := entity.Timer{}
	for _, v := range ss {
		switch v.DurationType {
		case "month":
			if t.MonthInterval != 0 && t.MonthInterval != v.Interval {
				return nil, fmt.Errorf("invalid interval: %d %d %d", v.Id, t.MonthInterval, v.Interval)
			}
			t.Month = append(t.Month, v.Number)
			t.MonthInterval = v.Interval
		case "day":
			if t.DayInterval != 0 && t.DayInterval != v.Interval {
				return nil, fmt.Errorf("invalid interval: %d %d %d", v.Id, t.DayInterval, v.Interval)
			}
			t.Day = append(t.Day, v.Number)
			t.DayInterval = v.Interval
		case "hour":
			if t.HourInterval != 0 && t.HourInterval != v.Interval {
				return nil, fmt.Errorf("invalid interval: %d %d %d", v.Id, t.HourInterval, v.Interval)
			}
			t.Hour = append(t.Hour, v.Number)
			t.HourInterval = v.Interval
		case "minute":
			if t.MinuteInterval != 0 && t.MinuteInterval != v.Interval {
				return nil, fmt.Errorf("invalid interval: %d %d %d", v.Id, t.MinuteInterval, v.Interval)
			}
			t.Minute = append(t.Minute, v.Number)
			t.MinuteInterval = v.Interval
		case "weekday":
			if t.WeekdayInterval != 0 && t.WeekdayInterval != v.Interval {
				return nil, fmt.Errorf("invalid interval: %d %d %d", v.Id, t.WeekdayInterval, v.Interval)
			}
			t.Weekday = append(t.Weekday, v.Number)
			t.WeekdayInterval = v.Interval
		}
	}

	return &t, nil
}

func (rep *Habits) nextId(table string) (int, error) {
	query := fmt.Sprintf("select max(id) from %s", table)
	row := rep.DB.QueryRowx(query)
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
