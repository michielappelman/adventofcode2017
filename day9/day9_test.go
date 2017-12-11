package main

import (
	"testing"
)

func TestStarOne(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"{}", 1},
		{"{}{}", 2},
		{"{{{}}}", 6},
		{"{{},{}}", 5},
		{"{{},!{{}}", 5},
		{"{{{},{},{{}}}}", 16},
		{"{<a>,<a>,<a>,<a>}", 1},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
	}
	for _, test := range tests {
		got := StarOne(test.input)
		if got != test.want {
			t.Errorf("for %s got %d, want %d", test.input, got, test.want)
		}
	}
}
