package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fuelRequired(mass int) int {
	return int(float64(mass/3)) - 2
}

func fuelRequiredRecurse(initialMass int) int {
	required := fuelRequired(initialMass)
	if required <= 0 {
		return 0
	}

	return required + fuelRequiredRecurse(required)
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf("Unable to open the file: %s\n", err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		v, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			fmt.Printf("Unable to parse int: %s\n", err)
			return
		}

		sum += fuelRequiredRecurse(int(v))
	}

	fmt.Printf("Required: %d\n", int(sum))
}
