package generic

func Sum(list []int) int {
	var sum int
	for _, num := range list {
		sum += num
	}
	return sum
}
