package main

import (
	"fmt"
	"log"
	"os"
	"robertbrignull/adventofcode2023/days/day1"
	"robertbrignull/adventofcode2023/days/day2"
	"robertbrignull/adventofcode2023/days/day3"
	"robertbrignull/adventofcode2023/days/day4"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 && args[0] == "--" {
		args = args[1:]
	}
	if len(args) < 2 {
		log.Fatal("Usage: ./run <day> <part>\n")
	}

	day := args[0]
	part := args[1]

	var result string
	var err error
	result, err = "", nil
	if day == "1" && part == "1" {
		result, err = day1.Part1()
	} else if day == "1" && part == "2" {
		result, err = day1.Part2()
	} else if day == "2" && part == "1" {
		result, err = day2.Part1()
	} else if day == "2" && part == "2" {
		result, err = day2.Part2()
	} else if day == "3" && part == "1" {
		result, err = day3.Part1()
	} else if day == "3" && part == "2" {
		result, err = day3.Part2()
	} else if day == "4" && part == "1" {
		result, err = day4.Part1()
	} else {
		err = fmt.Errorf("Unrecognised day/part: %s/%s", day, part)
	}

	if err != nil {
		log.Fatalf("%s\n", err)
	}

	fmt.Printf("%s\n", result)
}
