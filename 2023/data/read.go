package data

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadData[T any](input string, rowFunc func(row string) (T, error)) ([]T, error) {
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
