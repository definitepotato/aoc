package main

import (
	"fmt"
	"helpers"
)

func findStartingLoc(grid []string) [2]int {
	for y := range grid {
		for x := range y {
			switch grid[y][x] {
			case '^', '>', 'v', '<':
				return [2]int{int(y), int(x)}
			}
		}
	}

	return [2]int{0, 0}
}

func main() {
	path := helpers.ReadFile("../input.txt")

	grid := []string{}

	for _, v := range path {
		grid = append(grid, v)
	}

	visited := make(map[string]bool, 0)
	startingLoc := findStartingLoc(grid)
	y := startingLoc[0]
	x := startingLoc[1]
	dir := grid[y][x]
	currentCoord := fmt.Sprintf("%d,%d", y, x)
	visited[currentCoord] = true

	for y >= 0 && y < len(grid)-1 && x >= 0 && x < len(grid[x])-1 {
		if dir == '^' {
			if grid[y-1][x] == '#' {
				// fmt.Printf("Change dir: '>'\n")
				dir = '>'
				continue
			}
			y--
			currentCoord := fmt.Sprintf("%d,%d", y, x)
			visited[currentCoord] = true
		}

		if dir == '>' {
			if grid[y][x+1] == '#' {
				// fmt.Printf("Change dir: 'v'\n")
				dir = 'v'
				continue
			}
			x++
			currentCoord := fmt.Sprintf("%d,%d", y, x)
			visited[currentCoord] = true
		}

		if dir == 'v' {
			if grid[y+1][x] == '#' {
				// fmt.Printf("Change dir: '<'\n")
				dir = '<'
				continue
			}
			y++
			currentCoord := fmt.Sprintf("%d,%d", y, x)
			visited[currentCoord] = true
		}

		if dir == '<' {
			if grid[y][x-1] == '#' {
				// fmt.Printf("Change dir: '^'\n")
				dir = '^'
				continue
			}
			x--
			currentCoord := fmt.Sprintf("%d,%d", y, x)
			visited[currentCoord] = true
		}

		// fmt.Printf("%d,%d: %c => %d\n", y, x, grid[y][x], len(visited))
	}

	fmt.Printf("Part 1: %d\n", len(visited))
}
