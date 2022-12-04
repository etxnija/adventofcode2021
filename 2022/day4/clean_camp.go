package day4

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func oneContained(pair []string) bool {
	section1 := strings.Split(pair[0], "-")
	section2 := strings.Split(pair[1], "-")
	intSection1 := toIntPair(section1)
	intSection2 := toIntPair(section2)
	len1 := intSection1[1] - intSection1[0]
	len2 := intSection2[1] - intSection2[0]
	if len1 >= len2 {
		return intSection1[0] <= intSection2[0] && intSection1[1] >= intSection2[1]
	}
	return intSection2[0] <= intSection1[0] && intSection2[1] >= intSection1[1]

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
