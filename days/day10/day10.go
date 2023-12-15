package day10

import (
	"fmt"
	"robertbrignull/adventofcode2023/shared"
	"strconv"
)

type Direction int

const (
	N Direction = iota
	S
	E
	W
)

func (d Direction) opposite() Direction {
	switch d {
	case N:
		return S
	case S:
		return N
	case E:
		return W
	case W:
		return E
	}
	panic(fmt.Sprintf("Unhandled direction %v", d))
}

type Pipe int

const (
	None Pipe = iota
	Start
	NS
	EW
	NE
	ES
	SW
	WN
)

func readPipe(b byte) (Pipe, error) {
	if b == '.' {
		return None, nil
	} else if b == 'S' {
		return Start, nil
	} else if b == '|' {
		return NS, nil
	} else if b == '-' {
		return EW, nil
	} else if b == 'L' {
		return NE, nil
	} else if b == 'F' {
		return ES, nil
	} else if b == '7' {
		return SW, nil
	} else if b == 'J' {
		return WN, nil
	} else {
		return None, fmt.Errorf("Uknown character '%c'", b)
	}
}

func (p Pipe) connectsFrom(d Direction) bool {
	switch p {
	case NS:
		return d == N || d == S
	case EW:
		return d == E || d == W
	case NE:
		return d == N || d == E
	case ES:
		return d == E || d == S
	case SW:
		return d == S || d == W
	case WN:
		return d == W || d == N
	}
	return false
}

func (p Pipe) getConnectingDirection() (Direction, error) {
	switch p {
	case NS:
		return N, nil
	case EW:
		return E, nil
	case NE:
		return N, nil
	case ES:
		return E, nil
	case SW:
		return S, nil
	case WN:
		return W, nil
	}
	return N, fmt.Errorf("Pipe %v doesn't connect in any directions", p)
}

type PipeField [][]Pipe

func readPipeField(lines []string) (PipeField, error) {
	pf := make([][]Pipe, len(lines))
	for r, line := range lines {
		pl := make([]Pipe, len(line))
		for c := range line {
			p, err := readPipe(line[c])
			if err != nil {
				return [][]Pipe{}, err
			}
			pl[c] = p
		}
		pf[r] = pl
	}
	return pf, nil
}

func (pf PipeField) findStart() (int, int, error) {
	for y, line := range pf {
		for x, p := range line {
			if p == Start {
				return x, y, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("Unable to find start in pipe field")
}

func (pf PipeField) determineStartPipe(x int, y int) (Pipe, error) {
	connectsN := y > 0 && pf[y-1][x].connectsFrom(S)
	connectsS := y < len(pf)-1 && pf[y+1][x].connectsFrom(N)
	connectsE := x < len(pf[0])-1 && pf[y][x+1].connectsFrom(W)
	connectsW := x > 0 && pf[y][x-1].connectsFrom(E)

	if connectsN && !connectsE && connectsS && !connectsW {
		return NS, nil
	} else if !connectsN && connectsE && !connectsS && connectsW {
		return EW, nil
	} else if connectsN && connectsE && !connectsS && connectsW {
		return NE, nil
	} else if !connectsN && connectsE && connectsS && !connectsW {
		return ES, nil
	} else if !connectsN && !connectsE && connectsS && connectsW {
		return SW, nil
	} else if connectsN && !connectsE && !connectsS && connectsW {
		return WN, nil
	} else {
		return None, fmt.Errorf("Pipe at (%d, %d) does not connect to exactly two other pipes", x, y)
	}
}

func (pf PipeField) replaceStartPipe(x int, y int) error {
	p, err := pf.determineStartPipe(x, y)
	if err != nil {
		return err
	}
	pf[y][x] = p
	return nil
}

func (pf PipeField) findLengthOfPipeLoop(startX int, startY int) (int, error) {
	x := startX
	y := startY

	prevDirection, err := pf[startY][startX].getConnectingDirection()
	if err != nil {
		return 0, err
	}

	loopLength := 0

	for {
		p := pf[y][x]
		if p.connectsFrom(N) && prevDirection != N {
			y -= 1
			prevDirection = N.opposite()
		} else if p.connectsFrom(E) && prevDirection != E {
			x += 1
			prevDirection = E.opposite()
		} else if p.connectsFrom(S) && prevDirection != S {
			y += 1
			prevDirection = S.opposite()
		} else if p.connectsFrom(W) && prevDirection != W {
			x -= 1
			prevDirection = W.opposite()
		}

		loopLength += 1
		if x == startX && y == startY {
			return loopLength, nil
		}
	}
}

// Time taken: 39 minutes
func Part1() (string, error) {
	lines, err := shared.ReadFileLines("days/day10/input.txt")
	if err != nil {
		return "", err
	}

	pipeField, err := readPipeField(lines)
	if err != nil {
		return "", err
	}

	startX, startY, err := pipeField.findStart()
	if err != nil {
		return "", err
	}

	err = pipeField.replaceStartPipe(startX, startY)
	if err != nil {
		return "", err
	}

	loopLength, err := pipeField.findLengthOfPipeLoop(startX, startY)
	if err != nil {
		return "", err
	}

	farthestDistance := loopLength / 2

	return strconv.Itoa(farthestDistance), nil
}
