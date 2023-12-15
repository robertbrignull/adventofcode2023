package day8

import (
	"fmt"
	"regexp"
	"robertbrignull/adventofcode2023/shared"
	"sort"
	"strconv"
)

type Instruction int8

const (
	Left Instruction = iota
	Right
)

type Branch struct {
	left  string
	right string
}

type Branches map[string]Branch

func readInstruction(c byte) (Instruction, error) {
	if c == 'L' {
		return Left, nil
	} else if c == 'R' {
		return Right, nil
	}
	return Left, fmt.Errorf("'%c' is not a recognised instruction", c)
}

func readInstructions(line string) ([]Instruction, error) {
	is := []Instruction{}
	for i := range line {
		v, err := readInstruction(line[i])
		if err != nil {
			return []Instruction{}, err
		}
		is = append(is, v)
	}
	return is, nil
}

func readBranches(lines []string) (Branches, error) {
	bs := Branches{}
	re := regexp.MustCompile(`(...) = \((...), (...)\)`)
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) != 4 {
			return Branches{}, fmt.Errorf("line could not be parsed: %s", line)
		}
		bs[match[1]] = Branch{
			left:  match[2],
			right: match[3],
		}
	}
	return bs, nil
}

type StepsToNextDest struct {
	steps int
	dest  string
	index int
}

type ComputedStepsNeeded struct {
	instructions []Instruction
	bs           Branches
	storedSteps  map[string]StepsToNextDest
	ghostMode    bool
}

func (c ComputedStepsNeeded) key(node string, index int) string {
	return fmt.Sprintf("%s%d", node, index)
}

func (c ComputedStepsNeeded) isDest(node string) bool {
	if c.ghostMode {
		return node[2] == 'Z'
	} else {
		return node == "ZZZ"
	}
}

func (c ComputedStepsNeeded) getStepsToNextDest(node string, index int) (StepsToNextDest, error) {
	fmt.Printf("Computing next dest for %s at index %d\n", node, index)

	key := c.key(node, index)
	if v, ok := c.storedSteps[key]; ok {
		return v, nil
	}

	step := 0
	for {
		if c.isDest(node) {
			result := StepsToNextDest{steps: step, dest: node, index: index}
			c.storedSteps[key] = result
			return result, nil
		}

		if index >= len(c.instructions) {
			index = 0
		}

		branch, ok := c.bs[node]
		if !ok {
			return StepsToNextDest{}, fmt.Errorf("No branch found for node %s", node)
		}

		if c.instructions[index] == Left {
			node = branch.left
		} else {
			node = branch.right
		}
		step += 1
		index += 1

		if step%1000000 == 0 {
			fmt.Printf("Step = %d\n", step)
		}
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

	bs, err := readBranches(lines[2:])
	if err != nil {
		return "", err
	}

	computer := ComputedStepsNeeded{
		instructions: instructions,
		bs:           bs,
		storedSteps:  map[string]StepsToNextDest{},
		ghostMode:    false,
	}

	result, err := computer.getStepsToNextDest("AAA", 0)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(result.steps), nil
}

type GhostNode struct {
	node  string
	steps int
	index int
}

func countGhostStepsToDestination(instructions []Instruction, bs Branches) (int, error) {
	nodes := []GhostNode{}
	for node := range bs {
		if node[2] == 'A' {
			nodes = append(nodes, GhostNode{node: node, steps: 0, index: 0})
		}
	}

	computer := ComputedStepsNeeded{
		instructions: instructions,
		bs:           bs,
		storedSteps:  map[string]StepsToNextDest{},
		ghostMode:    false,
	}

	for {
		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i].steps < nodes[j].steps
		})

		if nodes[0].steps == nodes[len(nodes)-1].steps && nodes[0].steps != 0 {
			return nodes[0].steps, nil
		}

		result, err := computer.getStepsToNextDest(nodes[0].node, nodes[0].index)
		if err != nil {
			return 0, err
		}

		fmt.Printf("Advanced %s at step %d to %s at step %d\n", nodes[0].node, nodes[0].steps, result.dest, nodes[0].steps+result.steps)
		nodes[0] = GhostNode{
			node:  result.dest,
			steps: nodes[0].steps + result.steps,
			index: result.index,
		}
	}
}

// Time taken: 09:11-09:35, 09:56-10:29, 12:34
func Part2() (string, error) {
	lines, err := shared.ReadFileLines("days/day8/input.txt")
	if err != nil {
		return "", err
	}

	instructions, err := readInstructions(lines[0])
	if err != nil {
		return "", err
	}

	bs, err := readBranches(lines[2:])
	if err != nil {
		return "", err
	}

	steps, err := countGhostStepsToDestination(instructions, bs)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(steps), nil
}
