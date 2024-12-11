package day

type Day interface {
	DayNumber() float64
	Run(inputData string) string
}
type Days []Day

var DaysList = make(Days, 0, 24)

func (receiver Days) FindByNumber(dayNumber uint) []Day {
	result := make([]Day, 0)
	for _, day := range receiver {
		if day.DayNumber() >= float64(dayNumber) && day.DayNumber() < float64(dayNumber+1) {
			result = append(result, day)
		}
	}

	return result
}
