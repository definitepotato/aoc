package main

import (
	"fmt"

	"util"
)

func Part1(input []string) int {
	p1 := 0
	for i := 1; i < len(input); i++ {
		left := util.Stoi(input[i])
		right := util.Stoi(input[i-1])

		if left > right {
			p1++
		}
	}
	// fmt.Printf("Part 1: %d\n", p1)
	return p1
}

func Part2(input []string) int {
	p2 := 0
	// Sliding windows.
	for i := 0; i < len(input)-3; i++ {
		windowA := util.Stoi(input[i]) + util.Stoi(input[i+1]) + util.Stoi(input[i+2])
		windowB := util.Stoi(input[i+1]) + util.Stoi(input[i+2]) + util.Stoi(input[i+3])

		if windowB > windowA {
			p2++
		}
	}
	// fmt.Printf("Part 2: %d\n", p2)
	return p2
}

func main() {
	input := util.ReadFile("input.txt")

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
