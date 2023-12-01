package main

import (
	"fmt"
	"util"
)

type P struct {
	x1, y1 int
	x2, y2 int
}

type C struct {
	x, y int
}

func getAllPointsPart1(points []P) []C {
	allPoints := []C{}
	for _, p := range points {
		if p.y1 == p.y2 && p.x1 < p.x2 {
			for a := 0; a <= p.x2-p.x1; a++ {
				allPoints = append(allPoints, C{p.x1 + a, p.y1})
			}
		}

		if p.y1 == p.y2 && p.x1 > p.x2 {
			for a := 0; a <= p.x1-p.x2; a++ {
				allPoints = append(allPoints, C{p.x1 - a, p.y1})
			}
		}

		if p.x1 == p.x2 && p.y1 < p.y2 {
			for a := 0; a <= p.y2-p.y1; a++ {
				allPoints = append(allPoints, C{p.x1, p.y1 + a})
			}
		}

		if p.x1 == p.x2 && p.y1 > p.y2 {
			for a := 0; a <= p.y1-p.y2; a++ {
				allPoints = append(allPoints, C{p.x1, p.y1 - a})
			}
		}
	}

	return allPoints
}

func getAllPointsPart2(points []P) []C {
	allPoints := []C{}
	for _, p := range points {
		if p.y1 == p.y2 && p.x1 < p.x2 {
			for a := 0; a <= p.x2-p.x1; a++ {
				allPoints = append(allPoints, C{p.x1 + a, p.y1})
			}
		}
		if p.y1 == p.y2 && p.x1 > p.x2 {
			for a := 0; a <= p.x1-p.x2; a++ {
				allPoints = append(allPoints, C{p.x1 - a, p.y1})
			}
		}

		if p.x1 == p.x2 && p.y1 < p.y2 {
			for a := 0; a <= p.y2-p.y1; a++ {
				allPoints = append(allPoints, C{p.x1, p.y1 + a})
			}
		}

		if p.x1 == p.x2 && p.y1 > p.y2 {
			for a := 0; a <= p.y1-p.y2; a++ {
				allPoints = append(allPoints, C{p.x1, p.y1 - a})
			}
		}

		if p.x1 != p.x2 && p.y1 != p.y2 && p.x2 > p.x1 && p.y2 > p.y1 {
			for a := 0; a <= p.y2-p.y1; a++ {
				allPoints = append(allPoints, C{p.x1 + a, p.y1 + a})
			}
		}

		if p.x1 != p.x2 && p.y1 != p.y2 && p.x2 < p.x1 && p.y2 > p.y1 {
			for a := 0; a <= p.y2-p.y1; a++ {
				allPoints = append(allPoints, C{p.x1 - a, p.y1 + a})
			}
		}

		if p.x1 != p.x2 && p.y1 != p.y2 && p.x2 > p.x1 && p.y2 < p.y1 {
			for a := 0; a <= p.y1-p.y2; a++ {
				allPoints = append(allPoints, C{p.x1 + a, p.y1 - a})
			}
		}

		if p.x1 != p.x2 && p.y1 != p.y2 && p.x2 < p.x1 && p.y2 < p.y1 {
			for a := 0; a <= p.y1-p.y2; a++ {
				allPoints = append(allPoints, C{p.x1 - a, p.y1 - a})
			}
		}
	}
	return allPoints
}

func reducePoints(coordinates []C) map[C]bool {
	reduced := make(map[C]bool)
	for i := 0; i < len(coordinates)-1; i++ { // get an element of []C
		for j := i + 1; j < len(coordinates); j++ { // iterate every other element []C
			if coordinates[i] == coordinates[j] {
				reduced[coordinates[i]] = true
			}
		}
	}
	return reduced
}

func main() {
	input := util.ReadFile("input.txt")

	/*
				Difference between nil & empty slices
				=====================================

				If we think of a slice like this:
				[pointer] [length] [capacity]

				then:

				nil slice:   [nil][0][0]
				empty slice: [addr][0][0] // points to an address

		    - nil slice:
		    They’re useful when you want to represent a slice that doesn’t exist, such as when an exception occurs in a function that returns a slice.

		    // Create a nil slice of integers.
		    var slice []int

		    - empty slice:
		    Empty slices are useful when you want to represent an empty collection, such as when a database query returns zero results.

		    // Use make to create an empty slice of integers.
		    slice := make([]int, 0)

		    // Use a slice literal to create an empty slice of integers.
		    slice := []int{}

		    Regardless of whether you’re using a nil slice or an empty slice, the built-in functions append, len, and cap work the same.
	*/

	var points = make([]P, len(input))
	// point := []P{}	// This will return index out of range [0] with length 0 (see above)
	for i, p := range input {
		fmt.Sscanf(p, "%d,%d -> %d,%d",
			&points[i].x1,
			&points[i].y1,
			&points[i].x2,
			&points[i].y2,
		)
	}

	part1 := getAllPointsPart1(points)
	part2 := getAllPointsPart2(points)

	fmt.Printf("Part 1: %d\n", len(reducePoints(part1)))
	fmt.Printf("Part 2: %d\n", len(reducePoints(part2)))
}
