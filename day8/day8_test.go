package main

import (
	"testing"
)

func TestStarOne(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{[]string{
			"b inc 5 if a > 1",
			"a inc 1 if b < 5",
			"c dec -10 if a >= 1",
			"c inc -20 if c == 10",
		}, 1},
	}
	for _, test := range tests {
		got := StarOne(test.input)
		if got != test.want {
			t.Errorf("for %v got %d, want %d", test.input, got, test.want)
		}
	}
}
func TestStarTwo(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{[]string{
			"b inc 5 if a > 1",
			"a inc 1 if b < 5",
			"c dec -10 if a >= 1",
			"c inc -20 if c == 10",
		}, 10},
	}
	for _, test := range tests {
		got := StarTwo(test.input)
		if got != test.want {
			t.Errorf("for %v got %d, want %d", test.input, got, test.want)
		}
	}
}
