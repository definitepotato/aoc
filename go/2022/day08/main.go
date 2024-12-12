package main

import (
	"fmt"
	"strconv"
	"strings"
	"util"
)

// isVisible returns true if the given tree is visible from outside the forest
func isVisible(row []int, n int) bool {
	for _, v := range row {
		if v >= n {
			return false
		}
	}
	return true
}

// numVisible returns the number of trees visible from the given tree
func numVisible(row []int, n int, reverse bool) int {
	count := 0
	if reverse {
		for i := len(row) - 1; i >= 0; i-- {
			count++
			if row[i] >= n {
				return count
			}
		}
	} else {
		for _, v := range row {
			count++
			if v >= n {
				return count
			}
		}
	}

	return count
}

func main() {
	input := util.ReadFile("input.txt")

	grid := make([][]int, len(input))
	for i, v := range input {
		t := strings.Split(v, "")
		grid[i] = make([]int, 0)
		for _, s := range t {
			d, _ := strconv.Atoi(s)
			grid[i] = append(grid[i], d)
		}
	}

	cols := make([][]int, len(grid[0]))
	for i := 0; i < len(grid[0]); i++ {
		cols[i] = make([]int, 0)
		for j := 0; j < len(grid); j++ {
			cols[i] = append(cols[i], grid[j][i])
		}
	}

	visible := 0
	max := 0
	for ri, row := range grid {
		for ti, tree := range row {
			if ri == 0 || ti == 0 || ri == len(grid)-1 || ti == len(row)-1 {
				// top, bottom, left and right are always visible
				visible++
			} else if isVisible(row[ti+1:], tree) {
				// right
				visible++
			} else if isVisible(row[:ti], tree) {
				// left
				visible++
			} else if isVisible(cols[ti][ri+1:], tree) {
				// below
				visible++
			} else if isVisible(cols[ti][:ri], tree) {
				// above
				visible++
			}

			right := numVisible(row[ti+1:], tree, false)      // right
			left := numVisible(row[:ti], tree, true)          // left
			below := numVisible(cols[ti][ri+1:], tree, false) // below
			above := numVisible(cols[ti][:ri], tree, true)    // above
			score := right * left * below * above

			if score > max {
				max = score
			}
		}
	}
	fmt.Println("part1", visible)
	fmt.Println("part2", max)
}
