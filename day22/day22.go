package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N = iota
	E
	S
	W
)

const (
	Weakened = iota - 1
	Clean
	Infected
	Flagged
)

type Direction uint8
type State int8

type Pos struct {
	x, y int
}

type Carrier struct {
	pos       Pos
	direction Direction
}

func (c *Carrier) Burst(grid map[Pos]State) bool {
	var infected bool
	switch grid[c.pos] {
	case Infected:
		c.direction = (c.direction + 1) % 4
		grid[c.pos] = Clean
	case Clean:
		if c.direction == 0 {
			c.direction = 3
		} else {
			c.direction--
		}
		grid[c.pos] = Infected
		infected = true
	}
	switch c.direction {
	case N:
		c.pos.y++
	case E:
		c.pos.x++
	case S:
		c.pos.y--
	case W:
		c.pos.x--
	}
	return infected
}

func (c *Carrier) EvolvedBurst(grid map[Pos]State) bool {
	var infected bool
	switch grid[c.pos] {
	case Weakened:
		infected = true
		grid[c.pos] = Infected
	case Infected:
		c.direction = (c.direction + 1) % 4
		grid[c.pos] = Flagged
	case Clean:
		if c.direction == 0 {
			c.direction = 3
		} else {
			c.direction--
		}
		grid[c.pos] = Weakened
	case Flagged:
		c.direction = (c.direction + 2) % 4
		grid[c.pos] = Clean
	}
	switch c.direction {
	case N:
		c.pos.y++
	case E:
		c.pos.x++
	case S:
		c.pos.y--
	case W:
		c.pos.x--
	}
	return infected
}

func StarOne(input []string) int {
	grid := make(map[Pos]State)
	var w, h int
	for i, line := range input {
		for j, c := range line {
			if c == '#' {
				grid[Pos{j, -i}] = Infected
			}
			w = j
		}
		h = i
	}
	bursts := 10000
	var infections int
	carrier := &Carrier{Pos{w / 2, -(h / 2)}, N}
	for i := 0; i < bursts; i++ {
		if carrier.Burst(grid) {
			infections++
		}
	}
	return infections
}

func StarTwo(input []string) int {
	grid := make(map[Pos]State)
	var w, h int
	for i, line := range input {
		for j, c := range line {
			if c == '#' {
				grid[Pos{j, -i}] = Infected
			}
			w = j
		}
		h = i
	}
	bursts := 10000000
	var infections int
	carrier := &Carrier{Pos{w / 2, -(h / 2)}, N}
	for i := 0; i < bursts; i++ {
		if carrier.EvolvedBurst(grid) {
			infections++
		}
	}
	return infections
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Println("1:", StarOne(input))
	fmt.Println("2:", StarTwo(input))
}
