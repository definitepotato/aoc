package main

import (
	"fmt"
	"strings"
	"util"
)

type Position struct {
	horizontal int
	depth      int
	final      int
	aim        int
}

func Part1(instructions []string) int {
	p := &Position{
		horizontal: 0,
		depth:      0,
		final:      0,
	}

	for i := 0; i < len(instructions); i++ {
		instruction := strings.Split(instructions[i], " ")
		units := util.Stoi(instruction[1])

		if instruction[0] == "forward" {
			p.horizontal += units
		}

		if instruction[0] == "down" {
			p.depth += units
		}

		if instruction[0] == "up" {
			p.depth -= units
		}
	}

	p.final = p.horizontal * p.depth
	// fmt.Println("Part 1:", p.final)
	return p.final
}

func Part2(instructions []string) int {
	p := &Position{
		horizontal: 0,
		depth:      0,
		final:      0,
		aim:        0,
	}

	for i := 0; i < len(instructions); i++ {
		instruction := strings.Split(instructions[i], " ")
		units := util.Stoi(instruction[1])

		if instruction[0] == "forward" {
			p.horizontal += units
			p.depth += p.aim * units
		}

		if instruction[0] == "down" {
			p.aim += units
		}

		if instruction[0] == "up" {
			p.aim -= units
		}
	}

	p.final = p.horizontal * p.depth
	// fmt.Println("Part 2:", p.final)
	return p.final
}

func main() {
	input := util.ReadFile("input2.txt")

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
