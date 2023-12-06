package day3

import (
	"robertbrignull/adventofcode2023/shared"
	"strconv"
)

type PartNumber struct {
	number int
	row    int
	s      int
	e      int
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isSymbol(c byte) bool {
	return c != '.' && !isDigit(c)
}

func findNextDigit(line string, index int) int {
	for i := index; i < len(line); i++ {
		if isDigit(line[i]) {
			return i
		}
	}
	return len(line)
}

func findNextNonDigit(line string, index int) int {
	for i := index; i < len(line); i++ {
		if !isDigit(line[i]) {
			return i
		}
	}
	return len(line)
}

func isNextToSymbol(lines []string, row int, s int, e int) bool {
	for j := max(0, row-1); j < min(row+2, len(lines)); j++ {
		for i := max(0, s-1); i < min(e+1, len(lines[j])); i++ {
			if isSymbol(lines[j][i]) {
				return true
			}
		}
	}
	return false
}

func extractPartNumbers(lines []string) ([]PartNumber, error) {
	partNumbers := []PartNumber{}

	x, y := 0, 0
	for {
		// If we've reached the end of the line then move to the next line
		if x >= len(lines[y]) {
			x = 0
			y += 1
		}

		// If we've processed all lines then abort
		if y >= len(lines) {
			break
		}

		// Find the next digit in the current line
		s := findNextDigit(lines[y], x)
		x = s

		if s < len(lines[y]) {
			e := findNextNonDigit(lines[y], s+1)
			x = e

			partNumber, err := strconv.Atoi(lines[y][s:e])
			if err != nil {
				return []PartNumber{}, err
			}

			partNumbers = append(partNumbers, PartNumber{partNumber, y, s, e})
		}
	}

	return partNumbers, nil
}

// Time taken: 16 minutes
func Part1() (string, error) {
	lines, err := shared.ReadFileLines("days/day3/input.txt")
	if err != nil {
		return "", err
	}

	partNumbers, err := extractPartNumbers(lines)
	if err != nil {
		return "", err
	}

	partNumbersSum := 0
	for _, partNumber := range partNumbers {
		if isNextToSymbol(lines, partNumber.row, partNumber.s, partNumber.e) {
			partNumbersSum += partNumber.number
		}
	}

	return strconv.Itoa(partNumbersSum), nil
}
