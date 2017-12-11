package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/michielappelman/adventofcode2017/generic"
)

type Banks []int

func (b Banks) Realloc() Banks {
	highest := generic.IndexOfMax(b)

	dist := b[highest]
	b[highest] = 0
	for i := (highest + 1) % len(b); dist > 0; {
		b[i]++
		i = (i + 1) % len(b)
		dist--
	}
	return b
}

func StarOne(input string) int {
	var b Banks
	fields := strings.Fields(input)
	banks := generic.StringsToInts(fields)
	b = append(b, banks...)

	var cycles int
	seen := make(map[string]bool)
	for {
		cycles++
		c := b.Realloc()
		s := fmt.Sprintf("%v", c)
		if seen[s] {
			return cycles
		}
		seen[s] = true
	}
}

func StarTwo(input string) int {
	var b Banks
	fields := strings.Fields(input)
	banks := generic.StringsToInts(fields)
	b = append(b, banks...)

	var cycles int
	seen := make(map[string]bool)
	var repeat string
	for {
		cycles++
		c := b.Realloc()
		s := fmt.Sprintf("%v", c)
		if seen[s] {
			repeat = s
			break
		}
		seen[s] = true
	}
	cycles = 0
	for {
		cycles++
		c := b.Realloc()
		s := fmt.Sprintf("%v", c)
		if s == repeat {
			return cycles
		}
		seen[s] = true
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println("1:", StarOne(input))
		fmt.Println("2:", StarTwo(input))
	}
}
