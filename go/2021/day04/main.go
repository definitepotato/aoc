package main

import (
	"fmt"
	"strings"
	"util"
)

func checkWin(b []int) bool {
	win := true

	// Check rows.
	for j := 0; j < 25; j += 5 { // [1:0] [2:5] [3:10] [4:15] [5:20]
		win = true
		for i := j; i < j+5; i++ { // [1:0,1,2,3,4] [2:5,6,7,8,9] [3:10,11,12,13,14] [4:15,16,17,18,19] [5:20,21,22,23,24]
			if b[i] != 0 {
				win = false
			}
		}

		if win {
			return true
		}
	}

	// Check columns.
	for j := 0; j < 5; j++ { // [1:0] [2:1] [3:2] [4:3] [5:4]
		win = true
		for i := j; i < 25; i += 5 { // [1:0,5,10,15,20] [2:1,6,11,16,21] [3:2,7,12,17,22] [4:3,8,13,18,23] [5:4,9,14,19,24]
			if b[i] != 0 {
				win = false
			}
		}

		if win {
			return true
		}
	}

	return false
}

func part1() {
	input := util.ReadFile("input.txt")

	// Get numbers from input.
	var numbers []int
	for _, n := range strings.Split(input[0], ",") {
		numbers = append(numbers, util.Stoi(n))
	}

	// Get boards from input.
	var boards [][]int
	for i := 2; i < len(input); i += 6 {
		var board []int
		for _, s := range strings.Split(strings.Join(input[i:i+5], " "), " ") {
			if s == "" { // Ignore empty elements.
				continue
			}
			board = append(board, util.Stoi(s))
		}
		boards = append(boards, board)
	}

	for _, n := range numbers {
		for _, b := range boards {
			for i, v := range b {
				if v == n {
					b[i] = 0
					break
				}
			}

			if checkWin(b) {
				sum := 0
				for _, j := range b {
					sum += j
				}
				fmt.Println("Part 1:", sum*n)
				return
			}
		}
	}
}

func part2() {
	input := util.ReadFile("input.txt")

	// Get numbers from input.
	var numbers []int
	for _, n := range strings.Split(input[0], ",") {
		numbers = append(numbers, util.Stoi(n))
	}

	// Get boards from input.
	var boards [][]int
	for i := 2; i < len(input); i += 6 {
		var board []int
		for _, s := range strings.Split(strings.Join(input[i:i+5], " "), " ") {
			if s == "" { // Ignore empty elements.
				continue
			}
			board = append(board, util.Stoi(s))
		}
		boards = append(boards, board)
	}

	boardWin := make([]bool, len(boards))

	for _, n := range numbers {
		for b := range boards {

			if boardWin[b] {
				continue
			}

			for i, v := range boards[b] {
				if v == n {
					boards[b][i] = 0
					break
				}
			}

			if checkWin(boards[b]) {
				sum := 0
				for _, j := range boards[b] {
					sum += j
				}
				fmt.Println(n, boards[b], sum*n)
				boardWin[b] = true
			}
		}
	}
}

func main() {
	part1()
	part2()
}
