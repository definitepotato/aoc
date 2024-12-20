package main

import (
	"fmt"
	"helpers"
	"strings"
)

const Towels = "r, wr, b, g, bwu, rb, gb, br"

var mem = make(map[string]int, 0)

func completePattern(pattern string) int {
	if value, ok := mem[pattern]; ok {
		return value
	}

	if pattern == "" {
		return 1
	}

	match := 0
	for _, towel := range strings.Split(Towels, ", ") {
		if strings.HasPrefix(pattern, towel) {
			match += completePattern(pattern[len(towel):])
		}
	}

	mem[pattern] = match
	return match
}

func main() {
	input := helpers.ReadFile("../test.txt")
	part1 := 0
	part2 := 0

	for _, pattern := range input[2:] {
		result := completePattern(pattern)
		if result > 0 {
			part1 += 1
		}
		part2 += result
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
