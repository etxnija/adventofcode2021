package day5

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type point []int

type line struct {
	start point
	end   point
}

func (l line) xdelta() int {
	return l.end[0] - l.start[0]
}
func (l line) ydelta() int {
	return l.end[1] - l.start[1]
}

func (l line) points() []point {
	var p []point
	if l.ydelta() == 0 { // horizontal
		if l.xdelta() < 0 { // to left
			for i := l.start[0]; i >= l.end[0]; i-- {
				p = append(p, point{i, l.start[1]})
			}
		}
		if l.xdelta() > 0 { // to right
			for i := l.start[0]; i <= l.end[0]; i++ {
				p = append(p, point{i, l.start[1]})
			}
		}
	} else if l.xdelta() == 0 { // vertical
		if l.ydelta() < 0 { // down
			for i := l.start[1]; i >= l.end[1]; i-- {
				p = append(p, point{l.start[0], i})
			}
		}
		if l.ydelta() > 0 { // up
			for i := l.start[1]; i <= l.end[1]; i++ {
				p = append(p, point{l.start[0], i})
			}
		}
	} else if l.ydelta() < 0 { // down
		if l.xdelta() < 0 { //left
			for i := 0; i <= -l.xdelta(); i++ {
				p = append(p, point{l.start[0] - i, l.start[1] - i})
			}
		} else { // right
			for i := 0; i <= l.xdelta(); i++ {
				p = append(p, point{l.start[0] + i, l.start[1] - i})
			}
		}
	} else { //Up
		if l.xdelta() < 0 { //left
			for i := 0; i <= -l.xdelta(); i++ {
				p = append(p, point{l.start[0] - i, l.start[1] + i})
			}
		} else { // right
			for i := 0; i <= l.xdelta(); i++ {
				p = append(p, point{l.start[0] + i, l.start[1] + i})
			}
		}
	}
	return p
}

func newLine(points []string) line {
	s := strings.Split(points[0], ",")
	e := strings.Split(points[1], ",")
	sx, _ := strconv.Atoi(strings.TrimSpace(s[0]))
	sy, _ := strconv.Atoi(strings.TrimSpace(s[1]))

	ex, err := strconv.Atoi(strings.TrimSpace(e[0]))
	if err != nil {
		log.Fatalln(err)
	}
	ey, _ := strconv.Atoi(strings.TrimSpace(e[1]))
	return line{
		start: point{sx, sy},
		end:   point{ex, ey},
	}

}

func danagerMap(wents []line) int {
	dangerMap := make([][]int, 1000)
	overlap := 0
	for i := range dangerMap {
		dangerMap[i] = make([]int, 1000)
	}

	for _, v := range wents {
		lp := v.points()
		// log.Println("line", lp)
		for _, p := range lp {
			dangerMap[p[1]][p[0]] = dangerMap[p[1]][p[0]] + 1
			if dangerMap[p[1]][p[0]] == 2 {
				overlap++
			}
		}

	}
	log.Println("overlap", overlap)
	// log.Println("lines", wents, dangerMap)
	return overlap
}

func readWents(f string) []line {
	wents := make([]line, 0)

	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		t := scanner.Text()
		points := strings.Split(t, "->")

		wents = append(wents, newLine(points))
	}
	return wents
}
