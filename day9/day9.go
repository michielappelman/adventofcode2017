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
	noGbg, _ := deleteGarbage(noExcl)
	sumGroups := countGroups(noGbg)
	return sumGroups
}

func StarTwo(input string) int {
	noExcl := deleteExclamations(input)
	_, gl := deleteGarbage(noExcl)
	return gl
}

func countGroups(input string) int {
	var sum, count int

	rd := bufio.NewReader(strings.NewReader(input))
	for {
		r, _, err := rd.ReadRune()
		if err == io.EOF {
			return sum
		}
		switch r {
		case '{':
			count++
			sum += count
		case '}':
			count--
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
		} else {
			result += string(r)
		}
	}
}

func deleteGarbage(input string) (string, int) {
	re := regexp.MustCompile(`<(.*?)>`)
	result := re.ReplaceAllString(input, "")
	var garbage int
	matches := re.FindAllStringSubmatch(input, -1)
	if len(matches) > 0 {
		for _, g := range matches {
			for _, m := range g[1:] {
				garbage += len(m)
			}
		}
	}
	return result, garbage
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println("1:", StarOne(input))
		fmt.Println("2:", StarTwo(input))
	}
}
