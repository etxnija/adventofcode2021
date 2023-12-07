package day5

import (
	"bufio"
	"log"
	"nils/adventofcode2021/2023/data"
	"strconv"
	"strings"
)

func solve(in string) (int, error) {

	res, err := data.ReadFile(in, processRow)
	if err != nil {
		return 0, err
	}
	log.Println(res)
	return 0, nil
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
			sa := strings.Split(ss[1], " ")
			for _, seed := range sa {
				s, _ := strconv.Atoi(seed)
				r.seeds = append(r.seeds, s)
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
	seeds  []int
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
