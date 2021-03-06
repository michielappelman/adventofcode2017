package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const factorA = 16807
const factorB = 48271
const divConst = 2147483647

type Pair struct {
	a, b int
}

func StarOne(initA, initB int) int {
	rounds := 40000000
	pairs := make(chan Pair, 100)

	var count int
	var round int
	done := make(chan bool)
	go func() {
		for {
			if round > rounds {
				done <- true
			}
			select {
			case p := <-pairs:
				if checkPair(p) {
					count++
				}
				round++
			}
		}
	}()

	var pair func(p Pair)
	pair = func(p Pair) {
		pairs <- p
		newA := (p.a * factorA) % divConst
		newB := (p.b * factorB) % divConst
		go pair(Pair{newA, newB})
	}

	go pair(Pair{initA, initB})

	<-done
	return count
}

func StarTwo(initA, initB int) int {
	rounds := 5000000

	var count int
	var round int
	done := make(chan bool)
	chanA := make(chan int)
	chanB := make(chan int)
	go func() {
		for {
			if round > rounds {
				done <- true
			}
			select {
			case a := <-chanA:
				if checkPair(Pair{a, <-chanB}) {
					count++
				}
				round++
			case b := <-chanB:
				if checkPair(Pair{<-chanA, b}) {
					count++
				}
				round++
			}
		}
	}()

	var gen func(x, factor, div int, c chan int)
	gen = func(x, factor, div int, c chan int) {
		new := (x * factor) % divConst
		if new%div == 0 {
			c <- new
		}
		go gen(new, factor, div, c)
	}

	go gen(initA, factorA, 4, chanA)
	go gen(initB, factorB, 8, chanB)

	<-done
	return count
}

func checkPair(p Pair) bool {
	var lowA, lowB string
	binA := fmt.Sprintf("%b", p.a)
	if len(binA) < 16 {
		lowA = binA
	} else {
		lowA = binA[len(binA)-16:]
	}
	binB := fmt.Sprintf("%b", p.b)
	if len(binB) < 16 {
		lowB = binB
	} else {
		lowB = binB[len(binB)-16:]
	}
	if lowA == lowB {
		return true
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var init []int
	for scanner.Scan() {
		input := scanner.Text()
		fields := strings.Fields(input)
		num, _ := strconv.Atoi(fields[len(fields)-1])
		init = append(init, num)
	}
	fmt.Println("1:", StarOne(init[0], init[1]))
	fmt.Println("2:", StarTwo(init[0], init[1]))
}
