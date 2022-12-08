package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func fileToRemove(input string) int {

	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	location := []string{}
	// root := Tree{}
	files := collectFiles(scanner, location)
	log.Println(files)
	sorted := sortFiles(files)

	sn := spaceNeeded(sorted)
	need := 70000000
	for _, v := range sorted {
		if v > sn && v < need {
			need = v
		}
	}

	return need
}

func spaceNeeded(files map[string]int) int {
	fileSystemSize := 70000000
	neededFreeSpace := 30000000

	cf := fileSystemSize - files["root"]
	if cf > neededFreeSpace {
		return 0
	}

	return neededFreeSpace - cf
}

func flatTree(input string) int {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	location := []string{}
	// root := Tree{}
	files := collectFiles(scanner, location)
	sorted := sortFiles(files)
	tot := 0
	for _, v := range sorted {
		if v < 100000 {
			tot = tot + v
		}

	}
	return tot
}

func collectFiles(scanner *bufio.Scanner, location []string) map[string]int {
	files := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Split(line, " ")
		if args[0] == "$" {
			if args[1] == "cd" {
				if args[2] == ".." && len(location) > 0 {
					location = location[:len(location)-1]
					continue
				}
				location = append(location, args[2])
				continue
			}
		}
		if size, err := strconv.Atoi(args[0]); err == nil {
			files[locationPath(location)+args[1]] = size
		}

	}
	return files
}

func sortFiles(f map[string]int) map[string]int {
	s := make(map[string]int)
	for k, v := range f {
		sk := strings.Split(k, "/")
		sk = sk[:len(sk)-1]
		folder := ""
		for _, p := range sk {
			folder = folder + p
			if size, ok := s[folder]; ok {
				s[folder] = size + v
				continue
			}
			s[folder] = v
		}
	}
	return s
}

func locationPath(l []string) string {
	var s string

	for _, v := range l {
		if v == "/" {
			s = "root/"
			continue
		}
		s = fmt.Sprintf("%s%s/", s, v)
	}
	return s
}
