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

type DFAState struct {
	node         string
	index        int
	nextDestNode string
	stepsToDest  int
}

type DFA struct {
	nodes       []string
	destNodes   []string
	transitions map[string]DFAState
	numIndexes  int
}

func (d DFA) key(node string, index int) string {
	return fmt.Sprintf("%s%d", node, index)
}

func (d DFA) decodeKey(key string) (string, int, error) {
	node := key[0:3]
	index, err := strconv.Atoi(key[3:])
	if err != nil {
		return "", 0, err
	}
	return node, index, nil
}

func (d DFA) backwardsTransitions() (map[string][]string, error) {
	backwardTransitions := make(map[string][]string)
	for k, v := range d.transitions {
		key := d.key(v.node, v.index)
		if _, ok := backwardTransitions[key]; !ok {
			backwardTransitions[key] = make([]string, 0)
		}
		backwardTransitions[key] = append(backwardTransitions[key], k)
	}
	return backwardTransitions, nil
}

func (d DFA) computeStepsToDest() error {
	keysToProcess := make([]string, 0)
	for _, destNode := range d.destNodes {
		for index := 0; index < d.numIndexes; index++ {
			key := d.key(destNode, index)
			keysToProcess = append(keysToProcess, key)

			t := d.transitions[key]
			t.nextDestNode = destNode
			t.stepsToDest = 0
			d.transitions[key] = t
		}
	}

	backwardTransitions, err := d.backwardsTransitions()
	if err != nil {
		return err
	}

	for {
		if len(keysToProcess) == 0 {
			return nil
		}
		key := keysToProcess[0]
		keysToProcess = keysToProcess[1:]

		transition := d.transitions[key]

		prevKeys, foundPrevKeys := backwardTransitions[key]
		if foundPrevKeys {
			for _, prevKey := range prevKeys {
				prevTransition := d.transitions[prevKey]
				if prevTransition.stepsToDest == -1 {
					prevTransition.stepsToDest = transition.stepsToDest + 1
					prevTransition.nextDestNode = transition.nextDestNode
					keysToProcess = append(keysToProcess, prevKey)
					d.transitions[prevKey] = prevTransition
				}
			}
		}
	}
}

func getAllNodes(branches Branches) []string {
	s := make(map[string]bool)
	for k, v := range branches {
		s[k] = true
		s[v.left] = true
		s[v.right] = true
	}

	nodes := make([]string, 0, len(s))
	for k := range s {
		nodes = append(nodes, k)
	}
	return nodes
}

func getDestNodes(nodes []string, ghostMode bool) []string {
	destNodes := make([]string, 0)
	for _, node := range nodes {
		if (ghostMode && node[2] == 'Z') || node == "ZZZ" {
			destNodes = append(destNodes, node)
		}
	}
	return destNodes
}

func getNextNode(instructions []Instruction, branches Branches, node string, index int) (string, error) {
	branch, ok := branches[node]
	if !ok {
		return "", fmt.Errorf("No branches defined for node %s", node)
	}

	instruction := instructions[index]
	if instruction == Left {
		return branch.left, nil
	} else {
		return branch.right, nil
	}
}

func buildDFA(instructions []Instruction, branches Branches, ghostMode bool) (DFA, error) {
	nodes := getAllNodes(branches)
	transitions := make(map[string]DFAState)
	numIndexes := len(instructions)
	destNodes := getDestNodes(nodes, ghostMode)

	dfa := DFA{nodes, destNodes, transitions, numIndexes}

	for index := 0; index < numIndexes; index++ {
		for _, node := range nodes {
			nextNode, err := getNextNode(instructions, branches, node, index)
			if err != nil {
				return DFA{}, err
			}
			dfa.transitions[dfa.key(node, index)] = DFAState{
				node:         nextNode,
				index:        (index + 1) % numIndexes,
				nextDestNode: "",
				stepsToDest:  -1,
			}
		}
	}

	err := dfa.computeStepsToDest()
	if err != nil {
		return DFA{}, err
	}

	return dfa, nil
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

	dfa, err := buildDFA(instructions, bs, false)
	if err != nil {
		return "", err
	}

	result := dfa.transitions[dfa.key("AAA", 0)]
	if result.stepsToDest == -1 {
		return "", fmt.Errorf("No route from AAA to dest node")
	}

	return strconv.Itoa(result.stepsToDest), nil
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

	dfa, err := buildDFA(instructions, bs, false)
	if err != nil {
		return 0, err
	}

	for {
		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i].steps < nodes[j].steps
		})

		if nodes[0].steps == nodes[len(nodes)-1].steps && nodes[0].steps != 0 {
			return nodes[0].steps, nil
		}

		result := dfa.transitions[dfa.key(nodes[0].node, nodes[0].index)]
		if result.stepsToDest == -1 {
			return 0, fmt.Errorf("No route from %s to dest node", nodes[0].node)
		}

		fmt.Printf("Advanced %s at step %d to %s at step %d\n", nodes[0].node, nodes[0].steps, result.nextDestNode, nodes[0].steps+result.stepsToDest)
		nodes[0] = GhostNode{
			node:  result.nextDestNode,
			steps: nodes[0].steps + result.stepsToDest,
			index: result.index,
		}
	}
}

// Time taken: 2h 01m and I think confirmed unfinishable with my problem input :(
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
