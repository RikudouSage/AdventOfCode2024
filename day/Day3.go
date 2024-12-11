package day

import (
	"regexp"
	"strconv"
)

type Day3 struct{}

func (receiver *Day3) DayNumber() uint {
	return 3
}

func (receiver *Day3) Run(inputData string) string {
	regex := regexp.MustCompile("mul\\(([0-9]{1,3}),([0-9]{1,3})\\)")
	matches := regex.FindAllStringSubmatch(inputData, -1)

	var result uint = 0

	for _, match := range matches {
		numA, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		numB, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}

		result += uint(numA * numB)
	}

	return strconv.Itoa(int(result))
}

func init() {
	DaysList = append(DaysList, &Day3{})
}
