package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	writeValue bool
	move       int
	newState   string
}

func (o Operation) Run(pos int) (int, int, string) {
	var w int
	if o.writeValue {
		w = 1
	} else {
		w = 0
	}
	pos += o.move
	return w, pos, o.newState
}

func StarOne(input string) int {
	lines := strings.Split(input, "\n")
	state := strings.TrimRight(strings.Fields(lines[0])[3], ".")
	steps, _ := strconv.Atoi(strings.Fields(lines[1])[5])

	blueprint := make(map[string][]Operation)
	for i := 3; i < len(lines); i += 10 {
		state := strings.TrimRight(strings.Fields(lines[i])[2], ":")
		var operations []Operation
		for j := 2; j < 8; j += 4 {
			var val bool
			if strings.TrimRight(strings.Fields(lines[i+j])[4], ".") == "1" {
				val = true
			}
			var move int
			if strings.TrimRight(strings.Fields(lines[i+j+1])[6], ".") == "right" {
				move = 1
			} else {
				move = -1
			}
			newState := strings.TrimRight(strings.Fields(lines[i+j+2])[4], ".")
			operations = append(operations, Operation{val, move, newState})
		}
		blueprint[state] = operations
	}

	tape := make(map[int]int)
	var pos int

	for i := 0; i < steps; i++ {
		op := blueprint[state][tape[pos]]
		tape[pos], pos, state = op.Run(pos)
	}
	var sum int
	for _, v := range tape {
		sum += v
	}
	return sum
}

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln("error reading from stdin")
	}
	input := string(data)
	fmt.Println("1:", StarOne(input))
}
