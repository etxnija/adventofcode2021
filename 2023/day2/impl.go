package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type color int64

func (c color) String() string {
	switch c {
	case red:
		return "red"
	case green:
		return "green"
	case blue:
		return "blue"
	default:
		return ""
	}
}

func GetColor(s string) (color, error) {
	switch s {
	case "red":
		return red, nil
	case "green":
		return green, nil
	case "blue":
		return blue, nil
	default:
		return red, fmt.Errorf("No match for %s", s)
	}
}

const (
	red = iota
	blue
	green
)

func possible(in string, limits map[color]int) (int, error) {
	cv, err := possibleGames(in, processRow)
	if err != nil {
		return -1, err
	}
	sum := 0
out:
	for i := 1; i < len(cv)+1; i++ {

		for _, vv := range cv[i-1] {
			if vv[green] > limits[green] {
				continue out
			}
			if vv[red] > limits[red] {
				continue out
			}
			if vv[blue] > limits[blue] {
				continue out
			}

		}
		sum += i
	}
	// log.Println(cv)

	return sum, nil
}

func min(in string) (int, error) {
	cv, err := possibleGames(in, processRow)
	if err != nil {
		return -1, err
	}
	sum := 0
	for i := 1; i < len(cv)+1; i++ {
		rmax := 0
		bmax := 0
		gmax := 0
		for _, vv := range cv[i-1] {
			if vv[green] > gmax {
				gmax = vv[green]
			}
			if vv[red] > rmax {
				rmax = vv[red]
			}
			if vv[blue] > bmax {
				bmax = vv[blue]
			}

		}
		sum += (rmax * bmax * gmax)
	}
	// 	// log.Println(cv)

	return sum, nil
}

func possibleGames[T any](input string, rowFunc func(row string) (T, error)) ([]T, error) {
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

func processRow(row string) ([]map[color]int, error) {
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green

	split := strings.Split(row, ":") // "Game 1" "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	// game := strings.Split(split[0], " ")[1] // "Game" "1"
	// g, err := strconv.Atoi(game)
	// if err != nil {
	// 	return data, err
	// }

	return pulls(split[1]), nil
}

func pulls(s string) []map[color]int {
	p := strings.Split(s, ";") // "3 blue, 2 red" "1 red, 2 green" "...."
	var ps []map[color]int
	for _, pull := range p {
		hands := strings.Split(pull, ",") // "3 blue"
		ma := make(map[color]int)
		for _, h := range hands {
			c := strings.Split(strings.TrimSpace(h), " ")
			n, err := strconv.Atoi(c[0])
			if err != nil {
				log.Println(err)
			}
			col, err := GetColor(c[1])
			if err != nil {
				log.Println(err)
			}
			ma[col] = n

		}
		ps = append(ps, ma)

	}

	return ps

}
