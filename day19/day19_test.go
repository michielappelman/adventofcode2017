package main

import (
	"testing"
)

func TestStarOne(t *testing.T) {
	tests := []struct {
		input []string
		want  string
	}{
		{[]string{"     |          ",
			"     |  +--+    ",
			"     A  |  C    ",
			" F---|----E|--+ ",
			"     |  |  |  D ",
			"     +B-+  +--+ "}, "ABCDEF"},
	}
	for _, test := range tests {
		got := StarOne(test.input)
		if got != test.want {
			t.Errorf("for %s got %s, want %s", test.input, got, test.want)
		}
	}
}
func TestStarTwo(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{[]string{"     |          ",
			"     |  +--+    ",
			"     A  |  C    ",
			" F---|----E|--+ ",
			"     |  |  |  D ",
			"     +B-+  +--+ "}, 38},
	}
	for _, test := range tests {
		got := StarTwo(test.input)
		if got != test.want {
			t.Errorf("for %s got %d, want %d", test.input, got, test.want)
		}
	}
}
