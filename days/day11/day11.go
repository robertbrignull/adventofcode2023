package day11

import (
	"robertbrignull/adventofcode2023/shared"
	"strconv"
)

type Coord struct {
	x int
	y int
}

func (c Coord) distance(d Coord) int {
	deltaX := d.x - c.x
	if deltaX < 0 {
		deltaX = -deltaX
	}

	deltaY := d.y - c.y
	if deltaY < 0 {
		deltaY = -deltaY
	}

	return deltaX + deltaY
}

type Sky struct {
	width    int
	height   int
	galaxies []Coord
}

func (s *Sky) readSky(lines []string) {
	s.width = len(lines[0])
	s.height = len(lines)
	s.galaxies = make([]Coord, 0)

	for y, line := range lines {
		for x := range line {
			if line[x] == '#' {
				s.galaxies = append(s.galaxies, Coord{x, y})
			}
		}
	}
}

func (s *Sky) expandSpace() {
	for x := 0; x < s.width; x++ {
		columnHasGalaxies := false
		for _, galaxy := range s.galaxies {
			if galaxy.x == x {
				columnHasGalaxies = true
			}
		}

		if !columnHasGalaxies {
			for i := range s.galaxies {
				if s.galaxies[i].x > x {
					s.galaxies[i].x += 1
				}
			}
			s.width += 1
			x += 1
		}
	}

	for y := 0; y < s.height; y++ {
		columnHasGalaxies := false
		for _, galaxy := range s.galaxies {
			if galaxy.y == y {
				columnHasGalaxies = true
			}
		}

		if !columnHasGalaxies {
			for i := range s.galaxies {
				if s.galaxies[i].y > y {
					s.galaxies[i].y += 1
				}
			}
			s.height += 1
			y += 1
		}
	}
}

// func (s Sky) print() {
// 	lines := make([]string, s.height)
// 	for y := range lines {
// 		lines[y] = ""
// 		for x := 0; x < s.width; x++ {
// 			lines[y] += "."
// 		}
// 	}

// 	for _, galaxy := range s.galaxies {
// 		lines[galaxy.y] = lines[galaxy.y][:galaxy.x] + "#" + lines[galaxy.y][galaxy.x+1:]
// 	}

// 	for _, line := range lines {
// 		fmt.Printf("%s\n", line)
// 	}
// 	fmt.Println()
// }

type Pair struct {
	i int
	j int
}

func (s Sky) computeGalaxyPairs() []Pair {
	pairs := make([]Pair, 0)
	for i := 0; i < len(s.galaxies); i++ {
		for j := i + 1; j < len(s.galaxies); j++ {
			pairs = append(pairs, Pair{i, j})
		}
	}
	return pairs
}

func (s Sky) computeGalaxyDistances() int {
	pairs := s.computeGalaxyPairs()

	totalDistance := 0
	for _, pair := range pairs {
		totalDistance += s.galaxies[pair.i].distance(s.galaxies[pair.j])
	}

	return totalDistance
}

// Time taken: 40 minutes
func Part1() (string, error) {
	lines, err := shared.ReadFileLines("days/day11/input.txt")
	if err != nil {
		return "", err
	}

	sky := Sky{}
	sky.readSky(lines)

	sky.expandSpace()

	totalDistance := sky.computeGalaxyDistances()

	return strconv.Itoa(totalDistance), nil
}
