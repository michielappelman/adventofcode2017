#!/usr/bin/env bash
# Create a new directory for a new day of the adventofcode

echo "What date are you going to solve today?"
read day

DAYNAME=day$day

if [ -d $DAYNAME ] ; then
    echo "Directory $DAYNAME exists!"
    exit 1
fi

mkdir $DAYNAME
cd $DAYNAME
cat <<EOF > $DAYNAME.go
package main

func StarOne(input string) int {
    return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println("1:", StarOne(input))
		//fmt.Println("2:", StarTwo(input))
	}
}
EOF

cat <<EOF > ${DAYNAME}_test.go
package main

import (
	"testing"
)

func TestStarOne(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
		{"6644789", 10},
		{"578444444785", 25},
	}
	for _, test := range tests {
		got := StarOne(test.input)
		if got != test.want {
			t.Errorf("for %s got %d, want %d", test.input, got, test.want)
		}
	}
}

EOF
