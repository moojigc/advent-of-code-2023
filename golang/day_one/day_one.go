package dayone

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calibrator interface {
	FindIntegers(characters []string) int
}

func DayOne(filename string, calibrator Calibrator) (int, error) {
	fmt.Printf("Reading %s...\n", filename)
	if filename == "" {
		return 0, fmt.Errorf("please provide a filename")
	}

	lines, err := readLines(filename)

	if err != nil {
		panic(err)
	}

	var answer int

	for _, line := range lines {
		characters := strings.Split(line, "")
		intValue := calibrator.FindIntegers(characters)

		fmt.Printf("From %s got %d\n", line, intValue)

		answer += intValue
	}

	return answer, nil
}

type DayOneCalibrator struct{}

func (c DayOneCalibrator) FindIntegers(characters []string) int {
	var integerString []string
	var integerValue int

	for _, char := range characters {
		_, err := strconv.Atoi(char)

		if err == nil {
			integerString = append(integerString, char)
		}
	}

	switch len(integerString) {
	case 0:
		return 0
	case 1:
		integerValue, _ = strconv.Atoi(integerString[0] + integerString[0])
		return integerValue
	default:
		integerValue, _ = strconv.Atoi(integerString[0] + integerString[len(integerString)-1])
	}

	return integerValue
}

func readLines(filename string) ([]string, error) {
	inputFile, err := os.Open(filename)

	if err != nil {
		empty := make([]string, 0)
		return empty, err
	}

	defer inputFile.Close()

	var lines []string
	var scanner *bufio.Scanner = bufio.NewScanner(inputFile)

	for {
		if !scanner.Scan() {
			break
		}

		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
