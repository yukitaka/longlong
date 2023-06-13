package entity

type Habit struct {
	Id    int
	Timer Timer
	Name  string
	Exp   int
}

func NewHabit(id int, timer Timer, name string, exp int) *Habit {
	return &Habit{
		Id:    id,
		Timer: timer,
		Name:  name,
		Exp:   exp,
	}
}
