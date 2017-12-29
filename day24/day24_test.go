package main

import (
	"testing"
)

func TestStarOne(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{[]string{"0/2",
			"2/2",
			"2/3",
			"3/4",
			"3/5",
			"0/1",
			"10/1",
			"9/10"}, 31},
	}
	for _, test := range tests {
		got := StarOne(test.input)
		if got != test.want {
			t.Errorf("for %s got %d, want %d", test.input, got, test.want)
		}
	}
}
