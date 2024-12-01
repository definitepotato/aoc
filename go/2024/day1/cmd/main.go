package main

import (
	"fmt"
	"helpers"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file := helpers.ReadFile("../input.txt")

	leftString, rightString := []string{}, []string{}

	for _, line := range file {
		lineSplit := strings.Split(line, "   ")
		leftString = append(leftString, lineSplit[0])
		rightString = append(rightString, lineSplit[1])
	}

	// convert slice of strings to slice of ints
	stringLen := len(leftString)
	leftInt, rightInt := make([]int, stringLen), make([]int, stringLen)

	for i := 0; i < stringLen; i++ {
		leftInt[i], _ = strconv.Atoi(leftString[i])
		rightInt[i], _ = strconv.Atoi(rightString[i])
	}

	// sort ints
	sort.Ints(leftInt)
	sort.Ints(rightInt)

	// part1: calculate distance
	totalDistance := 0.00
	for i := 0; i < len(leftInt); i++ {
		// why doesn't go have a built-in abs func for integers?
		left := float64(leftInt[i])
		right := float64(rightInt[i])

		totalDistance += math.Abs(left - right)
	}

	fmt.Printf("Part 1: %d\n", int(totalDistance))

	// part2: calculate similarity
	totalSimilarity := 0
	countSimilarity := 0
	for i := 0; i < len(leftInt); i++ {
		for j := 0; j < len(rightInt); j++ {
			if leftInt[i] == rightInt[j] {
				countSimilarity += 1
			}
		}
		totalSimilarity += leftInt[i] * countSimilarity
		countSimilarity = 0
	}

	fmt.Printf("Part 2: %d\n", totalSimilarity)
}
