package day6

import (
	"io"
	"log"
	"os"
)

func findMarker(input string) int {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	return startOfPacketMarker(f)

}

func startOfPacketMarker(reader io.Reader) int {
	buffer := make([]byte, 1)
	var marker []string
	c := 0
	for {
		_, err := reader.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			break
		}
		yes, loc := contains(marker, string(buffer))
		if yes {
			marker = marker[loc+1:]
		}
		marker = append(marker, string(buffer))
		c++
		if len(marker) > 3 {
			return c
		}
	}
	return 0
}

func contains(s []string, r string) (bool, int) {
	for i, v := range s {
		if v == r {
			return true, i
		}
	}
	return false, -1
}
