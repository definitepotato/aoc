package main

import (
	"fmt"
	"helpers"
	"math"
	"regexp"
	"strconv"
	"time"
)

const (
	GridX = 101
	GridY = 103
)

type Robot struct {
	PosX int
	PosY int
	VelX int
	VelY int
}

func absolute(a int) int {
	return int(math.Abs(float64(a)))
}

func atoiNoError(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func maybeTree(r []Robot) (bool, [GridY][GridX]string) {
	var grid [GridY][GridX]string

	for y := 0; y < GridY; y++ {
		for x := 0; x < GridX; x++ {
			grid[y][x] = "."
		}
	}

	for _, robot := range r {
		grid[robot.PosY][robot.PosX] = "A"
	}

	count := 0
	for y := 0; y < GridY; y++ {
		for x := 0; x < GridX; x++ {
			if grid[y][x] == "A" {
				count++
			}
		}

		if count > 30 {
			return true, grid
		}

		count = 0
	}

	return false, [GridY][GridX]string{}
}

func quadrant(x, y int) int {
	if x >= 0 && x < GridX/2 { // left half (Q1 || Q2)
		if y >= 0 && y < GridY/2 { // (Q1)
			return 1
		}

		if y > GridY/2 && y <= GridY { // (Q2)
			return 2
		}
	}

	if x > GridX/2 && x <= GridX { // right half (Q3 || Q4)
		if y >= 0 && y < GridY/2 { // (Q3)
			return 3
		}

		if y > GridY/2 && y <= GridY { // (Q4)
			return 4
		}
	}

	// line void
	return 0
}

func countRobotsInQuadrants(r []Robot) map[int]int {
	count := make(map[int]int, 0)

	for _, robot := range r {
		q := quadrant(robot.PosX, robot.PosY)
		count[q] += 1
	}

	return count
}

func moveAllRobots(r []Robot) []Robot {
	var newPositions []Robot
	for _, robot := range r {
		r := Robot{
			PosX: (robot.PosX + robot.VelX + GridX) % GridX,
			PosY: (robot.PosY + robot.VelY + GridY) % GridY,
			VelX: robot.VelX,
			VelY: robot.VelY,
		}
		newPositions = append(newPositions, r)
	}

	return newPositions
}

func solve(m map[int]int) int {
	return m[1] * m[3] * m[2] * m[4]
}

func main() {
	pattern := regexp.MustCompile(`p=(-?\d+),(-?\d+)\s+v=(-?\d+),(-?\d+)`)

	input := helpers.ReadFile("../input.txt")
	var robots []Robot

	for _, line := range input {
		match := pattern.FindStringSubmatch(line)

		robot := Robot{
			PosX: atoiNoError(match[1]),
			PosY: atoiNoError(match[2]),
			VelX: atoiNoError(match[3]),
			VelY: atoiNoError(match[4]),
		}

		robots = append(robots, robot)
	}

	for i := 0; i < 100; i++ {
		robots = moveAllRobots(robots)
	}

	quadrantCount := countRobotsInQuadrants(robots)
	result := solve(quadrantCount)
	fmt.Printf("Part 1: %d\n", result)

	iter := 0
	for {
		robots = moveAllRobots(robots)
		iter += 1
		res, grid := maybeTree(robots)
		if res {
			for _, g := range grid {
				fmt.Println(g)
			}
			fmt.Printf("Part 2: %d\n", iter)
		}
		time.Sleep(5 * time.Millisecond)
	}
}
