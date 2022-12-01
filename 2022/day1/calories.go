package day1

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func findElfWithMostCalories(input string) int {
	cals := calories(input)
	sort.Ints(cals)

	return cals[len(cals)-1]
}

func threeTotalElfWithMostCalories(input string) int {
	cals := calories(input)
	sort.Ints(cals)
	cals = cals[len(cals)-3:]
	sum := 0
	for _, v := range cals {
		sum = sum + v
	}

	return sum
}

func calories(input string) []int {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var cals []int

	var totalC int
	for scanner.Scan() {

		row := strings.TrimSpace(scanner.Text())
		if len(row) < 1 {
			cals = append(cals, totalC)
			totalC = 0
			continue
		}
		c, err := strconv.Atoi(row)
		if err != nil {
			log.Fatal(err)
		}
		totalC = totalC + c
	}
	return cals
}
