package day12

import (
	"fmt"
	"robertbrignull/adventofcode2023/shared"
	"strconv"
	"strings"
)

func copySprings(springs []Spring, newValue Spring, index int) []Spring {
	newSprings := make([]Spring, len(springs))
	copy(newSprings, springs)
	newSprings[index] = newValue
	return newSprings
}

func copyBrokenGroups(brokenGroups []int, newValue int, index int) []int {
	newBrokenGroups := make([]int, len(brokenGroups))
	copy(newBrokenGroups, brokenGroups)
	newBrokenGroups[index] = newValue
	return newBrokenGroups
}

type Spring int

const (
	Working Spring = iota
	Broken
	Unknown
)

func readSpring(b byte) (Spring, error) {
	if b == '.' {
		return Working, nil
	} else if b == '#' {
		return Broken, nil
	} else if b == '?' {
		return Unknown, nil
	} else {
		return Unknown, fmt.Errorf("Unrecognised character: '%c'", b)
	}
}

func printSpring(s Spring) string {
	switch s {
	case Working:
		return "."
	case Broken:
		return "#"
	case Unknown:
		return "?"
	}
	panic(fmt.Sprintf("Unhandled spring %v", s))
}

func printSprings(ss []Spring) string {
	v := ""
	for _, s := range ss {
		v += printSpring(s)
	}
	return v
}

func printBrokenGroups(brokenGroups []int) string {
	v := ""
	for i, x := range brokenGroups {
		v += strconv.Itoa(x)
		if i < len(brokenGroups)-1 {
			v += ","
		}
	}
	return v
}

type SpringRow struct {
	springs      []Spring
	brokenGroups []int
}

func readSpringRow(line string) (SpringRow, error) {
	springsPart := strings.Split(line, " ")[0]
	springs := make([]Spring, len(springsPart))
	for i := range springsPart {
		s, err := readSpring(line[i])
		if err != nil {
			return SpringRow{}, err
		}
		springs[i] = s
	}

	groupsParts := strings.Split(strings.Split(line, " ")[1], ",")
	brokenGroups := make([]int, len(groupsParts))
	for i, s := range groupsParts {
		x, err := strconv.Atoi(s)
		if err != nil {
			return SpringRow{}, err
		}
		brokenGroups[i] = x
	}

	return SpringRow{springs, brokenGroups}, nil
}

func readSpringRows(lines []string) ([]SpringRow, error) {
	springRows := make([]SpringRow, len(lines))
	for i, line := range lines {
		springRow, err := readSpringRow(line)
		if err != nil {
			return []SpringRow{}, err
		}
		springRows[i] = springRow
	}
	return springRows, nil
}

func (r SpringRow) countPossibleCombinations(print bool, indent string) int {
	numWorking := 0
	for _, s := range r.springs {
		if s == Working {
			numWorking += 1
		} else {
			break
		}
	}
	if numWorking > 0 {
		r2 := SpringRow{
			springs:      r.springs[numWorking:],
			brokenGroups: r.brokenGroups,
		}
		v := r2.countPossibleCombinations(print, indent+"  ")
		if print {
			fmt.Printf("%s'%s' '%s' => %d\n", indent, printSprings(r.springs), printBrokenGroups(r.brokenGroups), v)
		}
		return v
	}

	if len(r.brokenGroups) == 0 {
		numBroken := 0
		for _, s := range r.springs {
			if s == Broken {
				numBroken += 1
			} else {
				break
			}
		}

		if numBroken > 0 {
			if print {
				fmt.Printf("%s'%s' '%s' => 0\n", indent, printSprings(r.springs), printBrokenGroups(r.brokenGroups))
			}
			return 0
		}
		if print {
			fmt.Printf("%s'%s' '%s' => 1\n", indent, printSprings(r.springs), printBrokenGroups(r.brokenGroups))
		}
		return 1
	}

	group := r.brokenGroups[0]

	if len(r.springs) < group {
		if print {
			fmt.Printf("%s'%s' '%s' => 0\n", indent, printSprings(r.springs), printBrokenGroups(r.brokenGroups))
		}
		return 0
	}

	numNonWorking := 0
	for _, s := range r.springs {
		if s != Working {
			numNonWorking += 1
		} else {
			break
		}
	}

	if r.springs[0] == Broken {
		if group == numNonWorking && group == len(r.springs) {
			r2 := SpringRow{
				springs:      r.springs[group:],
				brokenGroups: r.brokenGroups[1:],
			}
			v := r2.countPossibleCombinations(print, indent+"  ")
			if print {
				fmt.Printf("%s'%s' '%s' => %d\n", indent, printSprings(r.springs), printBrokenGroups(r.brokenGroups), v)
			}
			return v
		}

		if group > numNonWorking {
			if print {
				fmt.Printf("%s'%s' '%s' => 0\n", indent, printSprings(r.springs), printBrokenGroups(r.brokenGroups))
			}
			return 0
		}

		if group < numNonWorking && r.springs[group] == Broken {
			if print {
				fmt.Printf("%s'%s' '%s' => 0\n", indent, printSprings(r.springs), printBrokenGroups(r.brokenGroups))
			}
			return 0
		}

		r2 := SpringRow{
			springs:      r.springs[group+1:],
			brokenGroups: r.brokenGroups[1:],
		}
		v := r2.countPossibleCombinations(print, indent+"  ")
		if print {
			fmt.Printf("%s'%s' '%s' => %d\n", indent, printSprings(r.springs), printBrokenGroups(r.brokenGroups), v)
		}
		return v
	}

	r2 := SpringRow{
		springs:      copySprings(r.springs, Working, 0),
		brokenGroups: r.brokenGroups,
	}
	r3 := SpringRow{
		springs:      copySprings(r.springs, Broken, 0),
		brokenGroups: r.brokenGroups,
	}
	v := r2.countPossibleCombinations(print, indent+"W ") + r3.countPossibleCombinations(print, indent+"B ")
	if print {
		fmt.Printf("%s'%s' '%s' => %d\n", indent, printSprings(r.springs), printBrokenGroups(r.brokenGroups), v)
	}
	return v
}

// Time taken: 15:15-15:34, 17:11-18:49
func Part1() (string, error) {
	lines, err := shared.ReadFileLines("days/day12/input.txt")
	if err != nil {
		return "", err
	}

	springRows, err := readSpringRows(lines)
	if err != nil {
		return "", err
	}

	total := 0
	for i, r := range springRows {
		v := r.countPossibleCombinations(false, "")
		fmt.Printf("%s - %d arrangements\n", lines[i], v)
		total += v
	}

	return strconv.Itoa(total), nil
}
