package day2

import (
	"log/slog"
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

func processRow(row string) (bool, error) {
	// slog.Info("row", slog.Any("vals", row))
	vals := strings.Fields(row)
	valInt := make([]int, len(vals))
	for i, v := range vals {
		vi, _ := strconv.Atoi(v)
		valInt[i] = vi
	}

	return (allSameDirection(valInt) && withinAdjacent(valInt)), nil
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func allSameDirection(values []int) bool {
	plus := 0
	minus := 0
	for i := 0; i < len(values)-1; i++ {
		diff := (values[i+1] - values[i])
		if diff < 0 {
			minus++
		}
		if diff > 0 {
			plus++
		}
	}

	return !(plus > 0 && minus > 0)
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
