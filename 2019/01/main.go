package main

import (
	"fmt"
	"strconv"

	"github.com/mrasband/advent_of_code/2019/shared"
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
	sum := 0
	for line := range shared.ReadLines("./input.txt") {
		v, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			fmt.Printf("Unable to parse int: %s\n", err)
			return
		}

		sum += fuelRequiredRecurse(int(v))
	}
	fmt.Printf("Required: %d\n", int(sum))
}
