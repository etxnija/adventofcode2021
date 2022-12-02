package day2

import (
	"bufio"
	"log"
	"os"
)

func CalculateRPSScore(input string) int {

	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// A, X rock   		1
	// B, Y paper  		2
	// C, Z Scissors  	3
	scoreTable := map[string]int{
		"A X": 1 + 3,
		"A Y": 2 + 6,
		"A Z": 3 + 0,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 1 + 6,
		"C Y": 2 + 0,
		"C Z": 3 + 3,
	}

	scanner := bufio.NewScanner(f)
	score := 0
	for scanner.Scan() {
		score = score + scoreTable[scanner.Text()]
	}

	return score
}

func CalculateRPSSscoreStratergy2(input string) int {

	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// A, X rock   		1
	// B, Y paper  		2
	// C, Z Scissors  	3
	// X lose
	// Y Draw
	// Z win
	scoreTable := map[string]int{
		"A X": 3 + 0, // Z lose
		"A Y": 1 + 3, // X deaw
		"A Z": 2 + 6, // Y win
		"B X": 1 + 0, // X lose
		"B Y": 2 + 3, // Y draw
		"B Z": 3 + 6, // Z win
		"C X": 2 + 0, // Y lose
		"C Y": 3 + 3, // Z draw
		"C Z": 1 + 6, // X win
	}

	scanner := bufio.NewScanner(f)
	score := 0
	for scanner.Scan() {
		score = score + scoreTable[scanner.Text()]
	}

	return score
}
