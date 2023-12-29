package main

import (
	"fmt"
	"sort"
	"util"
)

func main() {
	calories := util.ReadFile("input2.txt")

	// Count calories per elf
	elvesAndCalories := map[int]int{}
	elfPos := 0
	for _, calorie := range calories {
		if calorie == "" {
			elfPos++
			continue
		}
		elvesAndCalories[elfPos] += util.Stoi(calorie)
	}

	// Find highest
	highest := 0
	for _, calorie := range elvesAndCalories {
		if calorie > highest {
			highest = calorie
		}
	}
	fmt.Println(highest)

	// Sum three highest
	var countedCalories []int
	for _, calorie := range elvesAndCalories {
		countedCalories = append(countedCalories, calorie)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(countedCalories)))
	// threeHighest := countedCalories[0] + countedCalories[1] + countedCalories[2]
	// fmt.Println(threeHighest)
}
