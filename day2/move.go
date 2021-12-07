package day2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	foward = iota
	down
	up
)

func move(start []int, movment []movment) []int {
	location := start
	for _, v := range movment {
		location = v.move(location)
	}

	return location
}

func moveWithAim(start []int, movment []movment) []int {
	location := start
	for _, v := range movment {
		location = v.moveWithAim(location)
		// log.Println(location)
	}

	return location
}

type movment struct {
	direction int
	length    int
}

func (m *movment) move(location []int) []int {
	switch m.direction {
	case foward:
		return []int{location[0] + m.length, location[1]}
	case up:
		return []int{location[0], location[1] - m.length}
	case down:
		return []int{location[0], location[1] + m.length}
	default:
		return location
	}
}

func (m *movment) moveWithAim(location []int) []int {
	switch m.direction {
	case foward:
		return []int{location[0] + m.length, location[1] + location[2]*m.length, location[2]}
	case up:
		return []int{location[0], location[1], location[2] - m.length}
	case down:
		return []int{location[0], location[1], location[2] + m.length}
	default:
		return location
	}
}

func NewMovement(direction string, lenth int) movment {
	switch direction {
	case "forward":
		return movment{foward, lenth}
	case "up":
		return movment{up, lenth}
	case "down":
		return movment{down, lenth}
	default:
		return movment{}
	}
}

func readMovements() []movment {
	movements := make([]movment, 100)

	file, err := os.Open("directions.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		t := scanner.Text()
		m := strings.Split(t, " ")
		li, _ := strconv.Atoi(m[1])
		movements = append(movements, NewMovement(m[0], li))
	}
	return movements
}
