package day4

import (
	"math"
	"nils/adventofcode2021/2023/data"
	"strconv"
	"strings"
)

func solve2(in string) (int, error) {

	res, err := data.ReadData(in, processRow)
	sum := 0
	if err != nil {
		return 0, err
	}
	res = play(res, []result{}, 0)

	// log.Println(res)
	return sum, nil
}

func play(res, add []result, start int) []result {
	res = append(res, add...)
	if len(res)-1 <= start {
		return res
	}
	for i := start; i < len(res); i++ {
		winners := scratch(res[i])
		if len(winners) < 1 {
			return play(res, []result{}, i+1)
		}
		next := i + 1 + len(winners)
		if next > len(res) {
			next = len(res) - 1
		}
		if next > i {
			play(res, res[i+1:next], i+1)
		} else {
			play(res, []result{}, i+1)
		}
	}
	return []result{}
}

func scratch(r result) []int {
	var winners []int
	for _, m := range r.my {
		for _, w := range r.winning {
			if w == m {
				winners = append(winners, m)
			}
		}
	}
	return winners
}

func solve(in string) (int, error) {

	res, err := data.ReadData(in, processRow)
	sum := 0
	if err != nil {
		return 0, err
	}

	for _, r := range res {
		var winners []int
		for _, m := range r.my {
			for _, w := range r.winning {
				if w == m {
					winners = append(winners, m)
				}
			}
		}
		sum += int(math.Exp2(float64(len(winners) - 1)))

	}

	// log.Println(res)
	return sum, nil
}

func processRow(row string) (result, error) {
	// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	sp := strings.Split(row, "|")
	winingNumbers := strings.Split(sp[0], ":")[1]
	myNumbers := sp[1]

	winInt := toIntArray(winingNumbers)
	myInt := toIntArray(myNumbers)

	return result{winning: winInt, my: myInt}, nil
}

func toIntArray(s string) []int {
	var ia []int
	for i := 0; i < len(s)-2; i = i + 3 {
		subString := s[i : i+3]
		subString = strings.TrimSpace(subString)
		n, _ := strconv.Atoi(subString)
		ia = append(ia, n)

	}
	return ia
}

type result struct {
	winning []int
	my      []int
	count   int
}

func (r result) winCount() int {
	var c int
	for _, m := range r.my {
		for _, w := range r.winning {
			if w == m {
				c++
			}
		}
	}
	return c
}
