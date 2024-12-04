package main

import (
	"fmt"
	"helpers"
)

const MagicWord = "XMAS"

func isXmas(a, b, c, d byte) bool {
	// XMAS
	if a == 88 && b == 77 && c == 65 && d == 83 {
		return true
	}

	// SAMX
	if a == 83 && b == 65 && c == 77 && d == 88 {
		return true
	}

	return false
}

func main() {
	file := helpers.ReadFile("../test.txt")

	var grid []string
	for _, line := range file {
		grid = append(grid, line)
	}

	foundXmas := 0
	// iterate each row of the grid on the Y axis
	for y := 0; y < len(grid); y++ {
		// iterate a row on X axis
		for x := range grid[y] {
			// which directions can we check?
			right := x < len(grid[y])-len(MagicWord)
			left := x >= len(MagicWord)
			down := y < len(grid)-len(MagicWord)
			up := y >= len(MagicWord)

			if right {
				if isXmas(grid[y][x], grid[y][x+1], grid[y][x+2], grid[y][x+3]) {
					foundXmas += 1
				}
			}

			if left {
				if isXmas(grid[y][x], grid[y][x-1], grid[y][x-2], grid[y][x-3]) {
					foundXmas += 1
				}
			}

			if up {
				if isXmas(grid[y][x], grid[y-1][x], grid[y-2][x], grid[y-3][x]) {
					foundXmas += 1
				}
			}

			if down {
				if isXmas(grid[y][x], grid[y+1][x], grid[y+2][x], grid[y+3][x]) {
					foundXmas += 1
				}
			}

			// check diagonals
			if right && down {
				if isXmas(grid[y][x], grid[y+1][x+1], grid[y+2][x+2], grid[y+3][x+3]) {
					foundXmas += 1
				}
			}

			if right && up {
				if isXmas(grid[y][x], grid[y-1][x+1], grid[y-2][x+2], grid[y-3][x+3]) {
					foundXmas += 1
				}
			}

			if left && down {
				if isXmas(grid[y][x], grid[y+1][x-1], grid[y+2][x-2], grid[y+3][x-3]) {
					foundXmas += 1
				}
			}

			if left && up {
				if isXmas(grid[y][x], grid[y-1][x-1], grid[y-2][x-2], grid[y-3][x-3]) {
					foundXmas += 1
				}
			}

		}
	}

	fmt.Printf("Part 1: %d\n", foundXmas)
}
