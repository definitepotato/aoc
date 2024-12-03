package main

import (
	"fmt"
	"helpers"
	"regexp"
	"strconv"
)

const (
	TestInput1 = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	TestInput2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
)

func getNumbersFromInstruction(input string) (int, int) {
	rgx, _ := regexp.Compile(`mul\((?P<n1>\d+),(?P<n2>\d+)\)`)
	instruction := rgx.FindStringSubmatch(input)

	leftNum, _ := strconv.Atoi(instruction[1])
	rightNum, _ := strconv.Atoi(instruction[2])

	return leftNum, rightNum
}

func getInstructions(input string, rgx *regexp.Regexp) []string {
	instructions := rgx.FindAllString(input, -1)
	return instructions
}

func main() {
	inputs := helpers.ReadFile("../input.txt")

	result1 := 0
	rgx1, _ := regexp.Compile(`mul\((?P<n1>\d+),(?P<n2>\d+)\)`)
	for _, input := range inputs {
		instructions := getInstructions(input, rgx1)

		for _, instr := range instructions {
			n1, n2 := getNumbersFromInstruction(instr)
			result1 += n1 * n2
		}
	}

	fmt.Printf("Part 1: %d\n", result1)

	result2 := 0
	process := true
	rgx2, _ := regexp.Compile(`(?:(?P<dont>don't)\(\)|(?<do>do)\(\))|mul\((?P<n1>\d+),(?P<n2>\d+)\)`)

	for _, input := range inputs {
		instructions := getInstructions(input, rgx2)

		for _, instr := range instructions {
			if instr == "don't()" {
				process = false
				continue
			}

			if instr == "do()" {
				process = true
				continue
			}

			if process {
				n1, n2 := getNumbersFromInstruction(instr)
				result2 += n1 * n2
			}
		}
	}

	fmt.Printf("Part 2: %d\n", result2)
}
