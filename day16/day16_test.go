package main

import (
	"testing"
)

func TestStarOne(t *testing.T) {
	tests := []struct {
		dancers int
		input   string
		want    string
	}{
		{5, "s1,x3/4,pe/b", "baedc"},
	}
	for _, test := range tests {
		got := StarOne(test.dancers, test.input)
		if got != test.want {
			t.Errorf("for %s got %s, want %s", test.input, got, test.want)
		}
	}
}
