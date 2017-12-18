package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	command string
	a, b    string
}

func getValue(id int, x string, reg map[string]int) int {
	n, err := strconv.Atoi(x)
	if err == nil {
		return n
	} else if _, ok := reg[x]; ok {
		return reg[x]
	} else {
		if x == "p" {
			reg[x] = id
			return id
		}
		reg[x] = 0
		return 0
	}
}

func loadMoves(input []string) []Instruction {
	var instr []Instruction
	for _, in := range input {
		fields := strings.Fields(in)
		fields = append(fields, "0")
		instr = append(instr, Instruction{fields[0], fields[1], fields[2]})
	}
	return instr
}

func StarOne(input []string) int {
	reg := make(map[string]int)
	instr := loadMoves(input)
	var id, sound int
	for i := 0; i < len(instr); {
		in := instr[i]
		va, vb := getValue(id, in.a, reg), getValue(id, in.b, reg)
		switch in.command {
		case "set":
			reg[in.a] = vb
		case "add":
			reg[in.a] += vb
		case "mul":
			reg[in.a] *= vb
		case "mod":
			reg[in.a] %= vb
		case "snd":
			sound = va
		case "rcv":
			if va != 0 {
				return sound
			}
		case "jgz":
			if va > 0 {
				i += vb
				continue
			}
		}
		i++
	}
	return 0
}

func Prog(id int, instr []Instruction, send, recv, countChan chan int, run *[2]bool) {
	reg := make(map[string]int)
	var count int
	for i := 0; i < len(instr); {
		run[id] = true
		in := instr[i]
		va, vb := getValue(id, in.a, reg), getValue(id, in.b, reg)
		switch in.command {
		case "set":
			reg[in.a] = vb
		case "add":
			reg[in.a] += vb
		case "mul":
			reg[in.a] *= vb
		case "mod":
			reg[in.a] %= vb
		case "snd":
			count++
			send <- va
		case "rcv":
			run[id] = false
			if len(recv) == 0 && run[0] == run[1] {
				if id == 1 {
					countChan <- count
				}
				return
			}
			reg[in.a] = <-recv
			run[id] = true
		case "jgz":
			if va > 0 {
				i += vb
				continue
			}
		}
		i++
	}
}

func StarTwo(input []string) int {
	instr := loadMoves(input)
	c0 := make(chan int, 100)
	c1 := make(chan int, 100)
	cC := make(chan int)
	run := &[2]bool{false, false}

	go Prog(0, instr, c1, c0, cC, run)
	go Prog(1, instr, c0, c1, cC, run)

	return <-cC
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
