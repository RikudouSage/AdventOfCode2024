package main

import (
	dayPkg "AdventOfCode2024/day"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("You must provide a day number and input file")
		os.Exit(1)
	}

	dayNumber, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Day must be a number")
		os.Exit(1)
	}

	file, err := os.Open(args[2])
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	days := dayPkg.DaysList.FindByNumber(uint(dayNumber))
	if len(days) == 0 {
		fmt.Println("Error: Day not found")
		os.Exit(1)
	}

	result := make([]string, len(days))

	for i, day := range days {
		result[i] = day.Run(string(content))
	}

	if len(result) == 1 {
		fmt.Println(result[0])
	} else {
		for i, dayResult := range result {
			fmt.Println("Part "+strconv.Itoa(i+1)+":", dayResult)
		}
	}
}
