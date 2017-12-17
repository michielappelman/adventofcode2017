package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func StarOne(input int) int {
	rounds := 2017
	i := 0
	s := []int{0}
	for r := 1; r <= rounds; r++ {
		i = ((i + input) % r) + 1
		s = append(s, 0)
		copy(s[i+1:], s[i:])
		s[i] = r
	}
	return s[i+1]
}

func StarTwo(input int) int {
	rounds := 50000000
	i := 0
	afterZero := 0
	for r := 1; r <= rounds; r++ {
		i = ((i + input) % r) + 1
		if i == 1 {
			afterZero = r
		}
	}
	return afterZero
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		n, err := strconv.Atoi(input)
		if err != nil {
			log.Fatalln("please supply integer, not", n)
		}
		fmt.Println("1:", StarOne(n))
		fmt.Println("2:", StarTwo(n))
	}
}
