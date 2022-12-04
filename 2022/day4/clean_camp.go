package day4

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func countOverlap(input string) int {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	conter := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		in := scanner.Text()
		pair := strings.Split(in, ",")

		if opverlap(pair) {
			conter++
		}
	}
	return conter
}

func countContained(input string) int {

	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	conter := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		in := scanner.Text()
		pair := strings.Split(in, ",")
		o := oneContained(pair)
		if o {
			conter++
		}
	}
	return conter
}

func opverlap(pair []string) bool {
	if oneContained(pair) {
		return true
	}
	s1 := toIntPair(strings.Split(pair[0], "-"))
	s2 := toIntPair(strings.Split(pair[1], "-"))

	if s1[1] >= s2[0] && s1[0] <= s2[0] {
		return true
	}
	if s2[1] >= s1[0] && s2[1] <= s1[1] {
		return true
	}
	return false

}

func oneContained(pair []string) bool {
	s1 := toIntPair(strings.Split(pair[0], "-"))
	s2 := toIntPair(strings.Split(pair[1], "-"))
	len1 := s1[1] - s1[0]
	len2 := s2[1] - s2[0]
	if len1 >= len2 {
		return s1[0] <= s2[0] && s1[1] >= s2[1]
	}
	return s2[0] <= s1[0] && s2[1] >= s1[1]

}

func toIntPair(s []string) []int {
	a, err := strconv.Atoi(s[0])
	if err != nil {
		log.Panicln(err)
	}
	b, err := strconv.Atoi(s[1])
	if err != nil {
		log.Panicln(err)
	}
	return []int{a, b}
}
