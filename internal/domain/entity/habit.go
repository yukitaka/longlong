package entity

type Habit struct {
	Id       int
	Schedule Schedule
	Name     string
	Exp      int
}

func NewHabit(id int, schedule Schedule, name string, exp int) *Habit {
	return &Habit{
		Id:       id,
		Schedule: schedule,
		Name:     name,
		Exp:      exp,
	}
}
