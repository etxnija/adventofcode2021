package day5

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func rearageStacks(s, m string) string {
	moves := loadMoves(m)
	stacks := loadStacks(s)
	for _, mo := range moves {
		if mo[1] == 0 {
			log.Println("stacks")
		}
		from := stacks[mo[1]-1]
		to := stacks[mo[2]-1]
		var c rune
		for i := 0; i < mo[0]; i++ {
			c, from = pop(from)
			to = push(c, to)
		}
		stacks[mo[1]-1] = from
		stacks[mo[2]-1] = to
	}
	return buildString(stacks)
}

func buildString(stacks [][]rune) string {
	str := make([]rune, len(stacks))
	for i, s := range stacks {
		str[i] = s[len(s)-1]
	}
	return string(str)
}

func push(c rune, to []rune) []rune {
	return append(to, c)
}

func pop(from []rune) (rune, []rune) {
	r := from[len(from)-1]
	return r, from[:len(from)-1]
}

func loadMoves(m string) [][]int {
	f, err := os.Open(m)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var moves [][]int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		movesString := scanner.Text()
		movesString = strings.ReplaceAll(movesString, "move ", "")
		movesString = strings.ReplaceAll(movesString, "from ", "")
		movesString = strings.ReplaceAll(movesString, "to ", "")
		movesString = strings.ReplaceAll(movesString, " ", ",")
		movesString = strings.TrimSpace(movesString)

		marr := strings.Split(movesString, ",")
		moves = append(moves, []int{strToInt(marr[0]), strToInt(marr[1]), strToInt(marr[2])})

	}
	return moves
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Panicln(err)
	}
	return i
}

func loadStacks(s string) [][]rune {
	f, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var stacks [][]rune
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		stacks = append(stacks, []rune(line))
	}
	return stacks
}
