package main

import (
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	raw, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	sample := string(raw)
	log.Printf("part1: %d", findStartByUniqueness(sample, 4))
	log.Printf("part2: %d", findStartByUniqueness(sample, 14))
}

func findStartByUniqueness(input string, uniqueCount int) int {
MAINLOOP:
	for i := 0; i < len(input)-uniqueCount; i++ {
		set := map[rune]bool{}
		for _, c := range input[i : i+uniqueCount] {
			if set[c] {
				continue MAINLOOP
			}
			set[c] = true
		}
		return i + uniqueCount
	}
	return -1
}
