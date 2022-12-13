package day8

import (
	"bufio"
	"log"
	"os"
)

func visableTrees(input string) int {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	ig := intGrid(scanner)
	tot := 4*len(ig) - 4
	var lineOfSight []int
	for i := 1; i < len(ig)-1; i++ {
	column:
		for j := 1; j < len(ig)-1; j++ {
			tree := ig[i][j]
			//up
			lineOfSight = make([]int, i)
			for k := 0; k < i; k++ {
				lineOfSight[k] = ig[i-k-1][j]
			}
			if !anyTaller(lineOfSight, tree) {
				tot++
				continue column
			}
			//down
			lineOfSight = make([]int, len(ig)-i-1)
			for k := 0; k < len(ig)-i-1; k++ {
				lineOfSight[k] = ig[i+k+1][j]
			}
			if !anyTaller(lineOfSight, tree) {
				tot++
				continue column
			}
			//left
			lineOfSight = make([]int, len(ig)-j-1)
			for k := 0; k < len(ig)-j-1; k++ {
				lineOfSight[k] = ig[i][j+k+1]
			}
			if !anyTaller(lineOfSight, tree) {
				tot++
				continue column
			}

			//right
			lineOfSight = make([]int, j)
			for k := 0; k < j; k++ {
				lineOfSight[k] = ig[i][j-k-1]
			}
			if !anyTaller(lineOfSight, tree) {
				tot++
				continue column
			}

		}
	}

	return tot
}

func anyTaller(lineOfSight []int, tree int) bool {
	for _, v := range lineOfSight {
		if v >= tree {
			return true
		}
	}
	return false
}

func intGrid(scanner *bufio.Scanner) [][]int {
	var g [][]int
	for scanner.Scan() {
		l := scanner.Text()
		row := make([]int, len(l))
		for i := 0; i < len(l); i++ {
			t := int(l[i] - '0')
			row[i] = t
		}
		g = append(g, row)

	}
	return g
}
