package main

import (
	"testing"
)

func TestStarOne(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{[]string{"p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>",
			"p=< 4,0,0>, v=< 0,0,0>, a=<-2,0,0>"}, 0},
	}
	for _, test := range tests {
		got := StarOne(test.input)
		if got != test.want {
			t.Errorf("for %s got %d, want %d", test.input, got, test.want)
		}
	}
}
