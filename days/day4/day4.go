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
	numMatches     int
}

func isWinningNumber(winningNumbers []int, yourNumber int) bool {
	for _, n := range winningNumbers {
		if n == yourNumber {
			return true
		}
	}
	return false
}

func computeScratchCardPoints(numMatches int) int {
	if numMatches == 0 {
		return 0
	}
	points := 1
	for i := 1; i < numMatches; i++ {
		points *= 2
	}
	return points
}

func computeNumMatches(winningNumbers []int, yourNumbers []int) int {
	numMatches := 0
	for _, yn := range yourNumbers {
		if isWinningNumber(winningNumbers, yn) {
			numMatches += 1
		}
	}
	return numMatches
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

	numMatches := computeNumMatches(winningNumbers, yourNumbers)

	return ScratchCard{
		cardNumber,
		winningNumbers,
		yourNumbers,
		numMatches,
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
		pointsSum += computeScratchCardPoints(scratchCard.numMatches)
	}

	return strconv.Itoa(pointsSum), nil
}

// Time taken: 11 minutes
func Part2() (string, error) {
	lines, err := shared.ReadFileLines("days/day4/input.txt")
	if err != nil {
		return "", err
	}

	scratchCards, err := extractScratchCards(lines)
	if err != nil {
		return "", err
	}

	numCopies := make([]int, len(scratchCards))
	for i := range numCopies {
		// We gain one original copy
		numCopies[i] += 1

		for j := 0; j < scratchCards[i].numMatches; j++ {
			if i+j+1 >= len(numCopies) {
				break
			}
			numCopies[i+j+1] += numCopies[i]
		}
	}

	cardsSum := 0
	for _, c := range numCopies {
		cardsSum += c
	}

	return strconv.Itoa(cardsSum), nil
}
