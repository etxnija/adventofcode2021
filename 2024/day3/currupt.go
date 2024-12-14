package day3

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"nils/adventofcode2021/2023/data"
)

// mul\(\d{1,3},\d{1,3}\)"
// mul\(\(d{1,3}\),\(d{1,3}\)\)

func calculate(file string) int {
	log.Println("file", file)
	res, err := data.ReadData(file, processRow)
	if err != nil {
		log.Println(err)
	}
	// log.Println(res)
	sum := 0
	for _, row := range res {
		for _, v := range row {
			sum = sum + v[0]*v[1]
		}
	}

	return sum
}

func calculate2(file string) int {
	log.Println("file111", file)
	res, err := data.ReadData(file, processRow2)
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
	sum := 0
	for _, row := range res {
		for _, v := range row {
			log.Println(v)
			sum = sum + v[0]*v[1]
		}
	}

	return sum
}

func processRow(row string) ([][]int, error) {
	// log.Printf("%q\n", pattern.FindAll([]byte(row), -1))
	return extractPairs(row)
}

func extractPairs(row string) ([][]int, error) {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	match := re.FindAllStringSubmatch(row, -1)
	result := make([][]int, len(match))

	for i, m := range match {

		b1, err := strconv.Atoi(m[1])
		if err != nil {
			log.Printf("not a number %s : %s", m[1], err)
			return nil, err
		}

		b2, err := strconv.Atoi(m[2])
		if err != nil {
			log.Printf("not a number %s : %s", m[2], err)
			return nil, err
		}
		result[i] = []int{b1, b2}
	}
	return result, nil
}

func processRow2(row string) ([][]int, error) {
	log.Println("processRow2")
	filteredRow := filterRow(row)
	return extractPairs(filteredRow)
}

func filterRow(row string) string {
	var filterdRow strings.Builder
	reDo := regexp.MustCompile(`do\(\)`)
	reDont := regexp.MustCompile(`don't\(\)`)

	for len(row) > 0 {
		locDont := reDont.FindStringIndex(row)
		if locDont == nil {
			filterdRow.WriteString(row)
			break
		}
		if locDont != nil {

			filterdRow.WriteString(row[:locDont[0]])
			row = row[locDont[1]:]
		}
		locDo := reDo.FindStringIndex(row)
		// last don't with no do after
		if locDont != nil && locDo == nil {
			row = ""
		}

		if locDo != nil {
			row = row[locDo[1]:]
			continue
		}
		log.Println("fall though")

	}
	return filterdRow.String()
}
