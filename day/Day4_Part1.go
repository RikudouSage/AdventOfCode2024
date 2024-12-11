package day

import (
	"AdventOfCode2024/day4"
	"strconv"
	"strings"
)

type Day4Part1 struct {
}

func (receiver *Day4Part1) DayNumber() float64 {
	return 4.1
}

func (receiver *Day4Part1) Run(inputData string) string {
	inputData = strings.TrimSpace(inputData)

	octagonal := day4.OctagonalContainerFromString(inputData)

	return strconv.Itoa(int(
		octagonal.CountOccurrences([]rune("XMAS")),
	))
}

func init() {
	DaysList = append(DaysList, &Day4Part1{})
}
