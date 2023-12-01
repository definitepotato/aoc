package main

import (
	"fmt"
	"util"
)

func outcome(play string, part int) int {
	switch play {
	case "A Y":
		if part == 1 {
			return 8
		} else {
			return 4
		}

	case "A X":
		if part == 1 {
			return 4
		} else {
			return 3
		}

	case "A Z":
		if part == 1 {
			return 3
		} else {
			return 8
		}

	case "B Y":
		if part == 1 {
			return 5
		} else {
			return 5
		}

	case "B X":
		if part == 1 {
			return 1
		} else {
			return 1
		}

	case "B Z":
		if part == 1 {
			return 9
		} else {
			return 9
		}

	case "C Y":
		if part == 1 {
			return 2
		} else {
			return 6
		}

	case "C X":
		if part == 1 {
			return 7
		} else {
			return 2
		}

	case "C Z":
		if part == 1 {
			return 6
		} else {
			return 7
		}
	}

	return 0
}

func main() {
	strategy := util.ReadFile("input.txt")

	// Part 1
	score1 := 0
	for _, play := range strategy {
		score1 += outcome(play, 1)
	}
	fmt.Println(score1)

	// Part 2
	score2 := 0
	for _, play := range strategy {
		score2 += outcome(play, 2)
	}
	fmt.Println(score2)
}
