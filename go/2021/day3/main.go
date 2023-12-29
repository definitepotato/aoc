package main

import (
	"fmt"
	"strconv"
	"util"
)

func binaryToDecimal(binary string) int64 {
	decimal, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		panic(err)
	}

	return decimal
}

func getRating(list []string, rating string) int64 {
	// We assume we won't exceed range boundaries because
	// we should get down to 1 element in our list before
	// we reach the end of the list.

	// Loop until we're down to one value in our list.
	// `i` will hold the current horizontal position.
	for i := 0; len(list) > 1; i++ {
		var countZero int
		var countOne int

		// Count zeroes and ones while iterating list vertically.
		// `i` is the current horizontal position from outer loop.
		for _, binary := range list {
			if binary[i] == '0' {
				countZero++
			} else {
				countOne++
			}
		}

		var keep string

		// Evaluate which bit to keep based on
		// rating parameter in func signature.
		if rating == "co2" {
			if countZero > countOne {
				keep = "1"
			} else if countZero == countOne {
				keep = "0"
			} else {
				keep = "0"
			}
		}

		if rating == "oxygen" {
			if countZero > countOne {
				keep = "0"
			} else if countZero == countOne {
				keep = "1"
			} else {
				keep = "1"
			}
		}

		// Using evaluated keep value, evaluate which items
		// in the list to keep as a new list.
		// `i` is the current horizontal position from outer loop.
		var newList []string
		for _, binary := range list {
			if string(binary[i]) == keep {
				newList = append(newList, binary)
			}
		}

		// Replace current list with new evaluated slice
		// before next iteration.
		list = newList
	}

	return binaryToDecimal(list[0])
}

func part1() int64 {
	input := util.ReadFile("input.txt")
	var gamma string
	var epsilon string

	// Iterate input list based on the width
	// of the first line.
	for i := 0; i < len(input[0]); i++ {
		var countZero int
		var countOne int

		// Count zeroes and ones.
		for _, binary := range input {
			if binary[i] == '0' {
				countZero++
			} else {
				countOne++
			}
		}

		// Evaluate gamma based count, epsilon
		// will always be opposite.
		if countZero > countOne {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammaValue := binaryToDecimal(gamma)
	epsilonValue := binaryToDecimal(epsilon)

	return gammaValue * epsilonValue
}

func part2() int64 {
	input := util.ReadFile("input.txt")

	oxygen := getRating(input, "oxygen")
	co2 := getRating(input, "co2")

	return oxygen * co2
}

func main() {
	p1 := part1()
	p2 := part2()

	fmt.Printf("Part1: %d\nPart2: %d\n", p1, p2)
}
