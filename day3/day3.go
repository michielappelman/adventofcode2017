package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/michielappelman/adventofcode2017/generic"
)

func StarOne(input int) int {
	var x, y int
	timesPM := 2
	timesXY := 1
	square := 1

	for i := 0; i < timesPM; i++ {
		for j := 0; j < timesXY; j++ {
			x++
			square++
			if square == input {
				return distance(x, y)
			}
		}
		for j := 0; j < timesXY; j++ {
			y++
			square++
			if square == input {
				return distance(x, y)
			}
		}
		timesXY++
		for j := 0; j < timesXY; j++ {
			x--
			square++
			if square == input {
				return distance(x, y)
			}
		}
		for j := 0; j < timesXY; j++ {
			y--
			square++
			if square == input {
				return distance(x, y)
			}
		}
		timesXY++
		timesPM += 4
	}
	return 0
}

func StarTwo(input int) int {
	grid := make(map[int]map[int]int)
	for i := -100; i < 100; i++ {
		grid[i] = map[int]int{}
	}
	grid[0][0] = 1

	var x, y int
	timesPM := 2
	timesXY := 1

	for i := 0; i < timesPM; i++ {
		for j := 0; j < timesXY; j++ {
			x++
			sum := surroundingSum(x, y, &grid)
			grid[y][x] = sum
			if sum >= input {
				return sum
			}
		}
		for j := 0; j < timesXY; j++ {
			y++
			sum := surroundingSum(x, y, &grid)
			grid[y][x] = sum
			if sum >= input {
				return sum
			}
		}
		timesXY++
		for j := 0; j < timesXY; j++ {
			x--
			sum := surroundingSum(x, y, &grid)
			grid[y][x] = sum
			if sum >= input {
				return sum
			}
		}
		for j := 0; j < timesXY; j++ {
			y--
			sum := surroundingSum(x, y, &grid)
			grid[y][x] = sum
			if sum >= input {
				return sum
			}
		}
		timesXY++
		timesPM += 4
	}
	return 0
}

func surroundingSum(x, y int, grid *map[int]map[int]int) int {
	var list []int
	list = append(list, (*grid)[y-1][x-1]) // LD
	list = append(list, (*grid)[y-1][x+1]) // RD
	list = append(list, (*grid)[y][x-1])   // L
	list = append(list, (*grid)[y-1][x])   // D
	list = append(list, (*grid)[y+1][x])   // U
	list = append(list, (*grid)[y][x+1])   // R
	list = append(list, (*grid)[y+1][x-1]) // LU
	list = append(list, (*grid)[y+1][x+1]) // RU
	return generic.Sum(list)
}

func distance(x, y int) int {
	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}
	return x + y
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("please supply an integer: %s", err)
		}
		fmt.Println("1:", StarOne(input))
		fmt.Println("2:", StarTwo(input))
	}
}
