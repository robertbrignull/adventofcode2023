package day6

import (
	"fmt"
	"math"
	"robertbrignull/adventofcode2023/shared"
	"strconv"
	"strings"
)

func readIntFields(line string) ([]int, error) {
	values := []int{}
	for _, str := range strings.Fields(line)[1:] {
		value, err := strconv.Atoi(str)
		if err != nil {
			return []int{}, err
		}
		values = append(values, value)
	}
	return values, nil
}

func computeNumWaysToWin(time int, recordDistance int) (int, error) {
	// r = h * (t - h) = h.t - h^2
	// =>  h^2 - t.h + r = 0
	// =>  (h - t/2)^2 - (t^2)/4 + r = 0
	// =>  (h - t/2)^2 = (t^2)/4 - r
	// =>  h = t/2 +- sqrt((t^2)/4 - r)
	// number of solutions = floor(sqrt((t^2)/4 - r)) * 2

	v := float64(time*time)/4 - float64(recordDistance)
	if v < 0 {
		return 0, fmt.Errorf("Unable to reach distance %d in time %d", recordDistance, time)
	}
	s := math.Sqrt(v)
	return int(math.Floor(float64(time)/2+s)) - int(math.Ceil(float64(time)/2-s)) + 1, nil
}

// Time taken: 53 minutes
func Part1() (string, error) {
	lines, err := shared.ReadFileLines("days/day6/input.txt")
	if err != nil {
		return "", err
	}

	times, err := readIntFields(lines[0])
	if err != nil {
		return "", err
	}

	distances, err := readIntFields(lines[1])
	if err != nil {
		return "", err
	}

	result := 1
	for i := range times {
		r, err := computeNumWaysToWin(times[i], distances[i]+1)
		if err != nil {
			return "", err
		}
		result *= r
	}

	return strconv.Itoa(result), nil
}
