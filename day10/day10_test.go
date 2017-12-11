package main

import (
	"testing"
)

func TestStarOne(t *testing.T) {
	tests := []struct {
		input string
		size  int
		want  int
	}{
		{"3,4,1,5", 5, 12},
	}
	for _, test := range tests {
		got := StarOne(test.input, test.size)
		if got != test.want {
			t.Errorf("for %s got %d, want %d", test.input, got, test.want)
		}
	}
}
