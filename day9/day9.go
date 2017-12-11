package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func StarOne(input string) int {
	noExcl := deleteExclamations(input)
	noGbg := deleteGarbage(noExcl)
	sumGroups := countGroups(noGbg, 0)
	return sumGroups
}

func countGroups(input string, myVal int) int {
	if len(input) == 0 {
		return myVal
	}
	sum := myVal

	rd := bufio.NewReader(strings.NewReader(input))
	var open int
	var new string
	for {
		r, _, err := rd.ReadRune()
		if err == io.EOF {
			return sum
		}
		switch r {
		case '{':
			if open > 0 {
				new += string(r)
			}
			open++
		case '}':
			open--
			if open > 0 {
				new += string(r)
			}
			if open == 0 {
				sum += countGroups(new, myVal+1)
			}
		}
	}
}

func deleteExclamations(input string) string {
	var result string

	rd := bufio.NewReader(strings.NewReader(input))
	for {
		r, _, err := rd.ReadRune()
		if err == io.EOF {
			return result
		}
		if r == '!' {
			rd.Discard(1)
		} else if r == ',' {
			continue
		} else {
			result += string(r)
		}
	}
}

func deleteGarbage(input string) string {
	re := regexp.MustCompile(`<.*?>`)
	return re.ReplaceAllString(input, "")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println("1:", StarOne(input))
		//fmt.Println("2:", StarTwo(input))
	}
}
