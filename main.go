package main

import (
	"fmt"
	"log"
	"os"
	Day1 "robertbrignull/adventofcode2023/days/day1"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 && args[0] == "--" {
		args = args[1:]
	}
	if len(args) < 1 {
		log.Fatal("Usage: ./run <day>\n")
	}

	var result string
	var err error
	result, err = "", nil
	switch args[0] {
	case "1":
		result, err = Day1.Run()
	default:
		err = fmt.Errorf("Unrecognised day: %s", args[0])
	}

	if err != nil {
		log.Fatalf("%s\n", err)
	}

	fmt.Printf("%s\n", result)
}
