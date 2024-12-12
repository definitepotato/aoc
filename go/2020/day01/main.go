package main

import (
	"fmt"
	"util"
)

func is2020(x, y int) bool {
	sum := x + y
	return sum == 2020
}

func part1() int {
	report := util.ReadFile("input.txt")

	for _, num1 := range report {
		for _, num2 := range report {
			if is2020(util.Stoi(num1), util.Stoi(num2)) {
				return util.Stoi(num1) * util.Stoi(num2)
			}
		}
	}

	return 0
}

//  1. Given an array of length [n] and a sum 2020
//  2. Create three nested loops:
//     2a. Loop counter i runs from start to end
//     2b. Loop counter j runs from i+1 to end
//     2c. Loop counter k runs from j+1 to end
//  3. The counter of each loop represents the index of 3 elements
//  4. If the sum of ith, jth and kth elements equal 2020, return
func part2() int {
	report := util.ReadFile("input.txt")

	for i := 0; i < len(report); i++ { // Loop counter i
		for j := i + 1; j < len(report); j++ { // Loop counter j
			for k := j + 1; k < len(report); k++ { // Loop counter k
				if util.Stoi(report[i])+util.Stoi(report[j])+util.Stoi(report[k]) == 2020 {
					return util.Stoi(report[i]) * util.Stoi(report[j]) * util.Stoi(report[k])
				}
			}
		}
	}

	return 0
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
