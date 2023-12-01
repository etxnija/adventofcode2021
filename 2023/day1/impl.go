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
	cv, err := calibrationValues(in)
	if err != nil {
		return -1, err
	}
	sum := 0
	for _, i := range cv {
		sum += i
	}

	return sum, nil
}

func calibrationValues(input string) ([]int, error) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var cals []int

	for scanner.Scan() {
		row := strings.TrimSpace(scanner.Text())
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
		if err != nil {
			return []int{}, err
		}
		cals = append(cals, cal)
	}
	return cals, nil
}
