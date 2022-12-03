package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	part1Totals := 0
	part2Totals := 0
	part2Slices := make([]string, 2)
	priority := func(c rune) int {
		if c > 96 {
			return int(c - 96)
		} else if c > 65 {
			return int(c - 38)
		}
		return 0
	}

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		mid := len(line) / 2
		one, two := line[:mid], line[mid:]
		if len(one) != len(two) {
			log.Fatalf("%s: mismatched lengths", line)
		}

		// part 1
		// golang doesn't have a set type for union ops :shrug:
		seen := map[rune]bool{}
		for _, c := range two {
			if strings.ContainsRune(one, c) && !seen[c] {
				seen[c] = true
				part1Totals += priority(c)
			}
		}

		idx := i % 3
		if idx == 2 {
			seen = map[rune]bool{}
			for _, c := range line {
				if strings.ContainsRune(part2Slices[0], c) && strings.ContainsRune(part2Slices[1], c) && !seen[c] {
					part2Totals += priority(c)
				}
				seen[c] = true
			}
		} else {
			part2Slices[idx] = line
		}
	}
	log.Printf("part1: %d", part1Totals)
	log.Printf("part2: %d", part2Totals)
}
