package main

import (
	"testing"
)

func TestRequiredFuel(t *testing.T) {
	cases := map[string]struct {
		InitialMass int
		Expected    int
	}{
		"12":     {12, 2},
		"14":     {14, 2},
		"1969":   {1969, 654},
		"100756": {100756, 33583},
	}

	for name, tc := range cases {
		res := fuelRequired(tc.InitialMass)
		if res != tc.Expected {
			t.Errorf("[%s] Expected %d, Got %d", name, tc.Expected, res)
		}
	}
}

func TestRequiredFuelPart2(t *testing.T) {
	cases := map[string]struct {
		InitialMass int
		Expected    int
	}{
		"14":     {14, 2},
		"1969":   {1969, 966},
		"100756": {100756, 50346},
	}

	for name, tc := range cases {
		res := fuelRequiredRecurse(tc.InitialMass)
		if res != tc.Expected {
			t.Errorf("[%s] Expected %d, got %d", name, tc.Expected, res)
		}
	}
}
