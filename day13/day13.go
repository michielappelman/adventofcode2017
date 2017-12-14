package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Layer struct {
	depth      int
	scannerPos int
	sev        int
	back       bool
	full       bool
}

func (l Layer) Step() Layer {
	r := Layer{l.depth, l.scannerPos, l.sev, l.back, l.full}
	if l.scannerPos-1 < 0 {
		r.back = false
	}
	if l.scannerPos+1 >= l.depth {
		r.back = true
	}

	if r.back {
		r.scannerPos--
	} else {
		r.scannerPos++
	}
	return r
}

func StarOne(input string) int {
	fwMap := make(map[int]Layer)
	var totaldepth int
	lines := bufio.NewScanner(strings.NewReader(input))
	for lines.Scan() {
		line := strings.Split(lines.Text(), ": ")
		l, _ := strconv.Atoi(line[0])
		r, _ := strconv.Atoi(line[1])
		fwMap[l] = Layer{r, 0, l * r, false, true}
		totaldepth = l + 1
	}
	fw := make([]Layer, totaldepth)
	for k, v := range fwMap {
		fw[k] = v
	}

	var step int
	var sev int
	for _ = range fw {
		if fw[step].full && fw[step].scannerPos == 0 {
			sev += fw[step].sev
		}
		for i, f := range fw {
			if f.full {
				fw[i] = f.Step()
			}
		}
		step++
	}
	return sev
}

func StarTwo(input string) int {
	fwMap := make(map[int]Layer)
	var totaldepth int
	lines := bufio.NewScanner(strings.NewReader(input))
	for lines.Scan() {
		line := strings.Split(lines.Text(), ": ")
		l, _ := strconv.Atoi(line[0])
		r, _ := strconv.Atoi(line[1])
		fwMap[l] = Layer{r, 0, l * r, false, true}
		totaldepth = l + 1
	}
	fw := make([]Layer, totaldepth)
	for k, v := range fwMap {
		fw[k] = v
	}

	var delay int
	for hit := true; hit; delay++ {
		var new []Layer
		for i, f := range fw {
			new = append(new, f)
			if f.full {
				for d := 0; d < delay; d++ {
					new[i] = new[i].Step()
				}
			}
		}
		hit = false
		for step := 0; step < len(new); step++ {
			if new[step].full && new[step].scannerPos == 0 {
				hit = true
				break
			}
			for i, f := range new {
				if f.full {
					new[i] = f.Step()
				}
			}
		}
	}
	return delay - 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var fws string
	for scanner.Scan() {
		fws += scanner.Text() + "\n"
	}
	fmt.Println("1:", StarOne(fws))
	fmt.Println("2:", StarTwo(fws))
}
