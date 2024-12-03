package main

import (
	"fmt"
	"helpers"
	"regexp"
	"strconv"
)

const TestInput = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

func getNumbersFromInstruction(input string) (int, int) {
	rgx, _ := regexp.Compile(`mul\((?P<n1>\d+),(?P<n2>\d+)\)`)
	instruction := rgx.FindStringSubmatch(input)

	leftNum, _ := strconv.Atoi(instruction[1])
	rightNum, _ := strconv.Atoi(instruction[2])

	return leftNum, rightNum
}

func getInstructions(input string) []string {
	rgx, _ := regexp.Compile(`mul\((?P<n1>\d+),(?P<n2>\d+)\)`)
	instructions := rgx.FindAllString(input, -1)

	return instructions
}

func main() {
	inputs := helpers.ReadFile("../input.txt")

	result := 0
	for _, input := range inputs {
		instructions := getInstructions(input)

		for _, instr := range instructions {
			n1, n2 := getNumbersFromInstruction(instr)
			result += n1 * n2
		}
	}

	fmt.Printf("Part 1: %d\n", result)
}
