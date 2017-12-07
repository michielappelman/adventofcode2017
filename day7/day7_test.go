package main

import (
	"testing"
)

func TestStarOne(t *testing.T) {
	tests := []struct {
		input []string
		want  string
	}{
		{[]string{
			"pbga (66)",
			"xhth (57)",
			"ebii (61)",
			"havc (66)",
			"ktlj (57)",
			"fwft (72) -> ktlj, cntj, xhth",
			"qoyq (66)",
			"padx (45) -> pbga, havc, qoyq",
			"tknk (41) -> ugml, padx, fwft",
			"jptl (61)",
			"ugml (68) -> gyxo, ebii, jptl",
			"gyxo (61)",
			"cntj (57)",
		}, "tknk"},
	}
	for _, test := range tests {
		got := StarOne(test.input)
		if got != test.want {
			t.Errorf("for %v got %s, want %s", test.input, got, test.want)
		}
	}
}

//func TestStarTwo(t *testing.T) {
//	tests := []struct {
//		input string
//		want  int
//	}{
//		{"1212", 6},
//		{"1221", 0},
//		{"123425", 4},
//		{"123123", 12},
//		{"12131415", 4},
//	}
//	for _, test := range tests {
//		got := StarTwo(test.input)
//		if got != test.want {
//			t.Errorf("for %s got %d, want %d", test.input, got, test.want)
//		}
//	}
//}
