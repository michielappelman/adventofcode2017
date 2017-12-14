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
}

func (l *Layer) Step() {
	if l.scannerPos-1 < 0 {
		l.back = false
	}
	if l.scannerPos+1 >= l.depth {
		l.back = true
	}

	if l.back {
		l.scannerPos--
	} else {
		l.scannerPos++
	}
}

func StarOne(input string) int {
	fwMap := make(map[int]*Layer)
	var totaldepth int
	lines := bufio.NewScanner(strings.NewReader(input))
	for lines.Scan() {
		line := strings.Split(lines.Text(), ": ")
		l, _ := strconv.Atoi(line[0])
		r, _ := strconv.Atoi(line[1])
		fwMap[l] = &Layer{r, 0, l * r, false}
		totaldepth = l + 1
	}
	fw := make([]*Layer, totaldepth)
	for k, v := range fwMap {
		fw[k] = v
	}

	var step int
	var sev int
	for _ = range fw {
		if fw[step] != nil && fw[step].scannerPos == 0 {
			sev += fw[step].sev
		}
		for _, f := range fw {
			if f != nil {
				f.Step()
			}
		}
		step++
	}
	return sev
}

func StarTwo(input string) int {
	fwMap := make(map[int]*Layer)
	var totaldepth int
	lines := bufio.NewScanner(strings.NewReader(input))
	for lines.Scan() {
		line := strings.Split(lines.Text(), ": ")
		l, _ := strconv.Atoi(line[0])
		r, _ := strconv.Atoi(line[1])
		fwMap[l] = &Layer{r, 0, l * r, false}
		totaldepth = l + 1
	}
	fw := make([]*Layer, totaldepth)
	for k, v := range fwMap {
		fw[k] = v
	}

	done := make(chan int)
	checkHit := func(delay int) {
		var hit bool
		for i, f := range fw {
			if f == nil {
				continue
			}
			hit = (i+delay)%(2*(f.depth-1)) == 0
			if hit {
				break
			}
		}
		if !hit {
			done <- delay
		}
	}

	for delay := 0; ; delay++ {
		select {
		case d := <-done:
			return d
		default:
			go checkHit(delay)
		}
	}
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
