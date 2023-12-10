package main

import (
	"fmt"
	"os"
	"strconv"

	dayone "github.com/moojigc/advent-of-code-2023/day_one"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) < 3 {
		fmt.Println("Usage: aoc [DAY# (1|2|3...)] [INPUT_FILE_PATH]")
		return
	}

	inputFile := os.Args[2]
	day, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Day must be an integer value from 1-25.")
	}

	switch day {
	case 1:
		answer, err := dayone.DayOne(inputFile, dayone.DayOneCalibrator{})
		if err != nil {
			fmt.Printf("got err from dayone.Dayone: %v\n", err)
			return
		}

		fmt.Printf("DAY ONE: The answer is... %d!\n", answer)
	}
}
