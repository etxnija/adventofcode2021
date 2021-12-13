package day4

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type bingoboard [][]string

func caluculateScore(boards []bingoboard, numbers []string) int {
	b, n := playBingo(boards, numbers)
	board := boards[b]
	return calc(board, numbers[n])
}
func calc(b bingoboard, n string) int {
	sum := 0
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b); j++ {
			if n, err := strconv.Atoi(b[i][j]); err == nil {
				sum = sum + n
			}
		}
	}
	nn, _ := strconv.Atoi(n)
	return nn * sum
}

func calculateLast(boards []bingoboard, numbers []string) int {

	b, nn := findLastWinner(boards, numbers)
	return calc(b, numbers[nn])
}

func findLastWinner(boards []bingoboard, numbers []string) (bingoboard, int) {

	var candidateBoard bingoboard
	var candidateNumber int
	loops := 0
	for i := 0; i < len(boards); i++ {
		bn, nn := playBingo(boards, numbers)
		candidateBoard = boards[bn]
		candidateNumber = nn
		boards[bn] = make(bingoboard, 0)
		loops++
	}
	return candidateBoard, candidateNumber
}

func playBingo(boards []bingoboard, numbers []string) (int, int) {
	var currentBoard int
	var currentNumber int
out:
	for i, v := range numbers {
		currentNumber = i
		for bn, b := range boards {
			currentBoard = bn

			for i := 0; i < len(b); i++ {
				xhits := 0
				yhits := 0
				for j := 0; j < len(b); j++ {
					if b[i][j] == v {
						b[i][j] = "x"
					}
					if b[i][j] == "x" {
						xhits++
					}
					if b[j][i] == "x" {
						yhits++
					}
					if xhits == len(b) || yhits == len(b) {
						break out
					}
				}
			}
		}
	}
	return currentBoard, currentNumber
}

func readBoards(f string) ([]bingoboard, []string) {
	bb := make([]bingoboard, 0)
	var input []string

	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var b bingoboard
	var y int

	for scanner.Scan() {
		t := scanner.Text()
		t = strings.Trim(t, " ")
		if len(t) > 15 {
			input = strings.Split(t, ",")

		} else if len(t) < 1 {
			b = make(bingoboard, 5)
			y = 0
			bb = append(bb, b)
		} else {
			// log.Println("board", t, b)
			in := strings.Fields(t)
			b[y] = in
			y++
		}
	}
	return bb, input
}
