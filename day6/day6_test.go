package main

import (
	"testing"
)

func TestStarOne(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"0   2   7   0", 5},
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
		input string
		want  int
	}{
		{"0   2   7   0", 4},
	}
	for _, test := range tests {
		got := StarTwo(test.input)
		if got != test.want {
			t.Errorf("for %s got %d, want %d", test.input, got, test.want)
		}
	}
}
