package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dancers []rune

func (d Dancers) Add(n int) Dancers {
	for i := 0; i < n; i++ {
		d = append(d, rune(97+i))
	}
	return d
}

func (d Dancers) Exchange(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d Dancers) Spin(i int) Dancers {
	l := len(d)
	p := d[l-i:]
	p = append(p, d[:l-i]...)
	return p
}

func (d Dancers) Partner(a, b rune) {
	var ia, ib int
	for i, c := range d {
		if c == a {
			ia = i
		}
		if c == b {
			ib = i
		}
	}
	d[ia], d[ib] = d[ib], d[ia]
}

type Move struct {
	move rune
	a, b rune
	x, y int
}

func loadMoves(input string) []Move {
	ins := strings.Split(input, ",")
	moves := make([]Move, len(ins))
	for _, j := range ins {
		m := Move{}
		x, j := j[0], j[1:]
		m.move = rune(x)
		switch x {
		case 'x':
			s := strings.Split(j, "/")
			m.x, _ = strconv.Atoi(s[0])
			m.y, _ = strconv.Atoi(s[1])
		case 'p':
			s := strings.Split(j, "/")
			m.a = []rune(s[0])[0]
			m.b = []rune(s[1])[0]
		case 's':
			m.x, _ = strconv.Atoi(j)
		}
		moves = append(moves, m)
	}
	return moves
}

func goOneRound(d Dancers, moves []Move) Dancers {
	for _, m := range moves {
		switch m.move {
		case 'x':
			d.Exchange(m.x, m.y)
		case 'p':
			d.Partner(m.a, m.b)
		case 's':
			d = d.Spin(m.x)
		}
	}
	return d
}

func StarOne(n int, moves string) string {
	m := loadMoves(moves)

	var d Dancers
	d = d.Add(n)

	d = goOneRound(d, m)
	return string(d)
}

func StarTwo(n int, moves string) string {
	rounds := 1000000000

	m := loadMoves(moves)

	var d Dancers
	d = d.Add(n)

	seen := make(map[string]bool)
	var r int
	for ; r < rounds; r++ {
		d = goOneRound(d, m)
		if seen[string(d)] {
			break
		}
		seen[string(d)] = true
	}
	for t := ((rounds / r) * r) + 1; t < rounds; t++ {
		d = goOneRound(d, m)
	}
	return string(d)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println("1:", StarOne(16, input))
		fmt.Println("2:", StarTwo(16, input))
	}
}
