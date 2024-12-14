package day2

import (
	"log/slog"
	"slices"
	"strconv"
	"strings"

	"nils/adventofcode2021/2023/data"
)

func safeReports(file string) int {
	slog.Info("file", slog.Any("file", file))
	res, err := data.ReadData(file, processRow)
	if err != nil {
		slog.Error("Error reading data", slog.Any("err", err))
	}

	sum := 0
	for _, v := range res {
		if v {
			sum++
		}
	}

	return sum
}

func safeReports2(file string) int {
	slog.Info("file", slog.Any("file", file))
	res, err := data.ReadData(file, processRow2)
	if err != nil {
		slog.Error("Error reading data", slog.Any("err", err))
	}

	sum := 0
	for _, v := range res {
		if v {
			sum++
		}
	}

	return sum
}

func processRow2(row string) (bool, error) {
	vals := strings.Fields(row)
	valInt := make([]int, len(vals))
	for i, v := range vals {
		vi, _ := strconv.Atoi(v)
		valInt[i] = vi
	}
	return allSameDirectionAndAdjacentWithDamper(valInt), nil
}

func processRow(row string) (bool, error) {
	// slog.Info("row", slog.Any("vals", row))
	vals := strings.Fields(row)
	valInt := make([]int, len(vals))
	for i, v := range vals {
		vi, _ := strconv.Atoi(v)
		valInt[i] = vi
	}

	return allSameDirectionAndAdjacent(valInt), nil
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func allSameDirection(values []int) bool {
	direction := 0
	fails := 0
	for i := 0; i < len(values)-1; i++ {
		switch diff := (values[i+1] - values[i]); {
		case diff < 0:
			if direction == 0 {
				direction = -1
				continue
			}
			if direction == 1 {
				fails++
			}
		case diff > 0:
			if direction == 0 {
				direction = 1
				continue
			}
			if direction == -1 {
				fails++
			}
		default:
			continue
		}
	}

	return fails < 1
}

func allSameDirectionAndAdjacentWithDamper(values []int) bool {
	if allSameDirectionAndAdjacent(values) {
		return true
	}
	for i := 0; i < len(values); i++ {
		tempVals := RemoveIndex(values, i)
		if allSameDirectionAndAdjacent(tempVals) {
			return true
		}
	}

	return false
}

func RemoveIndex(s []int, i int) []int {
	ns := slices.Clone(s)
	return append(ns[:i], ns[i+1:]...)
}

func allSameDirectionAndAdjacent(values []int) bool {
	direction := 0
	fails := 0
	for i := 0; i < len(values)-1; i++ {
		diff := (values[i+1] - values[i])
		switch {
		case diff < 0:
			if direction == 0 {
				direction = -1
			}
			if direction == 1 {
				fails++
			}
		case diff > 0:
			if direction == 0 {
				direction = 1
			}
			if direction == -1 {
				fails++
			}
		case diff == 0:
			fails++
		}
		diffAbs := Abs(diff)
		if diffAbs < 1 || diffAbs > 3 {
			fails++
		}
	}

	return fails < 1
}

func withinAdjacent(values []int) bool {
	for i := 0; i < len(values)-1; i++ {
		diff := Abs((values[i+1] - values[i]))
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}
