package entity

import (
	"reflect"
	"testing"
)

func TestScheduleByCron(t *testing.T) {
	sc, _ := NewScheduleByCron("1 * * * *")
	if sc.MinuteInterval != 1 {
		t.Errorf("MinuteInterval is not 1: %d", sc.HourInterval)
	}
	if sc.Minute == nil {
		t.Errorf("Minute is not 1: %d", sc.HourInterval)
	}
	sc, _ = NewScheduleByCron("1/2 * * * *")
	if sc.MinuteInterval != 2 {
		t.Errorf("MinuteInterval is not 2: %d", sc.MinuteInterval)
	}
	if sc.Minute == nil {
		t.Errorf("Minute is not 1: %d", sc.Minute)
	}
}

func TestSplitNumbersAndInterval(t *testing.T) {
	type expect struct {
		Numbers  []int
		Interval int
	}

	testCase := map[string]expect{
		"*":       {[]int{0}, 1},
		"*/10":    {[]int{0}, 10},
		"0":       {[]int{0}, 1},
		"0/3":     {[]int{0}, 3},
		"31":      {[]int{31}, 1},
		"31/2":    {[]int{31}, 2},
		"1,2,3":   {[]int{1, 2, 3}, 1},
		"1,2,3/2": {[]int{1, 2, 3}, 2},
	}
	for k, v := range testCase {
		n, i := splitNumbersAndInterval(k)
		if !reflect.DeepEqual(v, expect{n, i}) {
			t.Errorf("Error! Expect: %v, Numbers: %v, Interval: %v", v, n, i)
		}
	}
}
