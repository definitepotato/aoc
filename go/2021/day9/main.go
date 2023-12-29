package main

import (
	"fmt"
	"util"
)

func main() {
	input := util.ReadFile("input.txt")

	// This data structure allows for tracking of the row
	// location and each value of a given row. Makes it possible
	// to pull values from previous and next rows, relevant to current.
	grid := make([][]int, len(input))

	// Iterate each row in input.
	for i, row := range input {
		grid[i] = make([]int, len(row)) // Create a new slice for each row of input.
		for j, value := range row {
			grid[i][j] = util.Stoi(string(value)) // Store the current input row in the nested slice.
		}
	}

	var ans = 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			current := grid[y][x]

			// Up: grid[y-1][x]
			// Down: grid[y+1][x]
			// Right: grid[y][x+1]
			// Left: grid[y][x-1]

			// Start:
			// Skip all values with lower neighbors.
			// Staying within the boundaries of our grid size.
			if y > 0 && grid[y-1][x] <= current { // Look up if not in first row.
				continue
			}

			if y < len(input)-1 && grid[y+1][x] <= current { // Look down if not in last row.
				continue
			}

			if x > 0 && grid[y][x-1] <= current { // Look left if not in first column.
				continue
			}

			if x < len(input[0])-1 && grid[y][x+1] <= current { // Look right if not in last column.
				continue
			}
			// End.

			ans += current + 1
		}
	}

	fmt.Printf("Part 1: %d\n", ans)
}
