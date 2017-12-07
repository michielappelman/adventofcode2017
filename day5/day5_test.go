package main

import (
	"testing"
)

func TestStarOne(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{[]string{"0", "3", "0", "1", "-3"}, 5},
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
		{[]string{"0", "3", "0", "1", "-3"}, 10},
	}
	for _, test := range tests {
		got := StarTwo(test.input)
		if got != test.want {
			t.Errorf("for %v got %d, want %d", test.input, got, test.want)
		}
	}
}
