package day4

import (
	"fmt"
	"regexp"
	"robertbrignull/adventofcode2023/shared"
	"strconv"
	"strings"
)

type ScratchCard struct {
	cardNumber     int
	winningNumbers []int
	yourNumbers    []int
	points         int
}

func isWinningNumber(winningNumbers []int, yourNumber int) bool {
	for _, n := range winningNumbers {
		if n == yourNumber {
			return true
		}
	}
	return false
}

func computeScratchCardPoints(winningNumbers []int, yourNumbers []int) int {
	points := 0
	for _, yn := range yourNumbers {
		if isWinningNumber(winningNumbers, yn) {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}
	return points
}

func extractNumbers(line string) ([]int, error) {
	ns := []int{}
	for _, n := range strings.Fields(line) {
		n, err := strconv.Atoi(n)
		if err != nil {
			return []int{}, err
		}
		ns = append(ns, n)
	}
	return ns, nil
}

func extractScratchCard(line string) (ScratchCard, error) {
	re := regexp.MustCompile(`Card +(\d+): ([\d ]+) \| ([\d ]+)`)
	match := re.FindStringSubmatch(line)
	if len(match) != 4 {
		return ScratchCard{}, fmt.Errorf("invalid scratch card line: %s", line)
	}

	cardNumber, err := strconv.Atoi(match[1])
	if err != nil {
		return ScratchCard{}, err
	}

	winningNumbers, err := extractNumbers(match[2])
	if err != nil {
		return ScratchCard{}, err
	}

	yourNumbers, err := extractNumbers(match[3])
	if err != nil {
		return ScratchCard{}, err
	}

	points := computeScratchCardPoints(winningNumbers, yourNumbers)

	return ScratchCard{
		cardNumber,
		winningNumbers,
		yourNumbers,
		points,
	}, nil
}

func extractScratchCards(lines []string) ([]ScratchCard, error) {
	scratchCards := []ScratchCard{}

	for _, line := range lines {
		scratchCard, err := extractScratchCard(line)
		if err != nil {
			return []ScratchCard{}, err
		}
		scratchCards = append(scratchCards, scratchCard)
	}

	return scratchCards, nil
}

// Time taken: 16 minutes
func Part1() (string, error) {
	lines, err := shared.ReadFileLines("days/day4/input.txt")
	if err != nil {
		return "", err
	}

	scratchCards, err := extractScratchCards(lines)
	if err != nil {
		return "", err
	}

	pointsSum := 0
	for _, scratchCard := range scratchCards {
		pointsSum += scratchCard.points
	}

	return strconv.Itoa(pointsSum), nil
}
