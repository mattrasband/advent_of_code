package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Play int

const (
	Rock     Play = 1
	Paper    Play = 2
	Scissors Play = 3
)

func mapPlay(input string) Play {
	switch input {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	}
	return Play(0)
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	var part1Total int
	var part2Total int
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			log.Fatalf("parts too big: %d", len(parts))
		}

		// part 1
		theirs, ours := mapPlay(parts[0]), mapPlay(parts[1])
		var round int
		win := theirs + 1
		if win > 3 {
			win = win % 3
		}
		if theirs == ours { // draw
			round = 3
		} else if ours == win {
			round = 6
		}
		part1Total += round + int(ours)

		// part2
		lose := "X"
		draw := "Y"
		winner := "Z"
		round = 0
		switch parts[1] {
		case draw:
			round = 3 + int(theirs)
		case lose:
			l := int(theirs) - 1
			if l < 1 {
				l = 3
			}
			round = l
		case winner:
			round = 6 + int(win)
		}
		part2Total += round
	}
	log.Printf("part 1: %d", part1Total)
	log.Printf("part 2: %d", part2Total)

}
