package day

import (
	"AdventOfCode2024/day4"
	"strconv"
	"strings"
)

type Day4Part2 struct {
}

func (receiver *Day4Part2) DayNumber() float64 {
	return 4.2
}

func (receiver *Day4Part2) Run(inputData string) string {
	inputData = strings.TrimSpace(inputData)
	octagonal := day4.OctagonalContainerFromString(inputData)

	return strconv.Itoa(int(
		octagonal.CountCrossOccurrences([]rune("MAS")),
	))
}

func init() {
	DaysList = append(DaysList, &Day4Part2{})
}
