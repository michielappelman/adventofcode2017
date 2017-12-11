package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/michielappelman/adventofcode2017/generic"
)

type Position struct {
	x, y, z int
}

func (p *Position) Move(direction string) {
	switch direction {
	case "nw":
		p.x++
		p.y--
	case "n":
		p.x++
		p.z--
	case "ne":
		p.y++
		p.z--
	case "se":
		p.x--
		p.y++
	case "s":
		p.x--
		p.z++
	case "sw":
		p.y--
		p.z++
	}
}

func (p *Position) Distance() int {
	return (generic.Abs(p.x) + generic.Abs(p.y) + generic.Abs(p.z)) / 2
}

func StarOne(input string) int {
	in := strings.Split(input, ",")
	p := Position{0, 0, 0}

	for _, m := range in {
		p.Move(m)
	}
	return p.Distance()
}

func StarTwo(input string) int {
	in := strings.Split(input, ",")
	p := Position{0, 0, 0}
	var furthest int

	for _, m := range in {
		p.Move(m)
		d := p.Distance()
		if d > furthest {
			furthest = d
		}
	}
	return furthest
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println("1:", StarOne(input))
		fmt.Println("2:", StarTwo(input))
	}
}
