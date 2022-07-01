package day7

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	sa := strings.Split(in, " ")
	high := math.MinInt
	low := math.MaxInt
	for _, v := range sa {
		si, _ := strconv.Atoi(v)
		if si > high {
			high = si
		}
		if si < low {
			low = si
		}
	}
	// Code here or
	return fmt.Sprintf("%d %d", high, low)
}

func FindOutlier(integers []int) int {
	var r int
	l := len(integers)
	odd := []int{}
	even := []int{}
	for i := 0; i < l; i++ {
		if integers[i]%2 == 0 {
			even = append(even, i)
		} else {
			odd = append(odd, i)
		}
	}
	if len(even) > 1 {
		r = integers[odd[0]]
	} else {
		r = integers[even[0]]
	}
	return r
}

func TwoOldestAges(ages []int) [2]int {
	sort.Ints(ages)
	l := len(ages)
	return [2]int{ages[l-2], ages[l-1]}
}

func InArray(array1 []string, array2 []string) []string {
	r := []string{}
	for _, v := range array1 {
		if contains(r, v) {
			break
		}
		for _, v2 := range array2 {
			if strings.Contains(v2, v) {
				r = append(r, v)
				break
			}
		}
	}
	sort.Strings(r)
	return r
}

func contains(elms []string, e string) bool {
	for _, v := range elms {
		if v == e {
			return true
		}
	}
	return false
}
