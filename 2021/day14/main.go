package main

import (
	"fmt"
	"strings"
	"util"
)

// Count all occurences of `chars` in `pairs`.
func CountAll(chars []string, pairs []string) []int {
	final := []int{}

	for _, v := range chars {
		x := strings.Count(strings.Join(pairs, ""), v)
		final = append(final, x)
	}

	return final
}

// Match pairs to legend (insertion rules), return final polymer after insertions.
func MatchLegend(pairs []string, legend map[string]string) []string {
	final := []string{}
	last := pairs[len(pairs)-1]

	for i := range pairs {
		if i < len(pairs)-1 {
			c := pairs[i] + pairs[i+1]

			for x, y := range legend {
				if c == x {
					final = append(final, pairs[i])
					final = append(final, y)
				}
			}
		}
	}
	final = append(final, last)
	return final
}

func Part1(input []string) int {
	legend := make(map[string]string)
	ans := strings.Split(input[0], "")

	// Build legend to hold pair insertion rules.
	for i := 2; i < len(input); i++ {
		left := strings.Split(input[i], " -> ")[0]
		right := strings.Split(input[i], " -> ")[1]
		legend[left] = right
	}

	for i := 0; i < 40; i++ {
		ans = MatchLegend(ans, legend)
	}

	count := CountAll(util.Unique(ans), ans)
	min, max := util.MinMax(count)

	return max - min
}

func main() {
	input := util.ReadFile("input.txt")

	fmt.Printf("Part 1: %d\n", Part1(input))
}
