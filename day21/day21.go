package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func FromLine(s string) Grid {
	return Grid(strings.Replace(s, "/", "\n", -1))
}

type Grid string

func (g Grid) Size() int {
	return len(strings.Split(string(g), "\n")[0])
}

func (g Grid) Split() []Grid {
	var grids []Grid
	gridSize := g.Size()

	split := func(sub int) {
		rgrid := []rune(string(g))
		for j := 0; j < gridSize; j += sub {
			t := (gridSize + 1) * j
			for i := 0; i < gridSize; i += sub {
				var gr []rune
				for k := 0; k < sub; k++ {
					for l := 0; l < sub; l++ {
						gr = append(gr, rgrid[t+i+(k*(gridSize+1))+l])
					}
					gr = append(gr, '\n')
				}
				grids = append(grids, Grid(strings.TrimSpace(string(gr))))

			}
		}
	}

	switch {
	case gridSize == 2, gridSize == 3:
		grids = append(grids, g)
	case gridSize%2 == 0:
		split(2)
	case gridSize%3 == 0:
		split(3)
	default:
		log.Fatalf("grid size %d not divisable by 2 or 3\n", gridSize)
	}

	return grids
}

func (g Grid) ToLine() string {
	repl := strings.Replace(string(g), "\n", "/", -1)
	return strings.TrimRight(repl, "/")
}

func (g Grid) Flip() []Grid {
	var grids []Grid
	grids = append(grids, g)
	split := strings.Split(string(g), "\n")

	hFlip := make([]string, len(split))
	for i, j := 0, len(split)-1; i < len(split); i, j = i+1, j-1 {
		hFlip[i] = split[j]
	}
	grids = append(grids, Grid(strings.Join(hFlip, "\n")))

	vFlip := make([]string, len(split))
	for i, s := range split {
		r := []rune(s)
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		vFlip[i] = string(r)
	}
	grids = append(grids, Grid(strings.Join(vFlip, "\n")))

	return grids
}

func (g Grid) Rotate90() Grid {
	var m [][]rune
	split := strings.Split(string(g), "\n")
	for _, r := range split {
		m = append(m, []rune(r))
	}
	r := make([][]rune, len(m[0]))
	for x, _ := range r {
		r[x] = make([]rune, len(m))
	}
	for y, s := range m {
		for x, e := range s {
			r[x][y] = e
		}
	}
	var str []string
	for _, row := range r {
		for i, j := 0, len(row)-1; i < j; i, j = i+1, j-1 {
			row[i], row[j] = row[j], row[i]
		}
		str = append(str, string(row))
	}
	return Grid(strings.Join(str, "\n"))
}

func (g Grid) RotatedGrid() []Grid {
	var grids []Grid
	grid := g
	for i := 0; i < 4; i++ {
		grids = append(grids, grid)
		grid = grid.Rotate90()
	}
	return grids
}

func (g Grid) Transform() []string {
	var grids []Grid
	for _, rotGrid := range g.RotatedGrid() {
		grids = append(grids, rotGrid.Flip()...)
	}
	var stringGrids []string
	for _, grid := range grids {
		stringGrids = append(stringGrids, grid.ToLine())
	}
	return stringGrids
}

func (g Grid) Enhance(rules map[string]string) Grid {
	grids := g.Split()
	var enhancedGrids []Grid
	for _, grid := range grids {
		transforms := grid.Transform()
		for _, t := range transforms {
			if _, ok := rules[t]; ok {
				enhancedGrids = append(enhancedGrids, FromLine(rules[t]))
				break
			}
		}
	}
	return JoinGrids(enhancedGrids)
}

func (g Grid) Count(char string) int {
	return strings.Count(string(g), char)
}

func JoinGrids(grids []Grid) Grid {
	if len(grids) == 1 {
		return grids[0]
	}
	var grid string
	top := int(math.Sqrt(float64(len(grids))))
	for i := 0; i < len(grids); i += top {
		for k := 0; k < len(grids[0]); k += grids[0].Size() + 1 {
			for j := 0; j < top; j++ {
				for l := 0; l < grids[0].Size(); l++ {
					grid += string(grids[i+j][k+l])
				}
			}
			grid += "\n"
		}
	}
	return Grid(strings.TrimSpace(grid))
}

func loadRules(input []string) map[string]string {
	rules := make(map[string]string)
	for _, s := range input {
		fields := strings.Fields(s)
		rules[fields[0]] = fields[2]
	}
	return rules
}

func StarOne(input []string, rounds int) int {
	rules := loadRules(input)
	grid := Grid(".#.\n..#\n###")
	for r := 0; r < rounds; r++ {
		grid = grid.Enhance(rules)
	}
	return grid.Count("#")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Println("1:", StarOne(input, 5))
	fmt.Println("2:", StarOne(input, 18))
}
