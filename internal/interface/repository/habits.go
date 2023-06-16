package repository

import (
	"database/sql"
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
	t, err := h.timer(habit.Id)
	if err != nil {
		return nil, err
	}
	habit.Timer = *t

	return &habit, nil
}

func (h *Habits) Create(name, timer string) (*entity.Habit, error) {
	query := "select max(id) from habits"
	row := h.DB.QueryRowx(query)
	var nullableId sql.NullInt32
	err := row.Scan(&nullableId)
	if err != nil {
		return nil, err
	}
	id := 0
	if nullableId.Valid {
		id = int(nullableId.Int32)
		id++
	}

	t, err := entity.NewTimerByCronSyntax(timer)
	if err != nil {
		return nil, err
	}
	tx, err := h.DB.Begin()
	if err != nil {
		return nil, err
	}
	query = "insert into habits (id, name, exp) values ($1, $2, $3)"
	_, err = h.DB.Exec(query, id, name, 0)
	if err != nil {
		fmt.Println(id)
		return nil, err
	}
	err = tx.Commit()

	return &entity.Habit{Id: id, Name: name, Timer: *t}, nil
}

func (h *Habits) timer(habit_id int) (*entity.Timer, error) {
	query := "select t.id, duration_type, number, interval, reference_at from timers t join habits_timers t1 on t.id=t1.timer_id where t.id=$1"
	type s struct {
		Id           int       `db:"id"`
		DurationType string    `db:"duration_type"`
		Number       int       `db:"number"`
		Interval     int       `db:"interval"`
		ReferenceAt  time.Time `db:"reference_at"`
	}
	var ss []s

	err := h.DB.Select(&ss, query, habit_id)
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
