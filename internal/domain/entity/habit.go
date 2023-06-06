package entity

type Habit struct {
	Id       int
	Schedule string
	Name     string
	Exp      int
}

func NewHabit(id int, schedule, name string, exp int) *Habit {
	return &Habit{
		Id:       id,
		Schedule: schedule,
		Name:     name,
		Exp:      exp,
	}
}
