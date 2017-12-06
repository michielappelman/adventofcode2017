package main

import (
	"testing"
)

func TestStarOne(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{6, 1},
		{12, 3},
		{13, 4},
		{23, 2},
		{1024, 31},
	}
	for _, test := range tests {
		got := StarOne(test.input)
		if got != test.want {
			t.Errorf("for %d got %d, want %d", test.input, got, test.want)
		}
	}
}
