package day3

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func diagnose(reports []string) int {
	bytesize := len(reports[0])
	log.Println("bytesize", bytesize)
	gamma := make([]int, bytesize)
	epsilon := make([]int, bytesize)
	g := 0
	e := 0
	// epailon := make([]int, 8)

	for _, v := range reports {
		bi, _ := strconv.ParseInt(v, 2, 32)
		for i := 0; i < bytesize; i++ {
			sb := bi >> i
			// log.Println(sb)
			// log.Println(bi, sb, ((sb & 1) != 0))
			if sb&1 != 0 {
				gamma[bytesize-i-1] = gamma[bytesize-i-1] + 1
			}

		}
	}
	size := len(gamma)
	for i := 0; i < size; i++ {
		mask := 1 << (size - i - 1)
		if int(len(reports)/2) < gamma[i] {
			gamma[i] = 1
			g = g | mask
			epsilon[i] = 0
		} else {
			gamma[i] = 0
			epsilon[i] = 1
			e = e | mask
		}
	}

	log.Println("Gamma:, Epsilon:", gamma, epsilon)
	log.Println("g,e:", g, e)
	// log.Panicln(len(reports) / 2)

	return g * e
}

func oxygenRating(reports []string) int {
	bytesize := len(reports[0])
	mesurements := make([]int, 0)
	for _, v := range reports {
		bi, _ := strconv.ParseInt(v, 2, 32)
		mesurements = append(mesurements, int(bi))
	}
	r := findOxygenGen(mesurements, bytesize-1)
	s := findScrubbing(mesurements, bytesize-1)
	log.Println("Reurn of the mack", r, s)
	return r[0] * s[0]

}

func findOxygenGen(m []int, bp int) []int {
	if len(m) == 1 {
		return m
	}
	setBit, unsetBit := splitSlice(m, bp, matchOnSetBit)
	if len(setBit) >= len(unsetBit) {
		return findOxygenGen(setBit, bp-1)
	} else {
		return findOxygenGen(unsetBit, bp-1)
	}
}

func findScrubbing(m []int, bp int) []int {
	if len(m) == 1 {
		return m
	}
	setBit, unsetBit := splitSlice(m, bp, matchOnSetBit)
	if len(unsetBit) <= len(setBit) {
		return findScrubbing(unsetBit, bp-1)
	} else {
		return findScrubbing(setBit, bp-1)
	}
}

func splitSlice(mesurements []int, bitPossition int, f func(int, int) bool) (matched []int, unmatched []int) {

	var match []int
	var unmatch []int
	mask := 1 << bitPossition
	for _, v := range mesurements {
		if f(v, mask) {
			match = append(match, v)
		} else {
			unmatch = append(unmatch, v)
		}

	}

	return match, unmatch
}

func matchOnSetBit(v int, mask int) bool {
	return v&mask != 0
}
func matchOnUnSetBit(v int, mask int) bool {
	return v&mask == 0
}

func readReport() []string {
	mesurments := make([]string, 0)

	file, err := os.Open("report.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		t := scanner.Text()
		mesurments = append(mesurments, t)
	}
	return mesurments
}
