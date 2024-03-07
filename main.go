package main

import (
	"fmt"
	"os"
)

func panic(msg string, err error) {
	if err == nil {
		fmt.Fprintf(os.Stderr, "ERR: %s\n", msg)
	} else {
		fmt.Fprintf(os.Stderr, "ERR: %s: %v\n", msg, err)
	}

	os.Exit(1)
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		panic("No filepath provided", nil)
	}

	parsedMap := parse(args[0])

	plan(parsedMap)
	printSteps(parsedMap)
}
