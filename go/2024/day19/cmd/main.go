package main

import (
	"fmt"
	"helpers"
	"strings"
)

const Towels = "r, wr, b, g, bwu, rb, gb, br"

func completePattern(pattern string) bool {
	if pattern == "" {
		return true
	}

	done := false
	for _, towel := range strings.Split(Towels, ", ") {
		if done {
			return done
		}

		if strings.HasPrefix(pattern, towel) {
			done = completePattern(pattern[len(towel):])
		}
	}

	return false
}

func main() {
	input := helpers.ReadFile("../input.txt")
	part1 := 0

	for _, pattern := range input[2:] {
		if completePattern(pattern) {
			part1 += 1
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
}
