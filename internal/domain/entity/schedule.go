package entity

import (
	"strconv"
	"strings"
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

func NewScheduleByCron(cron string) *Schedule {
	s := &Schedule{}
	parts := strings.Split(cron, " ")
	for i := 0; i < len(parts); i++ {
		numbers, interval := splitNumbersAndInterval(parts[i])
		switch i {
		case 0:
			s.Minute = numbers
			s.MinuteInterval = interval
		case 1:
			s.Hour = numbers
			s.HourInterval = interval
		case 2:
			s.Day = numbers
			s.DayInterval = interval
		case 3:
			s.Month = numbers
			s.MonthInterval = interval
		case 4:
			s.Weekday = numbers
			s.WeekdayInterval = interval
		}
	}

	return &Schedule{}
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
			ret[i] = 0
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
