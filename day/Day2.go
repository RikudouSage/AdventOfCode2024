package day

import (
	. "AdventOfCode2024/helper"
	"math"
	"strconv"
	"strings"
)

type lineDirection int

const (
	up lineDirection = iota
	down
)

type Day2 struct {
}

func (receiver *Day2) DayNumber() uint {
	return 2
}

func getDirection(first int, second int) *lineDirection {
	if first > second {
		return Pointer(down)
	} else {
		return Pointer(up)
	}
}

func (receiver *Day2) Run(inputData string) string {
	inputLines := strings.Split(inputData, "\n")

	var result uint = 0

outer:
	for _, line := range inputLines {
		if line == "" {
			continue
		}
		split := Map(strings.Split(line, " "), func(item string) int {
			converted, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}

			return converted
		})

		var previous *int
		var direction *lineDirection

		for _, num := range split {
			skip := (func() (skip bool) {
				defer func() {
					previous = &num
				}()

				if previous == nil {
					return
				}

				diff := math.Abs(float64(num - *previous))
				if diff > 3 || diff < 1 {
					skip = true
					return
				}

				if direction == nil {
					direction = getDirection(*previous, num)
					return
				}

				currentDirection := getDirection(*previous, num)
				if *currentDirection != *direction {
					skip = true
					return
				}

				return
			})()
			if skip {
				continue outer
			}
		}

		result += 1
	}

	return strconv.Itoa(int(result))
}

func init() {
	DaysList = append(DaysList, &Day2{})
}
