package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/michielappelman/adventofcode2017/generic"
)

type String []int

func (s String) Reverse(p, l int) {
	if l > len(s) || l == 1 {
		return
	}
	if p+l > len(s) {
		new := String{}
		new = append(new, s[p:]...)
		new = append(new, s[:(p+l)%len(s)]...)
		new.Reverse(0, len(new))
		for k, _ := range new {
			s[(p+k)%len(s)] = new[k]
		}
		return
	}
	for i, j := p, p+l-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func StarOne(input string) int {
	const gridSize = 128
	var sum int
	for i := 0; i < gridSize; i++ {
		val := fmt.Sprintf("%s-%d", input, i)
		hash := BinKnotHash(val)
		split := strings.Split(hash, "")
		isplit := generic.StringsToInts(split)
		sum += generic.Sum(isplit)
	}
	return sum
}

func StarTwo(input string) int {
	const gridSize = 128

	grid := [][]int{}

	for i := 0; i < gridSize; i++ {
		val := fmt.Sprintf("%s-%d", input, i)
		hash := BinKnotHash(val)
		split := strings.Split(hash, "")
		isplit := generic.StringsToInts(split)
		grid = append(grid, isplit)
	}

	var visit func(x, y int)
	visit = func(x, y int) {
		grid[x][y] = 0
		if x > 0 && grid[x-1][y] == 1 {
			visit(x-1, y)
		}
		if y > 0 && grid[x][y-1] == 1 {
			visit(x, y-1)
		}
		if x < gridSize-1 && grid[x+1][y] == 1 {
			visit(x+1, y)
		}
		if y < gridSize-1 && grid[x][y+1] == 1 {
			visit(x, y+1)
		}
	}

	var regions int
	for r, row := range grid {
		for c, _ := range row {
			if grid[r][c] == 1 {
				visit(r, c)
				regions++
			}
		}
	}
	return regions
}

func BinKnotHash(input string) string {
	size := 256
	var dl []int
	for _, c := range input {
		dl = append(dl, int(c))
	}
	dl = append(dl, 17, 31, 73, 47, 23)

	s := String{}
	for i := 0; i < size; i++ {
		s = append(s, i)
	}

	var p, skip int
	for i := 0; i < 64; i++ {
		for _, l := range dl {
			s.Reverse(p, l)
			p = (p + l + skip) % size
			skip++
		}
	}
	var hex string
	for i := 0; i < size; i += 16 {
		var xor int
		for j := 0; j < 16; j++ {
			xor ^= s[i+j]
		}
		hex += fmt.Sprintf("%02x", xor)
	}
	var b string
	for _, c := range hex {
		ui, err := strconv.ParseUint(string(c), 16, 64)
		if err != nil {
			log.Fatalf("error parsing %v", c)
		}
		b += fmt.Sprintf("%04b", ui)
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println("1:", StarOne(input))
		fmt.Println("2:", StarTwo(input))
	}
}
