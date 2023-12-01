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
