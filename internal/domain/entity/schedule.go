package entity

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Schedule struct {
	Month           []int
	MonthInterval   int
	Day             []int
	DayInterval     int
	Hour            []int
	HourInterval    int
	Minute          []int
	MinuteInterval  int
	Weekday         []int
	WeekdayInterval int
}

func (s *Schedule) IsExecute(time time.Time) bool {
	if s.MinuteInterval > 0 {
		m := time.Minute()
		for _, v := range s.Minute {
			if v == -1 || v == m {
				return true
			}
		}
	}

	return false
}

func (s *Schedule) SetMinute(minutes []int, interval int) error {
	for _, m := range minutes {
		if m < 0 || m > 59 {
			return fmt.Errorf("Error: minute %d is out of range", m)
		}
	}
	s.Minute = minutes
	s.MinuteInterval = interval

	return nil
}

func (s *Schedule) SetHour(hours []int, interval int) error {
	for _, m := range hours {
		if m < 0 || m > 23 {
			return fmt.Errorf("Error: hour %d is out of range", m)
		}
	}
	s.Hour = hours
	s.HourInterval = interval

	return nil
}

func (s *Schedule) SetDay(days []int, interval int) error {
	for _, m := range days {
		if m < 0 || m > 31 {
			return fmt.Errorf("Error: day %d is out of range", m)
		}
	}
	s.Day = days
	s.DayInterval = interval

	return nil
}

func (s *Schedule) SetMonth(months []int, interval int) error {
	for _, m := range months {
		if m < 0 || m > 12 {
			return fmt.Errorf("Error: month %d is out of range", m)
		}
	}
	s.Month = months
	s.MonthInterval = interval

	return nil
}

func (s *Schedule) SetWeekday(weekdays []int, interval int) error {
	for _, m := range weekdays {
		if m < 0 || m > 7 {
			return fmt.Errorf("Error: weekday %d is out of range", m)
		}
	}
	s.Weekday = weekdays
	s.WeekdayInterval = interval

	return nil
}

func NewScheduleByCron(cron string) (*Schedule, error) {
	s := &Schedule{}
	parts := strings.Split(cron, " ")
	for i := 0; i < len(parts); i++ {
		numbers, interval := splitNumbersAndInterval(parts[i])
		err := error(nil)
		switch i {
		case 0:
			err = s.SetMinute(numbers, interval)
		case 1:
			err = s.SetHour(numbers, interval)
		case 2:
			err = s.SetDay(numbers, interval)
		case 3:
			s.Month = numbers
			s.MonthInterval = interval
		case 4:
			s.Weekday = numbers
			s.WeekdayInterval = interval
		}
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func splitNumbersAndInterval(s string) ([]int, int) {
	interval := 1
	parts := strings.Split(s, "/")
	if len(parts) == 2 {
		num, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		interval = num
	}
	numbers := strings.Split(parts[0], ",")
	ret := make([]int, len(numbers))
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == "*" {
			ret[i] = -1
		} else {
			num, err := strconv.Atoi(numbers[i])
			if err != nil {
				panic(err)
			}
			ret[i] = num
		}
	}

	return ret, interval
}
