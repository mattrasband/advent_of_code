package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var elfs sort.IntSlice
	i := 0
	for scanner.Scan() {
		if len(elfs) < i+1 {
			elfs = append(elfs, 0)
		}

		if text := scanner.Text(); text != "" {
			cals, err := strconv.Atoi(text)
			if err != nil {
				log.Printf("unable to convert int: %s", err)
				continue
			}
			elfs[i] += cals
		} else {
			i++
		}
	}
	elfs.Sort()

	fmt.Println(elfs[len(elfs)-1])

	parts := elfs[len(elfs)-3:]
	var amt int
	for _, part := range parts {
		amt += part
	}
	fmt.Println(amt)
}
