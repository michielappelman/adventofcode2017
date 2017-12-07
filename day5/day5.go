package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/michielappelman/adventofcode2017/generic"
)

func StarOne(input []string) int {
	l := generic.StringsToInts(input)
	var steps int
	for i := 0; i < len(l); {
		curr := i
		i = i + l[i]
		l[curr]++
		steps++
	}
	return steps
}

func StarTwo(input []string) int {
	l := generic.StringsToInts(input)
	var steps int
	for i := 0; i < len(l); {
		curr := i
		i = i + l[i]
		if l[curr] >= 3 {
			l[curr]--
		} else {
			l[curr]++
		}
		steps++
	}
	return steps
}

func main() {
	var instructions []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	fmt.Println("1:", StarOne(instructions))
	fmt.Println("2:", StarTwo(instructions))
}
