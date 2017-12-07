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

//func StarTwo(input string) int {
//    ...
//}

func main() {
	var instructions []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	fmt.Println("1:", StarOne(instructions))
	//fmt.Println("2:", StarTwo(instructions))
}
