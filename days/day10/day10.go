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

func allDirections() []Direction {
	return []Direction{N, S, E, W}
}

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

type Coord struct {
	x int
	y int
}

func (c Coord) moveDirection(d Direction) Coord {
	switch d {
	case N:
		return Coord{x: c.x, y: c.y - 1}
	case S:
		return Coord{x: c.x, y: c.y + 1}
	case E:
		return Coord{x: c.x + 1, y: c.y}
	case W:
		return Coord{x: c.x - 1, y: c.y}
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

func (p Pipe) getInsideAndOutsideCoords(c Coord) (Coord, Coord, error) {
	switch p {
	case NS:
		return c, c.moveDirection(E), nil
	case EW:
		return c, c.moveDirection(S), nil
	case NE:
		return c, c.moveDirection(E), nil
	case ES:
		return c.moveDirection(E), c.moveDirection(E).moveDirection(S), nil
	case SW:
		return c, c.moveDirection(S), nil
	case WN:
		return c, c.moveDirection(E), nil
	}
	return c, c, fmt.Errorf("Pipe %v cannot have inside and outside coords", p)
}

func (p Pipe) isMovementOutofTileBlocked(directionMoving Direction) bool {
	switch directionMoving {
	case N:
		return true
	case E:
		return p.connectsFrom(N)
	case S:
		return p.connectsFrom(W)
	case W:
		return true
	}
	panic(fmt.Sprintf("Unhandled direction %v", directionMoving))
}

func (p Pipe) isMovementIntoTileBlocked(directionMoving Direction) bool {
	switch directionMoving {
	case N:
		return p.connectsFrom(W)
	case E:
		return true
	case S:
		return true
	case W:
		return p.connectsFrom(N)
	}
	panic(fmt.Sprintf("Unhandled direction %v", directionMoving))
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

func (pf PipeField) getPipe(c Coord) Pipe {
	return pf[c.y][c.x]
}

func (pf PipeField) findStart() (Coord, error) {
	for y, line := range pf {
		for x, p := range line {
			if p == Start {
				return Coord{x, y}, nil
			}
		}
	}
	return Coord{}, fmt.Errorf("Unable to find start in pipe field")
}

func (pf PipeField) determineStartPipe(c Coord) (Pipe, error) {
	connectsN := c.y > 0 && pf.getPipe(c.moveDirection(N)).connectsFrom(S)
	connectsS := c.y < len(pf)-1 && pf.getPipe(c.moveDirection(S)).connectsFrom(N)
	connectsE := c.x < len(pf[0])-1 && pf.getPipe(c.moveDirection(E)).connectsFrom(W)
	connectsW := c.x > 0 && pf.getPipe(c.moveDirection(W)).connectsFrom(E)

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
		return None, fmt.Errorf("Pipe at (%d, %d) does not connect to exactly two other pipes", c.x, c.y)
	}
}

func (pf PipeField) replaceStartPipe(c Coord) error {
	p, err := pf.determineStartPipe(c)
	if err != nil {
		return err
	}
	pf[c.y][c.x] = p
	return nil
}

func (pf PipeField) findTilesOnLoop(start Coord) (map[Coord]bool, error) {
	loopTiles := make(map[Coord]bool, 0)

	c := start

	prevDirection, err := pf.getPipe(start).getConnectingDirection()
	if err != nil {
		return map[Coord]bool{}, err
	}

	for {
		p := pf.getPipe(c)
		for _, d := range allDirections() {
			if p.connectsFrom(d) && prevDirection != d {
				c = c.moveDirection(d)
				prevDirection = d.opposite()
				break
			}
		}

		loopTiles[c] = true

		if c == start {
			return loopTiles, nil
		}
	}
}

func (pf PipeField) findLengthOfPipeLoop(start Coord) (int, error) {
	loopTiles, err := pf.findTilesOnLoop(start)
	if err != nil {
		return 0, err
	}
	return len(loopTiles), nil
}

func (pf PipeField) cleanTilesNotOnLoop(start Coord) error {
	loopTiles, err := pf.findTilesOnLoop(start)
	if err != nil {
		return err
	}

	for y := range pf {
		for x := range pf[y] {
			if _, ok := loopTiles[Coord{x, y}]; !ok {
				pf[y][x] = None
			}
		}
	}

	return nil
}

func (pf PipeField) isCoordInsideLoop(start Coord, loopTiles map[Coord]bool) (bool, int, error) {
	coordsToProcess := make([]Coord, 1)
	coordsToProcess[0] = start

	coordsSeen := make(map[Coord]bool)

	tilesContained := make(map[Coord]bool)

	for {
		if len(coordsToProcess) == 0 {
			return true, len(tilesContained), nil
		}

		c := coordsToProcess[0]
		coordsToProcess = coordsToProcess[1:]

		if _, ok := coordsSeen[c]; ok {
			continue
		}
		coordsSeen[c] = true

		if c.x < 0 || c.x >= len(pf[0]) || c.y < 0 || c.y >= len(pf) {
			return false, 0, nil
		}

		if pf.getPipe(c) == None {
			tilesContained[c] = true
		}

		if c.y > 0 && !pf.getPipe(c.moveDirection(N)).isMovementIntoTileBlocked(N) {
			coordsToProcess = append(coordsToProcess, c.moveDirection(N))
		}
		if !pf.getPipe(c).isMovementOutofTileBlocked(E) {
			coordsToProcess = append(coordsToProcess, c.moveDirection(E))
		}
		if !pf.getPipe(c).isMovementOutofTileBlocked(S) {
			coordsToProcess = append(coordsToProcess, c.moveDirection(S))
		}
		if c.x > 0 && !pf.getPipe(c.moveDirection(W)).isMovementIntoTileBlocked(W) {
			coordsToProcess = append(coordsToProcess, c.moveDirection(W))
		}
	}
}

func (pf PipeField) findAreaEnclosedByPipeLoop(start Coord) (int, error) {
	loopTiles, err := pf.findTilesOnLoop(start)
	if err != nil {
		return 0, err
	}

	coordA, coordB, err := pf.getPipe(start).getInsideAndOutsideCoords(start)
	if err != nil {
		return 0, err
	}

	isInsideA, tilesContainedA, err := pf.isCoordInsideLoop(coordA, loopTiles)
	if err != nil {
		return 0, err
	}

	isInsideB, tilesContainedB, err := pf.isCoordInsideLoop(coordB, loopTiles)
	if err != nil {
		return 0, err
	}

	if isInsideA {
		return tilesContainedA, nil
	} else if isInsideB {
		return tilesContainedB, nil
	} else {
		return 0, fmt.Errorf("Both sides appear to be outside of the loop")
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

	start, err := pipeField.findStart()
	if err != nil {
		return "", err
	}

	err = pipeField.replaceStartPipe(start)
	if err != nil {
		return "", err
	}

	loopLength, err := pipeField.findLengthOfPipeLoop(start)
	if err != nil {
		return "", err
	}

	farthestDistance := loopLength / 2

	return strconv.Itoa(farthestDistance), nil
}

// Time taken: 59 minutes
func Part2() (string, error) {
	lines, err := shared.ReadFileLines("days/day10/input.txt")
	if err != nil {
		return "", err
	}

	pipeField, err := readPipeField(lines)
	if err != nil {
		return "", err
	}

	start, err := pipeField.findStart()
	if err != nil {
		return "", err
	}

	err = pipeField.replaceStartPipe(start)
	if err != nil {
		return "", err
	}

	err = pipeField.cleanTilesNotOnLoop(start)
	if err != nil {
		return "", err
	}

	areaEnclosed, err := pipeField.findAreaEnclosedByPipeLoop(start)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(areaEnclosed), nil
}
