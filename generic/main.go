package generic

import "strconv"

func Sum(list []int) int {
	var sum int
	for _, num := range list {
		sum += num
	}
	return sum
}

func StringsToInts(list []string) []int {
	var rowInts []int
	for _, c := range list {
		toInt, err := strconv.Atoi(c)
		if err != nil {
			continue
		}
		rowInts = append(rowInts, toInt)
	}
	return rowInts
}
