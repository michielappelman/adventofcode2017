package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/michielappelman/adventofcode2017/generic"
)

func StarOne(input string) int {
	pipes := make(map[int][]int)
	re := regexp.MustCompile(`(\d+) <-> (.*)`)
	matches := re.FindAllStringSubmatch(input, -1)
	for _, m := range matches {
		prog, _ := strconv.Atoi(m[1])
		subprog := generic.StringsToInts(strings.Split(m[2], ", "))
		pipes[prog] = subprog
	}

	var count int
	seen := make(map[int]bool)

	var visit func(progs []int)
	visit = func(progs []int) {
		for _, p := range progs {
			if !seen[p] {
				seen[p] = true
				visit(pipes[p])
				count++
			}
		}
	}

	visit([]int{0})
	return count
}

func StarTwo(input string) int {
	pipes := make(map[int][]int)
	re := regexp.MustCompile(`(\d+) <-> (.*)`)
	matches := re.FindAllStringSubmatch(input, -1)
	for _, m := range matches {
		prog, _ := strconv.Atoi(m[1])
		subprog := generic.StringsToInts(strings.Split(m[2], ", "))
		pipes[prog] = subprog
	}

	var count int
	seen := make(map[int]bool)

	var visit func(progs []int)
	visit = func(progs []int) {
		for _, p := range progs {
			if !seen[p] {
				seen[p] = true
				visit(pipes[p])
			}
		}
	}
	visit([]int{0})
	count++

	for k, _ := range pipes {
		if !seen[k] {
			count++
			seen[k] = true
			visit(pipes[k])
		}
	}

	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var pipes string
	for scanner.Scan() {
		pipes += scanner.Text() + "\n"
	}
	fmt.Println("1:", StarOne(pipes))
	fmt.Println("2:", StarTwo(pipes))
}
