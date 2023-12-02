package day2

import (
	"fmt"
	"robertbrignull/adventofcode2023/shared"
	"strconv"
	"strings"
)

type hand struct {
	red   int
	green int
	blue  int
}

type game struct {
	id    int
	hands []hand
}

func extractHand(line string) (hand, error) {
	commaParts := strings.Split(line, ", ")

	h := hand{}
	for _, commaPart := range commaParts {
		spaceParts := strings.Split(commaPart, " ")

		n, err := strconv.Atoi(spaceParts[0])
		if err != nil {
			return hand{}, err
		}

		if spaceParts[1] == "red" {
			h.red = n
		} else if spaceParts[1] == "green" {
			h.green = n
		} else if spaceParts[1] == "blue" {
			h.blue = n
		} else {
			return hand{}, fmt.Errorf("Unrecognised colour: %s", spaceParts[1])
		}
	}

	return h, nil
}

func extractHands(line string) ([]hand, error) {
	parts := strings.Split(line, "; ")

	var hands []hand
	for _, part := range parts {
		h, err := extractHand(part)
		if err != nil {
			return []hand{}, err
		}
		hands = append(hands, h)
	}

	return hands, nil
}

func extractGameInfo(line string) (game, error) {
	parts := strings.Split(line, ": ")

	id, err := strconv.Atoi(parts[0][5:])
	if err != nil {
		return game{}, err
	}

	hands, err := extractHands(parts[1])
	if err != nil {
		return game{}, err
	}

	return game{id, hands}, nil
}

func isGamePossible(g game, maxRed int, maxGreen int, maxBlue int) bool {
	for _, h := range g.hands {
		if h.red > maxRed || h.green > maxGreen || h.blue > maxBlue {
			return false
		}
	}
	return true
}

// Time taken: 19 minutes
func Part1() (string, error) {
	lines, err := shared.ReadFileLines("days/day2/input2.txt")
	if err != nil {
		return "", err
	}

	var games []game
	for _, line := range lines {
		g, err := extractGameInfo(line)
		if err != nil {
			return "", err
		}
		games = append(games, g)
	}

	sum := 0
	for _, g := range games {
		if isGamePossible(g, 12, 13, 14) {
			sum += g.id
		}
	}

	return strconv.Itoa(sum), nil
}
