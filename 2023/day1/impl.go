package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calibrate(in string) (int, error) {
	cv, err := calibrationValues(in, processRow)
	if err != nil {
		return -1, err
	}
	sum := 0
	for _, i := range cv {
		sum += i
	}

	return sum, nil
}

func calibrationValues[T any](input string, rowFunc func(row string) (T, error)) ([]T, error) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var cals []T

	for scanner.Scan() {
		row := strings.TrimSpace(scanner.Text())
		cal, err := rowFunc(row)
		if err != nil {
			return []T{}, err
		}
		cals = append(cals, cal)
	}
	return cals, nil
}

func processRow(row string) (int, error) {
	sone := ""
	stwo := ""
	num := ""
	for i := 0; i < len(row); i++ {
		char := string(row[i])
		_, err := strconv.Atoi(char)
		if err != nil {
			ok := false
			num, ok = matchNumber(row, i)
			if !ok {
				continue
			}
		} else {
			num = string(char)
		}
		if sone == "" {
			sone = num
		}

		stwo = num
	}
	cal, err := strconv.Atoi(fmt.Sprintf("%s%s", sone, stwo))
	if err != nil {
		return 0, nil
	}
	return cal, err
}

func matchNumber(row string, i int) (string, bool) {
	for _, v := range numbers {
		s := v.String()
		if len(row) >= i+len(s) && s == row[i:i+len(s)] {
			return strconv.Itoa(int(v)), true
		}
	}
	return "", false
}

func processRowTwo(row string) (int, error) {
	sone := ""
	stwo := ""
	for i := 0; i < len(row); i++ {
		char := string(row[i])
		_, err := strconv.Atoi(char)
		if err != nil {
			continue
		}
		if sone == "" {
			sone = string(char)
		}
		stwo = string(char)
	}

	cal, err := strconv.Atoi(fmt.Sprintf("%s%s", sone, stwo))
	return cal, err
}
