package entity

type Schedule struct {
	Month           int
	MonthInterval   int
	Day             int
	DayInterval     int
	Hour            int
	HourInterval    int
	Minute          int
	MinuteInterval  int
	Weekday         string
	WeekdayInterval int
}

func NewScheduleByCron(cron string) *Schedule {
	return &Schedule{}
}
