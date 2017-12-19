package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	N = iota
	E
	S
	W
)

type direction int

func getNext(y, x int, dir direction, grid [][]string, letters string) (int, int, direction, string) {
	l := grid[y][x]
	grid[y][x] = "!"
	if l == " " {
		return 12345, 12345, dir, letters
	}
	if l == "+" {
		switch {
		case y > 0 && grid[y-1][x] != " " && grid[y-1][x] != "!":
			return y - 1, x, direction(N), letters
		case x < len(grid[y]) && grid[y][x+1] != " " && grid[y][x+1] != "!":
			return y, x + 1, direction(E), letters
		case y < len(grid) && grid[y+1][x] != " " && grid[y+1][x] != "!":
			return y + 1, x, direction(S), letters
		case x > 0 && grid[y][x-1] != " " && grid[y][x-1] != "!":
			return y, x - 1, direction(W), letters
		}
	} else {
		if rune(l[0]) >= 65 && rune(l[0]) < 91 {
			letters += l
		}
		switch dir {
		case N:
			return y - 1, x, dir, letters
		case E:
			return y, x + 1, dir, letters
		case S:
			return y + 1, x, dir, letters
		case W:
			return y, x - 1, dir, letters
		}
	}
	return 0, 0, dir, letters
}

func StarOne(input []string) string {
	var letters string
	start := strings.IndexRune(input[0], '|')
	var grid [][]string
	for i := 1; i < len(input); i++ {
		grid = append(grid, strings.Split(input[i], ""))
	}

	y, x, dir := 0, start, direction(S)
	for {
		y, x, dir, letters = getNext(y, x, dir, grid, letters)
		if y == 12345 {
			return letters
		}
	}
}

func StarTwo(input []string) int {
	var steps int
	var letters string
	start := strings.IndexRune(input[0], '|')
	var grid [][]string
	for i := 1; i < len(input); i++ {
		grid = append(grid, strings.Split(input[i], ""))
	}

	y, x, dir := 0, start, direction(S)
	for {
		steps++
		y, x, dir, letters = getNext(y, x, dir, grid, letters)
		if y == 12345 {
			return steps
		}
	}
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
