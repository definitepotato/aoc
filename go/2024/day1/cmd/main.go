package main

import (
	"fmt"
	"helpers"
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
	totalDistance := 0
	for i := 0; i < len(leftInt); i++ {
		if rightInt[i] > leftInt[i] {
			totalDistance += rightInt[i] - leftInt[i]
		}

		if rightInt[i] < leftInt[i] {
			totalDistance += leftInt[i] - rightInt[i]
		}
	}

	fmt.Printf("Part 1: %d\n", totalDistance)

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
