package Day1

import (
	"fmt"
	Shared "robertbrignull/adventofcode2023/shared"
	"strconv"
)

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func firstNumber(line string) (int, error) {
	for i := range line {
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
	}
	return 0, fmt.Errorf("string does not contain any digits: %s", line)
}

// Time taken 30 minutes
func Part1() (string, error) {
	lines, err := Shared.ReadFileLines("days/day1/input1.txt")
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
