package main

import (
	"testing"
)

func TestStarOne(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{[]string{"..#",
			"#..",
			"..."}, 5587},
	}
	for _, test := range tests {
		got := StarOne(test.input)
		if got != test.want {
			t.Errorf("for %s got %d, want %d", test.input, got, test.want)
		}
	}
}

func TestStarTwo(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{[]string{"..#",
			"#..",
			"..."}, 2511944},
	}
	for _, test := range tests {
		got := StarTwo(test.input)
		if got != test.want {
			t.Errorf("for %s got %d, want %d", test.input, got, test.want)
		}
	}
}
