package day

import (
	"AdventOfCode2024/day4"
	"strconv"
	"strings"
)

type Day4 struct {
}

func (receiver *Day4) DayNumber() float64 {
	return 4.1
}

func (receiver *Day4) Run(inputData string) string {
	inputData = strings.TrimSpace(inputData)

	octagonal := day4.OctagonalContainerFromString(inputData)

	return strconv.Itoa(int(
		octagonal.CountOccurrences([]rune("XMAS")),
	))
}

func init() {
	DaysList = append(DaysList, &Day4{})
}
