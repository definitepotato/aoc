package main

import (
	"fmt"
	"helpers"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) Push(point string) {
	*s = append(*s, point)
}

func (s *Stack) Pop() string {
	if len(*s) == 0 {
		return "-1"
	}

	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func sumPoints(points map[string]int) int {
	sum := 0
	for _, p := range points {
		sum += p
	}

	return sum
}

func startingPoints(grid []string) map[string]int {
	points := make(map[string]int, 0)

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '0' {
				pointY := strconv.Itoa(y)
				pointX := strconv.Itoa(x)

				point := fmt.Sprintf("%s,%s", pointY, pointX)
				points[point] = 0
			}
		}
	}

	return points
}

func traverse(startPos string, grid []string, points map[string]int, points2 map[string]int) {
	var stack Stack
	stack.Push(startPos)

	visited := make(map[string]bool, 0)

	for !stack.IsEmpty() {
		curPos := stack.Pop()
		pointY, _ := strconv.Atoi(strings.Split(curPos, ",")[0])
		pointX, _ := strconv.Atoi(strings.Split(curPos, ",")[1])
		curPosN, _ := strconv.Atoi(string(grid[pointY][pointX]))

		right := pointX < len(grid[0])-1
		left := pointX > 0
		up := pointY > 0
		down := pointY < len(grid)-1

		if right {
			nextPos := fmt.Sprintf("%s,%s", strconv.Itoa(pointY), strconv.Itoa(pointX+1))
			nextPosN, _ := strconv.Atoi(string(grid[pointY][pointX+1]))

			if nextPosN == 9 && curPosN == 8 {
				points2[startPos] += 1
				if !visited[nextPos] {
					points[startPos] += 1
					visited[nextPos] = true
				}
			}

			if nextPosN == curPosN+1 && curPosN != 8 {
				stack.Push(nextPos)
			}
		}

		if left {
			nextPos := fmt.Sprintf("%s,%s", strconv.Itoa(pointY), strconv.Itoa(pointX-1))
			nextPosN, _ := strconv.Atoi(string(grid[pointY][pointX-1]))

			if nextPosN == 9 && curPosN == 8 {
				points2[startPos] += 1
				if !visited[nextPos] {
					points[startPos] += 1
					visited[nextPos] = true
				}
			}

			if nextPosN == curPosN+1 && curPosN != 8 {
				stack.Push(nextPos)
			}
		}

		if up {
			nextPos := fmt.Sprintf("%s,%s", strconv.Itoa(pointY-1), strconv.Itoa(pointX))
			nextPosN, _ := strconv.Atoi(string(grid[pointY-1][pointX]))

			if nextPosN == 9 && curPosN == 8 {
				points2[startPos] += 1
				if !visited[nextPos] {
					points[startPos] += 1
					visited[nextPos] = true
				}
			}

			if nextPosN == curPosN+1 && curPosN != 8 {
				stack.Push(nextPos)
			}
		}

		if down {
			nextPos := fmt.Sprintf("%s,%s", strconv.Itoa(pointY+1), strconv.Itoa(pointX))
			nextPosN, _ := strconv.Atoi(string(grid[pointY+1][pointX]))

			if nextPosN == 9 && curPosN == 8 {
				points2[startPos] += 1
				if !visited[nextPos] {
					points[startPos] += 1
					visited[nextPos] = true
				}
			}

			if nextPosN == curPosN+1 && curPosN != 8 {
				stack.Push(nextPos)
			}
		}
	}
}

func main() {
	input := helpers.ReadFile("../input.txt")
	var grid []string

	for _, line := range input {
		grid = append(grid, line)
	}

	points := startingPoints(grid)
	points2 := startingPoints(grid)

	var trailheadStack Stack
	for point := range points {
		trailheadStack.Push(point)
	}

	for !trailheadStack.IsEmpty() {
		traverse(trailheadStack.Pop(), grid, points, points2)
	}

	fmt.Printf("Part 1: %d\n", sumPoints(points))
	fmt.Printf("Part 2: %d\n", sumPoints(points2))
}
