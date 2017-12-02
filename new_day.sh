#!/usr/bin/env bash
# Create a new directory for a new day of the adventofcode

echo "What date are you going to solve today?"
read day

if [ $day -le 9 ] ; then
    DAYNAME=day0$day
else
    DAYNAME=day$day
fi

if [ -d $DAYNAME ] ; then
    echo "Directory $DAYNAME exists!"
    exit 1
fi

mkdir $DAYNAME
cd $DAYNAME
cat <<EOF > main.go
package main

import (
	"github.com/michielappelman/adventofcode2016/generic"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Input filename not given...")
	}

	input_lines := generic.LoadLines(args[1])
	instructions := generic.SplitLine(input_lines[0], ", ")

    ...

	fmt.Println("Star 1:", dist)
}
EOF

git add .
