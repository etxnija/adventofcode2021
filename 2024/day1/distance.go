package day1

import (
	"fmt"
	"log/slog"
	"sort"
	"strconv"
	"strings"

	"nils/adventofcode2021/2023/data"
)

func distance(file string) int {
	slog.Info("file ", slog.String("file ", file))

	res, err := data.ReadData(file, processRow)
	if err != nil {
		slog.Error("Unable to read data", slog.Any("err", err))
	}
	slog.Info("file content", slog.Any("file", res))

	list1 := make([]int, len(res))
	list2 := make([]int, len(res))

	for i := range res {
		list1[i] = res[i][0]
		list2[i] = res[i][1]
	}
	sort.Ints(list1)
	sort.Ints(list2)
	sum := 0

	for i := range list1 {
		sum = sum + Abs(list1[i]-list2[i])
	}

	return sum
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type result []int

func processRow(row string) (result, error) {
	vals := strings.Fields(row)
	if len(vals) > 2 {
		return result{}, fmt.Errorf("To many fields on the row")
	}
	v1, err := strconv.Atoi(vals[0])
	if err != nil {
		return result{}, fmt.Errorf("v1 not a int")
	}
	v2, err := strconv.Atoi(vals[1])
	if err != nil {
		return result{}, fmt.Errorf("v2 not a int")
	}
	return result{v1, v2}, nil
}
