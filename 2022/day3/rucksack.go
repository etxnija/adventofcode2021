package day3

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func SumOfBadges(input string) int {

	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	score := 0
	for scanner.Scan() {

		firstItems := scanner.Text()
		scanner.Scan()
		secountItems := scanner.Text()
		scanner.Scan()
		thirdItems := scanner.Text()
		potential := findPotential(firstItems, secountItems)

		common := findCommon(string(potential), thirdItems)
		p := findPrioritie(common)
		score = score + p
	}

	return score

}

func SumOfpriorities(input string) int {

	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	score := 0
	for scanner.Scan() {
		items := scanner.Text()
		one := items[:len(items)/2]
		two := items[len(items)/2:]

		common := findCommon(one, two)
		p := findPrioritie(common)
		score = score + p
	}

	return score
}

func findPrioritie(common rune) int {
	if common >= 'a' {
		return int(common) - int('a') + 1
	}
	return int(common) - int('A') + 27
}

func findCommon(one, two string) rune {
	for _, v := range one {
		if strings.ContainsRune(two, v) {
			return v
		}
	}
	return ' '
}

func findPotential(one, two string) []rune {
	var p []rune
	for _, v := range one {
		if strings.ContainsRune(two, v) {
			p = append(p, v)
		}
	}
	return p
}
