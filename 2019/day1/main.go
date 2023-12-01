package day1

import (
	"fmt"
	"util"
)

func CalculateFuel(mass int) int {
	return mass/3 - 2
}

func CalculateRecurseFuel(mass int, fuel int) int {
	newMass := CalculateFuel(mass)

	if newMass <= 0 {
		return fuel
	}

	fuel += newMass
	return CalculateRecurseFuel(newMass, fuel)
}

func main() {
	mass := util.ReadFile("input.txt")

	// Part one.
	var fuelRequirement int
	for i := 0; i < len(mass); i++ {
		fuelRequirement += CalculateFuel(util.Stoi(mass[i]))
	}
	fmt.Println(fuelRequirement)

	// Part two.
	var fuelRequirementRecurse int
	for i := 0; i < len(mass); i++ {
		fuelRequirementRecurse += CalculateRecurseFuel(util.Stoi(mass[i]), 0)
	}
	fmt.Println(fuelRequirementRecurse)
}
