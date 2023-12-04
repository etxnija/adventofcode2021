package day3

import (
	"math"
	"nils/adventofcode2021/2023/data"
	"strconv"
)

const special = "special"

func solve(in string) (int, error) {

	res, err := data.ReadData(in, processRow)
	sum := 0
	if err != nil {
		return 0, err
	}
	sp := [][]int{}
	for i, row := range res {
		s := row.special
		for _, j := range s {
			sp = append(sp, []int{i, j})

		}
	}
	for i, row := range res {
	part:
		for _, ia := range row.parts {
			for _, partLoc := range ia.location {
				for _, s := range sp {
					if vectorDist(s, []int{i, partLoc}) < 1.5 {
						pn, err := strconv.Atoi(ia.name)
						if err != nil {
							return 0, err
						}
						sum += pn
						continue part
					}
				}
			}
		}
	}
	// log.Println(res)
	return sum, nil
}

func solve2(in string) (int, error) {

	res, err := data.ReadData(in, processRow)
	sum := 0
	if err != nil {
		return 0, err
	}
	sp := [][]int{}
	for i, row := range res {
		s := row.special
		for _, j := range s {
			sp = append(sp, []int{i, j})

		}
	}
	// A gear is any * symbol that is adjacent to exactly two part numbers.
	gearRations := [][]string{}

	for _, s := range sp {
		gearRation := []string{}
		for i, row := range res {
			parts := row.parts
		part:
			for _, p := range parts {
				for _, l := range p.location {
					if vectorDist(s, []int{i, l}) < 1.5 {
						gearRation = append(gearRation, p.name)
						continue part
					}
				}
			}
		}
		if len(gearRation) == 2 {
			gearRations = append(gearRations, gearRation)
		}
	}

	for _, gr := range gearRations {
		g1, err := strconv.Atoi(gr[0])
		if err != nil {
			return 0, err
		}
		g2, err := strconv.Atoi(gr[1])
		if err != nil {
			return 0, err
		}
		sum += (g1 * g2)

	}

	// log.Println(res)
	return sum, nil
}

func vectorDist(p1, p2 []int) float64 {
	sum := math.Pow(float64((p1[0]-p2[0])), 2) + math.Pow(float64((p1[1]-p2[1])), 2)
	return float64(math.Sqrt(sum))
}

type result struct {
	parts   []part
	special []int
}
type part struct {
	name     string
	location []int
}

func processRow(row string) (result, error) {
	// var s [][][]int
	zero := int('0')
	nine := int('9')
	var onPart []rune
	var partLoc []int
	var spc []int
	var parts []part

	for i, char := range row {
		c := int(char)
		if c >= zero && c <= nine {
			onPart = append(onPart, char)
			partLoc = append(partLoc, i)
			if i == len(row)-1 && len(onPart) > 0 {
				parts = append(parts, part{string(onPart), partLoc})
			}
			continue
		}
		if len(onPart) > 0 {
			parts = append(parts, part{string(onPart), partLoc})
			onPart = make([]rune, 0)
			partLoc = make([]int, 0)
		}

		if char == '.' {
			continue
		}
		spc = append(spc, i)
	}
	return result{parts: parts, special: spc}, nil
}
