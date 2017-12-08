package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Program struct {
	name        string
	weight      int
	subPrograms []string
}

func StarOne(input []string) string {
	programs := make(map[string]*Program)
	re := regexp.MustCompile(`(\w+) \((\d+)\)( -> (.+))?`)
	for _, line := range input {
		result := re.FindStringSubmatch(line)
		name := result[1]
		weight, _ := strconv.Atoi(result[2])
		var subprogs []string
		if len(result) > 3 {
			subprogs = strings.Split(result[4], ", ")
		}
		programs[name] = &Program{name, weight, subprogs}
	}
	return findFirst(programs)
}

func StarTwo(input []string) int {
	programs := make(map[string]*Program)
	re := regexp.MustCompile(`(\w+) \((\d+)\)( -> (.+))?`)
	for _, line := range input {
		result := re.FindStringSubmatch(line)
		name := result[1]
		weight, _ := strconv.Atoi(result[2])
		var subprogs []string
		if len(result) > 3 {
			subprogs = strings.Split(result[4], ", ")
		}
		programs[name] = &Program{name, weight, subprogs}
	}
	first := findFirst(programs)
	_, newWeight := sumOfSubProgs(programs[first], programs, 0)
	return newWeight
}

func sumOfSubProgs(program *Program, programs map[string]*Program, firstW int) (map[string]int, int) {
	if len(program.subPrograms) <= 1 || len(program.subPrograms[0]) == 0 {
		return map[string]int{program.name: program.weight}, firstW
	}
	weighted := make(map[string]int)
	var newWeight int
	for _, subprog := range program.subPrograms {
		var w map[string]int
		var nw int
		w, nw = sumOfSubProgs(programs[subprog], programs, firstW)
		if nw > 0 {
			return nil, nw
		}
		for k, v := range w {
			weighted[k] = v
		}
	}
	uniqueWeights := make(map[int]int)
	var weight int
	for _, w := range weighted {
		uniqueWeights[w]++
		weight += w
	}
	var weirdWeight, correctWeight int
	if len(uniqueWeights) > 1 && newWeight == 0 {
		for k, v := range uniqueWeights {
			if v == 1 {
				weirdWeight = k
			} else {
				correctWeight = k
			}
		}
		var weirdOwner string
		for k, v := range weighted {
			if v == weirdWeight {
				weirdOwner = k
			}
		}
		diff := weirdWeight - correctWeight
		firstW = programs[weirdOwner].weight - diff
	}
	return map[string]int{program.name: weight + program.weight}, firstW
}

func findFirst(programs map[string]*Program) string {
	hasPointer := make(map[string]bool, len(programs))
	for _, prog := range programs {
		hasPointer[prog.name] = false
	}
	for _, prog := range programs {
		for _, subprog := range prog.subPrograms {
			delete(hasPointer, subprog)
		}
	}
	for k, _ := range hasPointer {
		return k
	}
	return ""
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
