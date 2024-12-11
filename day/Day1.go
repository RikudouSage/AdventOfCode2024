package day

import (
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Day1 struct {
}

func (receiver *Day1) DayNumber() uint {
	return 1
}

func (receiver *Day1) Run(inputData string) string {
	spaceSplitRegex := regexp.MustCompile(" +")
	lineAmount := len(strings.Split(inputData, "\n"))

	split := make([][]int, 2)
	split[0] = make([]int, 0, lineAmount)
	split[1] = make([]int, 0, lineAmount)

	for _, line := range strings.Split(inputData, "\n") {
		if line == "" {
			continue
		}
		parts := spaceSplitRegex.Split(line, 2)
		parsedA, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		parsedB, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		split[0] = append(split[0], parsedA)
		split[1] = append(split[1], parsedB)
	}

	sortFunc := func(a, b int) int {
		if a == b {
			return 0
		}

		if a < b {
			return -1
		}

		return 1
	}

	slices.SortFunc(split[0], sortFunc)
	slices.SortFunc(split[1], sortFunc)

	var result uint = 0
	for i := range split[0] {
		result += uint(math.Abs(float64(split[0][i] - split[1][i])))
	}

	return strconv.Itoa(int(result))
}

func init() {
	DaysList = append(DaysList, &Day1{})
}
