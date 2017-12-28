package main

import (
	"testing"
)

//func TestListTransformations(t *testing.T) {
//tests := []struct {
//input string
//size  int
//want  map[string]bool
//}{
//{"#.\n..", 2, map[string]bool{
//"#./..": true,
//".#/..": true,
//"../.#": true,
//"../#.": true}},
//{"##\n..", 2, map[string]bool{
//"##/..": true,
//"../##": true,
//".#/.#": true,
//"#./#.": true}},
//{"#.\n.#", 2, map[string]bool{
//".#/#.": true,
//"#./.#": true}},
//{"##\n#.", 2, map[string]bool{
//"#./##": true,
//"##/.#": true,
//".#/##": true,
//"##/#.": true}},
//{"#..\n...\n...", 3, map[string]bool{
//"#../.../...": true,
//"..#/.../...": true,
//".../.../..#": true,
//".../.../#..": true}},
//{"#.#\n...\n...", 3, map[string]bool{
//"#.#/.../...": true,
//"..#/.../..#": true,
//".../.../#.#": true,
//"#../.../#..": true}},
//{"#.#\n#..\n...", 3, map[string]bool{
//"#.#/#../...": true,
//".##/.../..#": true,
//".../..#/#.#": true,
//"#../.../##.": true}},
//}
//for _, test := range tests {
//got := listTransformations(test.input, test.size)
//if len(got) != len(test.want) {
//t.Errorf("for %s len %d not %d", test.input, len(got), len(test.want))
//}
//for _, result := range got {
//if !test.want[result] {
//t.Errorf("for %s result %s not in %v", test.input, result, test.want)
//}
//}
//}
//}

func TestStarOne(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{[]string{"../.# => ##./#../...",
			".#./..#/### => #..#/..../..../#..#"}, 12},
	}
	for _, test := range tests {
		got := StarOne(test.input)
		if got != test.want {
			t.Errorf("for %s got %d, want %d", test.input, got, test.want)
		}
	}
}
