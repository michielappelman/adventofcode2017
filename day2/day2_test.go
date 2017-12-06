package main

import (
	"strings"
	"testing"
)

func TestStarOne(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{`5 1 9 5
7 5 3
2 4 6 8
`, 18},
	}
	for _, test := range tests {
		var spreadsheet [][]string
		for _, row := range strings.Split(test.input, "\n") {
			row := strings.Fields(row)
			spreadsheet = append(spreadsheet, row)
		}
		got := StarOne(spreadsheet)
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
		{`5 9 2 8
9 4 7 3
3 8 6 5
`, 9},
	}
	for _, test := range tests {
		var spreadsheet [][]string
		for _, row := range strings.Split(test.input, "\n") {
			row := strings.Fields(row)
			spreadsheet = append(spreadsheet, row)
		}
		got := StarTwo(spreadsheet)
		if got != test.want {
			t.Errorf("for %s got %d, want %d", test.input, got, test.want)
		}
	}
}

func TestFindEvenDividers(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"5 9 2 8", 4},
		{"9 4 7 3", 3},
		{"3 8 6 5", 2},
	}
	for _, test := range tests {
		got := findEvenDividers(strings.Split(test.input, " "))
		if got != test.want {
			t.Errorf("for %s got %d, want %d", test.input, got, test.want)
		}
	}
}

func TestDiffMinMax(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"5 1 9 5", 8},
		{"7 5 3", 4},
		{"2 4 6 8", 6},
	}
	for _, test := range tests {
		got := diffMinMax(strings.Split(test.input, " "))
		if got != test.want {
			t.Errorf("for %s got %d, want %d", test.input, got, test.want)
		}
	}
}
