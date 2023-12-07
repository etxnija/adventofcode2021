package day5

import (
	"bufio"
	"math"
	"nils/adventofcode2021/2023/data"
	"strconv"
	"strings"
)

func solve(in string) (int, error) {

	res, err := data.ReadFile(in, processRow)
	if err != nil {
		return 0, err
	}
	// log.Println(res)

	var start []int
	for i, r := range res {
		if len(r.Seeds) > 0 {
			start = r.Seeds
			if i == 0 {
				res = res[1:]
			} else {
				res = append(res[:i-1], res[i+1:]...)
			}
			break
		}
	}

	min := math.MaxInt
	for _, s := range start {
		converter := convert(res, "seed", s)
		if converter < min {
			min = converter
		}
	}

	return min, nil
}

// assume order
func convert(res []result, sourceName string, source int) int {

	ls := source
	sn := sourceName
	for {
		s, ok := getNext(res, sn)
		if !ok {
			return ls
		}
		ls = s.convertToDest(ls)
		sn = s.To
	}

}

func getNext(res []result, sourceName string) (result, bool) {
	ok := false
	var firstSource result
	for _, r := range res {
		if sourceName == r.From {
			firstSource = r
			ok = true
			break
		}
	}
	return firstSource, ok
}

func processRow(scanner bufio.Scanner) ([]result, error) {
	var res []result
	var r result
	for scanner.Scan() {
		row := strings.TrimSpace(scanner.Text())
		if strings.Contains(row, "-to-") {
			mp := strings.Split(strings.Split(row, " ")[0], "-to-")
			r = result{
				From: strings.TrimSpace(mp[0]),
				To:   strings.TrimSpace(mp[1]),
			}
			continue
		}
		if len(row) > 0 && r.From != "" {
			sa := strings.Split(row, " ")
			lowRange, _ := strconv.Atoi(sa[0])
			highRange, _ := strconv.Atoi(sa[1])
			rangeLength, _ := strconv.Atoi(sa[2])
			r.ranges = append(r.ranges,
				ranges{
					DestinationRange: lowRange,
					SourceRange:      highRange,
					RangeLength:      rangeLength,
				})
		} else {
			if r.From != "" {
				res = append(res, r)
			}
		}
		if strings.Contains(row, "seeds:") {

			ss := strings.Split(row, ":")
			sa := strings.Split(strings.TrimSpace(ss[1]), " ")
			for _, seed := range sa {
				s, _ := strconv.Atoi(seed)
				r.Seeds = append(r.Seeds, s)
			}
			res = append(res, r)

		}

	}
	if r.From != "" {
		res = append(res, r)
	}
	return res, nil
}

type result struct {
	Seeds  []int
	From   string
	To     string
	ranges []ranges
}

func (r result) convertToDest(source int) int {
	for _, r := range r.ranges {
		if source >= r.SourceRange && source <= r.SourceRange+r.RangeLength {
			loc := source - r.SourceRange
			return r.DestinationRange + loc
		}

	}
	return source
}

type ranges struct {
	DestinationRange int
	SourceRange      int
	RangeLength      int
}
