package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func flatTree(input string) int {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	location := []string{}
	files := make(map[string]int)
	// root := Tree{}
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
	log.Println(files)
	sorted := sortFiles(files)
	tot := 0
	for _, v := range sorted {
		if v < 100000 {
			tot = tot + v
		}

	}
	return tot
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
