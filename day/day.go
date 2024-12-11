package day

import "strconv"

type Day interface {
	DayNumber() uint
	Run(inputData string) string
}
type Days []Day

var DaysList = make(Days, 0, 24)

func (receiver Days) FindByNumber(dayNumber uint) Day {
	for _, day := range receiver {
		if day.DayNumber() == dayNumber {
			return day
		}
	}

	panic("day number not found: " + strconv.Itoa(int(dayNumber)))
}
