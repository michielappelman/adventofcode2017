package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/michielappelman/adventofcode2017/generic"
)

func StarOne(spreadsheet [][]string) int {
	var rowResult []int
	for _, row := range spreadsheet {
		if len(row) < 1 {
			continue
		}
		rowResult = append(rowResult, diffMinMax(row))
	}
	return generic.Sum(rowResult)
}

func StarTwo(spreadsheet [][]string) int {
	var rowResult []int
	for _, row := range spreadsheet {
		if len(row) < 1 {
			continue
		}
		rowResult = append(rowResult, findEvenDividers(row))
	}
	return generic.Sum(rowResult)
}

func diffMinMax(row []string) int {
	rowInts := generic.StringsToInts(row)
	sort.Ints(rowInts)
	min := rowInts[0]
	max := rowInts[len(rowInts)-1]
	return max - min
}

func findEvenDividers(row []string) int {
	rowInts := generic.StringsToInts(row)
	sort.Sort(sort.Reverse(sort.IntSlice(rowInts)))
	for i := range rowInts {
		for j := i + 1; j < len(rowInts); j++ {
			r := rowInts[i] % rowInts[j]
			if r == 0 {
				return rowInts[i] / rowInts[j]
			}
		}
	}
	return 0
}

func main() {
	var spreadsheet [][]string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := strings.Fields(scanner.Text())
		spreadsheet = append(spreadsheet, row)
	}
	fmt.Println("1:", StarOne(spreadsheet))
	fmt.Println("2:", StarTwo(spreadsheet))
}
