package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Op int
type Comp int

type Condition struct {
	register string
	comp     Comp
	value    int
}

type Instruction struct {
	register  string
	op        Op
	amount    int
	condition *Condition
}

const (
	dec = -1
	inc = 1
)

const (
	EQ = iota
	NE
	GT
	GE
	LT
	LE
)

func (i Instruction) Execute(registers map[string]int) {
	if i.condition.Test(registers) {
		registers[i.register] += i.amount * int(i.op)
	}
}

func (c Condition) Test(registers map[string]int) bool {
	switch c.comp {
	case EQ:
		return registers[c.register] == c.value
	case NE:
		return registers[c.register] != c.value
	case GT:
		return registers[c.register] > c.value
	case GE:
		return registers[c.register] >= c.value
	case LT:
		return registers[c.register] < c.value
	case LE:
		return registers[c.register] <= c.value
	default:
		return false
	}
}

func StarOne(input []string) int {
	var instructions []*Instruction
	registers := make(map[string]int)

	var compConsts = map[string]Comp{
		"==": 0,
		"!=": 1,
		">":  2,
		">=": 3,
		"<":  4,
		"<=": 5,
	}

	reI := regexp.MustCompile(`(\w+) (\w+) ([0-9-]+) if (\w+) (.+) ([0-9-]+)`)
	for _, line := range input {
		result := reI.FindStringSubmatch(line)
		register := result[1]
		var op Op
		if result[2] == "inc" {
			op = inc
		} else {
			op = dec
		}
		amount, err := strconv.Atoi(result[3])
		if err != nil {
			log.Fatalf("%s not an integer", result[3])
		}

		val, err := strconv.Atoi(result[6])
		if err != nil {
			log.Fatalf("%s not an integer", result[6])
		}
		condition := &Condition{result[4], compConsts[result[5]], val}
		instructions = append(instructions, &Instruction{register, op, amount, condition})

	}
	for _, i := range instructions {
		i.Execute(registers)
	}
	var highest int
	for _, v := range registers {
		if v > highest {
			highest = v
		}
	}
	return highest
}

func StarTwo(input []string) int {
	var instructions []*Instruction
	registers := make(map[string]int)

	var compConsts = map[string]Comp{
		"==": 0,
		"!=": 1,
		">":  2,
		">=": 3,
		"<":  4,
		"<=": 5,
	}

	reI := regexp.MustCompile(`(\w+) (\w+) ([0-9-]+) if (\w+) (.+) ([0-9-]+)`)
	for _, line := range input {
		result := reI.FindStringSubmatch(line)
		register := result[1]
		var op Op
		if result[2] == "inc" {
			op = inc
		} else {
			op = dec
		}
		amount, err := strconv.Atoi(result[3])
		if err != nil {
			log.Fatalf("%s not an integer", result[3])
		}

		val, err := strconv.Atoi(result[6])
		if err != nil {
			log.Fatalf("%s not an integer", result[6])
		}
		condition := &Condition{result[4], compConsts[result[5]], val}
		instructions = append(instructions, &Instruction{register, op, amount, condition})

	}
	var highest int
	for _, i := range instructions {
		i.Execute(registers)
		for _, v := range registers {
			if v > highest {
				highest = v
			}
		}
	}
	return highest
}

func main() {
	var instructions []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	fmt.Println("1:", StarOne(instructions))
	fmt.Println("2:", StarTwo(instructions))
}
