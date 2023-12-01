package main

import (
	"fmt"
	"util"
)

func crabWalk(crabs []int, pos int) int {
	var moves int

	for i := 0; i < len(crabs); i++ {
		if crabs[i] > pos {
			moves += crabs[i] - pos
		}

		if crabs[i] < pos {
			moves += pos - crabs[i]
		}
	}

	return moves
}

func crabRun(crabs []int, pos int) int {
	var allMoves int

	for i := 0; i < len(crabs); i++ {
		var moves int
		var step int

		if crabs[i] > pos {
			moves = crabs[i] - pos

			for j := 0; j < moves; j++ {
				step += j + 1
			}
			allMoves += step
		}

		if crabs[i] < pos {
			moves += pos - crabs[i]

			for j := 0; j < moves; j++ {
				step += j + 1
			}
			allMoves += step
		}
	}

	return allMoves
}

func part1() {
	input := util.ReadFile("input.txt")
	crabs := util.SliceStringToInt(input)
	_, max := util.MinMax(crabs)

	var allWalkMoves []int
	for i := 0; i < max; i++ {
		allWalkMoves = append(allWalkMoves, (crabWalk(crabs, i)))
	}
	min, _ := util.MinMax(allWalkMoves)

	fmt.Printf("Part 1: %d\n", min)
}

func part2() {
	input := util.ReadFile("input.txt")
	crabs := util.SliceStringToInt(input)
	_, max := util.MinMax(crabs)

	var allRunMoves []int
	for i := 0; i < max; i++ {
		allRunMoves = append(allRunMoves, (crabRun(crabs, i)))
	}
	min, _ := util.MinMax(allRunMoves)

	fmt.Printf("Part 2: %d\n", min)
}

func main() {
	part1()
	part2()
}
