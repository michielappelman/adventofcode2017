package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func StarOne(input string, size int) int {
	lengths := strings.Split(input, ",")

	s := String{}
	for i := 0; i < size; i++ {
		s = append(s, i)
	}

	var p, skip int
	for _, l := range lengths {
		i, _ := strconv.Atoi(l)
		s.Reverse(p, i)
		p = (p + i + skip) % size
		skip++
	}
	return s[0] * s[1]
}

func StarTwo(input string, size int) string {
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
		hex += fmt.Sprintf("%02x", s[i+0]^s[i+1]^s[i+2]^s[i+3]^s[i+4]^s[i+5]^s[i+6]^s[i+7]^s[i+8]^s[i+9]^s[i+10]^s[i+11]^s[i+12]^s[i+13]^s[i+14]^s[i+15])
	}
	return hex
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println("1:", StarOne(input, 256))
		fmt.Println("2:", StarTwo(input, 256))
	}
}
