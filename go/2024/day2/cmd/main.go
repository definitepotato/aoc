package main

import (
	"fmt"
	"helpers"
	"math"
	"strings"
)

func diffIsSafe(x, y int) bool {
	xFloat, yFloat := float64(x), float64(y)
	difference := math.Abs(xFloat - yFloat)

	if difference >= 1 && difference <= 3 {
		return true
	}

	return false
}

func isIncreasing(list []string) bool {
	listInt := helpers.SliceStringToInt(list)

	for idx := range listInt {
		// Toss the first value.
		if idx == 0 {
			continue
		}

		if listInt[idx-1] > listInt[idx] {
			return false
		}

		if !diffIsSafe(listInt[idx-1], listInt[idx]) {
			return false
		}
	}

	return true
}

// Yes this essentially a copy of isIncreasing with 1 char difference
// I'm okay with this, I want it to read nicely in main()
func isDecreasing(list []string) bool {
	listInt := helpers.SliceStringToInt(list)

	for idx := range listInt {
		// Toss the first value.
		if idx == 0 {
			continue
		}

		if listInt[idx-1] < listInt[idx] {
			return false
		}

		if !diffIsSafe(listInt[idx-1], listInt[idx]) {
			return false
		}
	}

	return true
}

func removeIdx(slice []string, idx int) []string {
	// go slices are pointers so we need to deep copy here
	// to avoid modifying the original slice
	c := make([]string, len(slice))
	copy(c, slice)

	return append(c[:idx], c[idx+1:]...)
}

func dampen(report []string) bool {
	for idx := 0; idx < len(report); idx++ {
		modifiedReport := removeIdx(report, idx)

		if isIncreasing(modifiedReport) {
			return true
		}

		if isDecreasing(modifiedReport) {
			return true
		}

	}

	return false
}

func main() {
	file := helpers.ReadFile("../input.txt")

	safeReports := 0
	safeDampenedReports := 0
	isSafe := false
	for _, report := range file {
		reportArray := strings.Split(report, " ")

		isSafe = isIncreasing(reportArray)
		if !isSafe {
			isSafe = isDecreasing(reportArray)
		}

		if isSafe {
			safeReports += 1
		}

		if !isSafe {
			if dampen(reportArray) {
				safeDampenedReports += 1
			}
		}
	}

	fmt.Printf("Part 1: %d\n", safeReports)
	fmt.Printf("Part 2: %d\n", safeDampenedReports+safeReports)
}
