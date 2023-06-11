package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/yukitaka/longlong/internal/domain/entity"
	rep "github.com/yukitaka/longlong/internal/domain/repository"
	"time"
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

func (h *Habits) Find(id int) (*entity.Habit, error) {
	query := "select id, name, exp from habits where id=$1"
	habit := entity.Habit{}
	err := h.DB.Get(&habit, query, id)
	if err != nil {
		return nil, err
	}
	sc, err := h.schedule(habit.Id)
	if err != nil {
		return nil, err
	}
	habit.Schedule = *sc

	return &habit, nil
}

func (h *Habits) schedule(habit_id int) (*entity.Schedule, error) {
	query := "select t.id, duration_type, number, interval, start_at, end_at from schedules t join habits_schedules t1 on t.id=t1.schedule_id where t.id=$1"
	type s struct {
		Id           int       `db:"id"`
		DurationType string    `db:"duration_type"`
		Number       int       `db:"number"`
		Interval     int       `db:"interval"`
		StartAt      time.Time `db:"start_at"`
		EndAt        time.Time `db:"start_at"`
	}
	var ss []s

	err := h.DB.Select(&ss, query, habit_id)
	if err != nil {
		return nil, err
	}
	sc := entity.Schedule{}
	for _, v := range ss {
		switch v.DurationType {
		case "month":
			if sc.MonthInterval != 0 && sc.MonthInterval != v.Interval {
				return nil, fmt.Errorf("invalid interval: %d %d %d", v.Id, sc.MonthInterval, v.Interval)
			}
			sc.Month = append(sc.Month, v.Number)
			sc.MonthInterval = v.Interval
		case "day":
			if sc.DayInterval != 0 && sc.DayInterval != v.Interval {
				return nil, fmt.Errorf("invalid interval: %d %d %d", v.Id, sc.DayInterval, v.Interval)
			}
			sc.Day = append(sc.Day, v.Number)
			sc.DayInterval = v.Interval
		case "hour":
			if sc.HourInterval != 0 && sc.HourInterval != v.Interval {
				return nil, fmt.Errorf("invalid interval: %d %d %d", v.Id, sc.HourInterval, v.Interval)
			}
			sc.Hour = append(sc.Hour, v.Number)
			sc.HourInterval = v.Interval
		case "minute":
			if sc.MinuteInterval != 0 && sc.MinuteInterval != v.Interval {
				return nil, fmt.Errorf("invalid interval: %d %d %d", v.Id, sc.MinuteInterval, v.Interval)
			}
			sc.Minute = append(sc.Minute, v.Number)
			sc.MinuteInterval = v.Interval
		case "weekday":
			if sc.WeekdayInterval != 0 && sc.WeekdayInterval != v.Interval {
				return nil, fmt.Errorf("invalid interval: %d %d %d", v.Id, sc.WeekdayInterval, v.Interval)
			}
			sc.Weekday = append(sc.Weekday, v.Number)
			sc.WeekdayInterval = v.Interval
		}
	}

	return &sc, nil
}
