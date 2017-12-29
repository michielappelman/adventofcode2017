package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/michielappelman/adventofcode2017/generic"
)

type Component struct {
	a, b int
}

func (c Component) HasPort(port int) bool {
	return c.a == port || c.b == port
}
func (c Component) OtherPort(port int) int {
	if c.a == port {
		return c.b
	} else {
		return c.a
	}
}

type Bridge []Component

func (b Bridge) Strength() int {
	var sum int
	for _, c := range b {
		sum += c.a + c.b
	}
	return sum
}

func loadComps(input []string) map[Component]struct{} {
	comps := make(map[Component]struct{})
	for _, line := range input {
		split := strings.Split(line, "/")
		a, _ := strconv.Atoi(split[0])
		b, _ := strconv.Atoi(split[1])
		comps[Component{a, b}] = struct{}{}
	}
	return comps
}

func newMap(c Component, comps map[Component]struct{}) map[Component]struct{} {
	newComps := make(map[Component]struct{})
	for old, _ := range comps {
		if old != c {
			newComps[old] = struct{}{}
		}
	}
	return newComps
}

func getBridges(allComps map[Component]struct{}) []Bridge {
	var bridges []Bridge
	var getNext func(val int, comps map[Component]struct{}, b Bridge)
	getNext = func(val int, comps map[Component]struct{}, b Bridge) {
		for c, _ := range comps {
			if c.HasPort(val) {
				b := b
				b = append(b, c)
				bridges = append(bridges, b)
				newComps := newMap(c, comps)
				getNext(c.OtherPort(val), newComps, b)
			}
		}
	}
	getNext(0, allComps, Bridge{})
	return bridges
}

func StarOne(input []string) int {
	allComps := loadComps(input)
	bridges := getBridges(allComps)
	var strengths []int
	for _, b := range bridges {
		strengths = append(strengths, b.Strength())
	}

	return generic.Max(strengths)
}

func StarTwo(input []string) int {
	allComps := loadComps(input)
	bridges := getBridges(allComps)

	var lengths []int
	bridgeLengths := make(map[int][]Bridge)
	for _, b := range bridges {
		l := len(b)
		lengths = append(lengths, l)
		bridgeLengths[l] = append(bridgeLengths[l], b)
	}
	var strengths []int
	for _, b := range bridgeLengths[generic.Max(lengths)] {
		strengths = append(strengths, b.Strength())
	}
	return generic.Max(strengths)
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
