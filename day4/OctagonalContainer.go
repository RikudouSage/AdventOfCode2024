package day4

import (
	"AdventOfCode2024/helper"
	"strings"
)

type OctagonalContainer [][]rune

func OctagonalContainerFromString(input string) OctagonalContainer {
	return helper.Map(strings.Split(input, "\n"), func(item string) []rune {
		return []rune(item)
	})
}

func (receiver *OctagonalContainer) findHorizontal(input []rune, rowIndex int, columnIndex int) bool {
	for inputIndex := range len(input) {
		if columnIndex+inputIndex >= len((*receiver)[0]) || (*receiver)[rowIndex][columnIndex+inputIndex] != input[inputIndex] {
			return false
		}
	}

	return true
}

func (receiver *OctagonalContainer) findHorizontalReverse(input []rune, rowIndex int, columnIndex int) bool {
	for inputIndex := range len(input) {
		if columnIndex-inputIndex < 0 || (*receiver)[rowIndex][columnIndex-inputIndex] != input[inputIndex] {
			return false
		}
	}

	return true
}

func (receiver *OctagonalContainer) findVertical(input []rune, rowIndex int, columnIndex int) bool {
	for inputIndex := 0; inputIndex < len(input); inputIndex++ {
		if rowIndex+inputIndex >= len(*receiver) || (*receiver)[rowIndex+inputIndex][columnIndex] != input[inputIndex] {
			return false
		}
	}

	return true
}

func (receiver *OctagonalContainer) findVerticalReverse(input []rune, rowIndex int, columnIndex int) bool {
	for inputIndex := range len(input) {
		if rowIndex-inputIndex < 0 || (*receiver)[rowIndex-inputIndex][columnIndex] != input[inputIndex] {
			return false
		}
	}

	return true
}

func (receiver *OctagonalContainer) findDiagonal(input []rune, rowIndex int, columnIndex int) bool {
	for inputIndex := range len(input) {
		if rowIndex+inputIndex >= len(*receiver) || columnIndex+inputIndex >= len((*receiver)[0]) || (*receiver)[rowIndex+inputIndex][columnIndex+inputIndex] != input[inputIndex] {
			return false
		}
	}

	return true
}

func (receiver *OctagonalContainer) findDiagonalReverse(input []rune, rowIndex int, columnIndex int) bool {
	for inputIndex := range len(input) {
		if rowIndex-inputIndex < 0 || columnIndex-inputIndex < 0 || (*receiver)[rowIndex-inputIndex][columnIndex-inputIndex] != input[inputIndex] {
			return false
		}
	}

	return true
}

func (receiver *OctagonalContainer) findOffDiagonal(input []rune, rowIndex int, columnIndex int) bool {
	for inputIndex := range len(input) {
		if rowIndex-inputIndex < 0 || columnIndex+inputIndex >= len((*receiver)[0]) || (*receiver)[rowIndex-inputIndex][columnIndex+inputIndex] != input[inputIndex] {
			return false
		}
	}

	return true
}

func (receiver *OctagonalContainer) findOffDiagonalReverse(input []rune, rowIndex int, columnIndex int) bool {
	for inputIndex := range len(input) {
		if rowIndex+inputIndex >= len(*receiver) || columnIndex-inputIndex < 0 || (*receiver)[rowIndex+inputIndex][columnIndex-inputIndex] != input[inputIndex] {
			return false
		}
	}

	return true
}

func (receiver *OctagonalContainer) CountOccurrences(input []rune) (count uint) {
	for rowIndex := range len(*receiver) {
		for columnIndex := range len((*receiver)[0]) {
			if (*receiver)[rowIndex][columnIndex] != input[0] {
				continue
			}

			if receiver.findHorizontal(input, rowIndex, columnIndex) {
				count++
			}
			if receiver.findHorizontalReverse(input, rowIndex, columnIndex) {
				count++
			}
			if receiver.findVertical(input, rowIndex, columnIndex) {
				count++
			}
			if receiver.findVerticalReverse(input, rowIndex, columnIndex) {
				count++
			}
			if receiver.findDiagonal(input, rowIndex, columnIndex) {
				count++
			}
			if receiver.findDiagonalReverse(input, rowIndex, columnIndex) {
				count++
			}
			if receiver.findOffDiagonal(input, rowIndex, columnIndex) {
				count++
			}
			if receiver.findOffDiagonalReverse(input, rowIndex, columnIndex) {
				count++
			}
		}
	}

	return
}

func (receiver *OctagonalContainer) CountCrossOccurrences(input []rune) (count uint) {
	// todo make this generic, I kinda gave up in the middle

	middleIndex := len(input) / 2
	rowLen := len(*receiver)
	colLen := len((*receiver)[0])

	val := func(row, column int) rune {
		if row < 0 || column < 0 || row >= rowLen || column >= colLen {
			return rune(0)
		}
		return (*receiver)[row][column]
	}

	for rowIndex := range rowLen {
		for columnIndex := range colLen {
			if (*receiver)[rowIndex][columnIndex] != input[middleIndex] {
				continue
			}

			if val(rowIndex-1, columnIndex-1) == input[middleIndex-1] &&
				val(rowIndex-1, columnIndex+1) == input[middleIndex-1] &&
				val(rowIndex+1, columnIndex+1) == input[middleIndex+1] &&
				val(rowIndex+1, columnIndex-1) == input[middleIndex+1] {
				count++
			}

			if val(rowIndex-1, columnIndex-1) == input[middleIndex+1] &&
				val(rowIndex-1, columnIndex+1) == input[middleIndex+1] &&
				val(rowIndex+1, columnIndex+1) == input[middleIndex-1] &&
				val(rowIndex+1, columnIndex-1) == input[middleIndex-1] {
				count++
			}

			if val(rowIndex-1, columnIndex-1) == input[middleIndex-1] &&
				val(rowIndex-1, columnIndex+1) == input[middleIndex+1] &&
				val(rowIndex+1, columnIndex+1) == input[middleIndex+1] &&
				val(rowIndex+1, columnIndex-1) == input[middleIndex-1] {
				count++
			}

			if val(rowIndex-1, columnIndex-1) == input[middleIndex+1] &&
				val(rowIndex-1, columnIndex+1) == input[middleIndex-1] &&
				val(rowIndex+1, columnIndex+1) == input[middleIndex-1] &&
				val(rowIndex+1, columnIndex-1) == input[middleIndex+1] {
				count++
			}
		}
	}

	return
}
