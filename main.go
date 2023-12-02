package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 && args[0] == "--" {
		args = args[1:]
	}
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: ./run <day>\n")
		os.Exit(1)
	}
}
