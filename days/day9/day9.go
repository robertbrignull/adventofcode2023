package day9

import (
	"robertbrignull/adventofcode2023/shared"
	"strconv"
	"strings"
)

func readSequence(line string) ([]int, error) {
	xs := make([]int, 0)
	for _, v := range strings.Fields(line) {
		x, err := strconv.Atoi(v)
		if err != nil {
			return xs, err
		}
		xs = append(xs, x)
	}
	return xs, nil
}

func readSequences(lines []string) ([][]int, error) {
	ss := make([][]int, 0)
	for _, line := range lines {
		s, err := readSequence(line)
		if err != nil {
			return ss, err
		}
		ss = append(ss, s)
	}
	return ss, nil
}

func computeDeltas(xs []int) []int {
	ds := make([]int, len(xs)-1)
	for i := 0; i < len(xs)-1; i++ {
		ds[i] = xs[i+1] - xs[i]
	}
	return ds
}

func isAllZeros(xs []int) bool {
	for _, x := range xs {
		if x != 0 {
			return false
		}
	}
	return true
}

func nextValueInSequence(xs []int) int {
	if isAllZeros(xs) {
		return 0
	} else {
		return xs[len(xs)-1] + nextValueInSequence(computeDeltas(xs))
	}
}

func prevValueInSequence(xs []int) int {
	if isAllZeros(xs) {
		return 0
	} else {
		return xs[0] - prevValueInSequence(computeDeltas(xs))
	}
}

// Time taken: 15 minutes
func Part1() (string, error) {
	lines, err := shared.ReadFileLines("days/day9/input.txt")
	if err != nil {
		return "", err
	}

	sequences, err := readSequences(lines)
	if err != nil {
		return "", err
	}

	total := 0
	for _, sequence := range sequences {
		total += nextValueInSequence(sequence)
	}

	return strconv.Itoa(total), nil
}

// Time taken: 3 minutes
func Part2() (string, error) {
	lines, err := shared.ReadFileLines("days/day9/input.txt")
	if err != nil {
		return "", err
	}

	sequences, err := readSequences(lines)
	if err != nil {
		return "", err
	}

	total := 0
	for _, sequence := range sequences {
		total += prevValueInSequence(sequence)
	}

	return strconv.Itoa(total), nil
}
