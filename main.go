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

	day := dayPkg.DaysList.FindByNumber(uint(dayNumber))
	result := day.Run(string(content))

	fmt.Println(result)
}
