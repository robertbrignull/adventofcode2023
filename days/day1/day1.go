package day1

import (
	"fmt"
	"robertbrignull/adventofcode2023/shared"
	"strconv"
)

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func decodeNumber(str string) int {
	if len(str) >= 3 && str[0:3] == "one" {
		return 1
	} else if len(str) >= 3 && str[0:3] == "two" {
		return 2
	} else if len(str) >= 5 && str[0:5] == "three" {
		return 3
	} else if len(str) >= 4 && str[0:4] == "four" {
		return 4
	} else if len(str) >= 4 && str[0:4] == "five" {
		return 5
	} else if len(str) >= 3 && str[0:3] == "six" {
		return 6
	} else if len(str) >= 5 && str[0:5] == "seven" {
		return 7
	} else if len(str) >= 5 && str[0:5] == "eight" {
		return 8
	} else if len(str) >= 4 && str[0:4] == "nine" {
		return 9
	} else {
		return 0
	}
}

func firstDigit(line string) (int, error) {
	for i := range line {
		if isDigit(line[i]) {
			return strconv.Atoi(line[i : i+1])
		}
	}
	return 0, fmt.Errorf("string does not contain any digits: %s", line)
}

func firstNumber(line string) (int, error) {
	for i := range line {
		if isDigit(line[i]) {
			return strconv.Atoi(line[i : i+1])
		}

		n := decodeNumber(line[i:])
		if n != 0 {
			return n, nil
		}
	}
	return 0, fmt.Errorf("string does not contain any digits: %s", line)
}

func lastDigit(line string) (int, error) {
	for j := range line {
		i := len(line) - 1 - j
		if isDigit(line[i]) {
			return strconv.Atoi(line[i : i+1])
		}
	}
	return 0, fmt.Errorf("string does not contain any digits: %s", line)
}

func lastNumber(line string) (int, error) {
	for j := range line {
		i := len(line) - 1 - j
		if isDigit(line[i]) {
			return strconv.Atoi(line[i : i+1])
		}

		n := decodeNumber(line[i:])
		if n != 0 {
			return n, nil
		}
	}
	return 0, fmt.Errorf("string does not contain any digits: %s", line)
}

// Time taken: 30 minutes
func Part1() (string, error) {
	lines, err := shared.ReadFileLines("days/day1/input.txt")
	if err != nil {
		return "", err
	}

	sum := 0
	for _, line := range lines {
		f, err := firstDigit(line)
		if err != nil {
			return "", err
		}

		l, err := lastDigit(line)
		if err != nil {
			return "", err
		}

		sum += f*10 + l
	}

	return strconv.Itoa(sum), nil
}

// Time taken: 11 minutes
func Part2() (string, error) {
	lines, err := shared.ReadFileLines("days/day1/input.txt")
	if err != nil {
		return "", err
	}

	sum := 0
	for _, line := range lines {
		f, err := firstNumber(line)
		if err != nil {
			return "", err
		}

		l, err := lastNumber(line)
		if err != nil {
			return "", err
		}

		sum += f*10 + l
	}

	return strconv.Itoa(sum), nil
}
