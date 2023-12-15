package main

import (
	"fmt"
	"log"
	"os"
	"robertbrignull/adventofcode2023/days/day1"
	"robertbrignull/adventofcode2023/days/day10"
	"robertbrignull/adventofcode2023/days/day2"
	"robertbrignull/adventofcode2023/days/day3"
	"robertbrignull/adventofcode2023/days/day4"
	"robertbrignull/adventofcode2023/days/day5"
	"robertbrignull/adventofcode2023/days/day6"
	"robertbrignull/adventofcode2023/days/day8"
	"robertbrignull/adventofcode2023/days/day9"
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
	} else if day == "4" && part == "2" {
		result, err = day4.Part2()
	} else if day == "5" && part == "1" {
		result, err = day5.Part1()
	} else if day == "5" && part == "2" {
		result, err = day5.Part2()
	} else if day == "6" && part == "1" {
		result, err = day6.Part1()
	} else if day == "6" && part == "2" {
		result, err = day6.Part2()
	} else if day == "8" && part == "1" {
		result, err = day8.Part1()
	} else if day == "8" && part == "2" {
		result, err = day8.Part2()
	} else if day == "9" && part == "1" {
		result, err = day9.Part1()
	} else if day == "9" && part == "2" {
		result, err = day9.Part2()
	} else if day == "10" && part == "1" {
		result, err = day10.Part1()
	} else if day == "10" && part == "2" {
		result, err = day10.Part2()
	} else {
		err = fmt.Errorf("Unrecognised day/part: %s/%s", day, part)
	}

	if err != nil {
		log.Fatalf("%s\n", err)
	}

	fmt.Printf("%s\n", result)
}
