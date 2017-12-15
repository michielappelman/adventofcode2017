package main

import (
	"testing"
)

func TestStarOne(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{65, 8921}, 588},
	}
	for _, test := range tests {
		got := StarOne(test.input[0], test.input[1])
		if got != test.want {
			t.Errorf("for %v got %d, want %d", test.input, got, test.want)
		}
	}
}
func TestStarTwo(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{65, 8921}, 309},
	}
	for _, test := range tests {
		got := StarTwo(test.input[0], test.input[1])
		if got != test.want {
			t.Errorf("for %v got %d, want %d", test.input, got, test.want)
		}
	}
}
