package day8

import (
	"fmt"
	"regexp"
	"robertbrignull/adventofcode2023/shared"
	"strconv"
)

type instruction int8

const (
	left instruction = iota
	right
)

type branch struct {
	left  string
	right string
}

func readInstruction(c byte) (instruction, error) {
	if c == 'L' {
		return left, nil
	} else if c == 'R' {
		return right, nil
	}
	return left, fmt.Errorf("'%s' is not a recognised instruction", c)
}

func readInstructions(line string) ([]instruction, error) {
	is := []instruction{}
	for i := range line {
		v, err := readInstruction(line[i])
		if err != nil {
			return []instruction{}, err
		}
		is = append(is, v)
	}
	return is, nil
}

func readBranches(lines []string) (map[string]branch, error) {
	bs := map[string]branch{}
	re := regexp.MustCompile(`(...) = \((...), (...)\)`)
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) != 4 {
			return map[string]branch{}, fmt.Errorf("line could not be parsed: %s", line)
		}
		bs[match[1]] = branch{
			left:  match[2],
			right: match[3],
		}
	}
	return bs, nil
}

func countStepsToDestination(instructions []instruction, branches map[string]branch, start string, dest string) (int, error) {
	step := 0
	index := 0
	node := start
	for {
		if node == dest {
			return step, nil
		}

		if index >= len(instructions) {
			index = 0
		}

		branch, ok := branches[node]
		if !ok {
			return 0, fmt.Errorf("No branch found for node %s", node)
		}

		if instructions[index] == left {
			node = branch.left
		} else {
			node = branch.right
		}
		step += 1
		index += 1
	}
}

// Time taken: 35 minutes
func Part1() (string, error) {
	lines, err := shared.ReadFileLines("days/day8/input.txt")
	if err != nil {
		return "", err
	}

	instructions, err := readInstructions(lines[0])
	if err != nil {
		return "", err
	}

	branches, err := readBranches(lines[2:])
	if err != nil {
		return "", err
	}

	steps, err := countStepsToDestination(instructions, branches, "AAA", "ZZZ")
	if err != nil {
		return "", err
	}

	return strconv.Itoa(steps), nil
}
