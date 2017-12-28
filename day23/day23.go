package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	command string
	a, b    string
}

func getValue(x string, reg map[string]int) int {
	n, err := strconv.Atoi(x)
	if err == nil {
		return n
	} else if _, ok := reg[x]; ok {
		return reg[x]
	} else {
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
	var count int
	for i := 0; i < len(instr); {
		in := instr[i]
		va, vb := getValue(in.a, reg), getValue(in.b, reg)
		switch in.command {
		case "set":
			reg[in.a] = vb
		case "sub":
			reg[in.a] -= vb
		case "mul":
			count++
			reg[in.a] *= vb
		case "jnz":
			if va != 0 {
				i += vb
				continue
			}
		}
		i++
	}
	return count
}

func StarTwo(input []string) int {
	reg := make(map[string]int)
	reg["a"] = 1
	instr := loadMoves(input)
	for i := 0; i < 8; {
		in := instr[i]
		va, vb := getValue(in.a, reg), getValue(in.b, reg)
		switch in.command {
		case "set":
			reg[in.a] = vb
		case "sub":
			reg[in.a] -= vb
		case "mul":
			reg[in.a] *= vb
		case "jnz":
			if va != 0 {
				i += vb
				continue
			}
		}
		i++
	}
	for {
		i := big.NewInt(int64(reg["b"]))
		if !i.ProbablyPrime(10) {
			reg["h"]++
		}
		if reg["b"] == reg["c"] {
			break
		}
		reg["b"] -= -17
	}
	return reg["h"]
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
