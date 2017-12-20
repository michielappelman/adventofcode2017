package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/michielappelman/adventofcode2017/generic"
)

type Particle struct {
	p, v, a  []int
	distance int
}

func (p *Particle) Tick() {
	p.v[0] += p.a[0]
	p.v[1] += p.a[1]
	p.v[2] += p.a[2]
	p.p[0] += p.v[0]
	p.p[1] += p.v[1]
	p.p[2] += p.v[2]
	p.distance = generic.Abs(p.p[0]) + generic.Abs(p.p[1]) + generic.Abs(p.p[2])
}

func loadParticles(input []string) []*Particle {
	var p []*Particle
	for _, line := range input {
		var comps [][]int
		for i, sub := range strings.Split(line, "<") {
			if i > 0 {
				var comp []int
				for j, n := range strings.Split(sub, ",") {
					if j < 3 {
						num, _ := strconv.Atoi(strings.Trim(n, " >"))
						comp = append(comp, num)
					}
				}
				comps = append(comps, comp)
			}
		}
		p = append(p, &Particle{comps[0], comps[1], comps[2], 0})
	}
	return p
}

func StarOne(input []string) int {
	p := loadParticles(input)
	simulations := 500
	for i := 0; i < simulations; i++ {
		for _, x := range p {
			x.Tick()
		}
	}
	list := make(map[int]int)
	var slist []int
	for i, x := range p {
		slist = append(slist, x.distance)
		list[x.distance] = i
	}
	sort.Ints(slist)
	return list[slist[0]]
}

func StarTwo(input []string) int {
	p := loadParticles(input)
	simulations := 500
	for i := 0; i < simulations; i++ {
		seen := make(map[string][]int)
		for j, x := range p {
			if x != nil {
				x.Tick()
				ts := fmt.Sprintf("%v", x.p)
				seen[ts] = append(seen[ts], j)
			}
		}
		for _, s := range seen {
			if len(s) > 1 {
				for _, l := range s {
					// Delete without preserving order
					p[l] = p[len(p)-1]
					p = p[:len(p)-1]
				}
			}
		}
	}
	return len(p)
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
