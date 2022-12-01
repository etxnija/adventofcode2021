package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func findElfWithMostCalories(input string) int {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var totalC int
	max := totalC
	for scanner.Scan() {
		if totalC > max {
			max = totalC
		}
		row := strings.TrimSpace(scanner.Text())
		if len(row) < 1 {
			totalC = 0
			continue
		}
		c, err := strconv.Atoi(row)
		if err != nil {
			log.Fatal(err)
		}
		totalC = totalC + c

	}
	return max
}
