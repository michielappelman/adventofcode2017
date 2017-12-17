package main

import (
	"testing"
)

func TestStarOne(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{3, 638},
	}
	for _, test := range tests {
		got := StarOne(test.input)
		if got != test.want {
			t.Errorf("for %d got %d, want %d", test.input, got, test.want)
		}
	}
}
