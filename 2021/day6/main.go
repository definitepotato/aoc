package main

import (
	"fmt"
	"strings"
	"util"
)

func convertSchool(fish []string) []int {
	var newfish []int

	for _, v := range strings.Split(fish[0], ",") {
		newfish = append(newfish, util.Stoi(v))
	}

	return newfish
}

func convertSchoolBig(fish []string) []int {
	// Different data structure.
	// Track number of days, then track number of fish in that day.
	var newfish = make([]int, 9) // Up to 9 days in life of fish.

	for _, v := range strings.Split(fish[0], ",") {
		daysLeft := util.Stoi(v)
		newfish[daysLeft]++
	}

	return newfish
}

func scanSchool(fish []int) []int {
	var newfish []int

	for i := range fish {
		if fish[i] == 0 {
			newfish = append(newfish, 8)
			fish[i] = 6
		} else {
			fish[i]--
		}
	}

	return append(fish, newfish...)
}

func scanSchoolBig(fish []int) []int {
	var next = make([]int, 9)

	for i := 1; i < 9; i++ {
		next[i-1] = fish[i]
	}

	next[6] += fish[0]
	next[8] += fish[0]

	return next
}

func countFishBig(fish []int) int {
	allFish := 0
	for _, n := range fish {
		allFish += n
	}

	return allFish
}

func main() {
	input := util.ReadFile("input.txt")
	fish := convertSchool(input)
	biggerFish := convertSchoolBig(input)

	for j := 0; j < 80; j++ {
		fish = scanSchool(fish)
	}

	for j := 0; j < 256; j++ {
		biggerFish = scanSchoolBig(biggerFish)
	}

	fmt.Printf("Part 1: %d\n", len(fish))
	fmt.Printf("Part 2: %d\n", countFishBig(biggerFish))
}
