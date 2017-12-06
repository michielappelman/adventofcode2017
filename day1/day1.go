package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/michielappelman/adventofcode2017/generic"
)

func StarOne(input string) int {
	var doubles []int
	length := len(input)

	input += string(input[0])

	for i := 0; i < length; i++ {
		if input[i] == input[i+1] {
			toInt, _ := strconv.Atoi(string(input[i]))
			doubles = append(doubles, toInt)
		}
	}
	return generic.Sum(doubles)
}

func StarTwo(input string) int {
	var doubles []int
	length := len(input)

	input += string(input[0])

	for i := 0; i < length; i++ {
		if input[i] == input[(i+length/2)%length] {
			toInt, _ := strconv.Atoi(string(input[i]))
			doubles = append(doubles, toInt)
		}
	}
	return generic.Sum(doubles)
}

func Sum(list []int) int {
	var sum int
	for _, num := range list {
		sum += num
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println("1:", StarOne(input))
		fmt.Println("2:", StarTwo(input))
	}
}
