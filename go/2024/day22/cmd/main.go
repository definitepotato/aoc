package main

import (
	"fmt"
	"helpers"
	"strconv"
)

func Mix(a, b int) int {
	return a ^ b
}

func Prune(a int) int {
	return a % 16777216
}

func NextSecret(n int) int {
	// Round 1
	n = Mix(n*64, n)
	n = Prune(n)

	// Round 2
	n = Mix(n/32, n)
	n = Prune(n)

	// Round 3
	n = Mix(n*2048, n)
	n = Prune(n)

	return n
}

func main() {
	// secret := 123
	// for i := 0; i < 10; i++ {
	// 	secret = NextSecret(secret)
	// 	fmt.Println(secret)
	// }

	// test_input := []string{"1", "10", "100", "2024"}

	file := helpers.ReadFile("../input.txt")

	puzzle_input := []string{}
	for _, line := range file {
		puzzle_input = append(puzzle_input, line)
	}

	fmt.Println(len(puzzle_input))

	ans := 0
	for _, input := range puzzle_input {
		secret, _ := strconv.Atoi(input)

		for i := 0; i < 2000; i++ {
			secret = NextSecret(secret)
		}

		ans += secret
	}

	fmt.Printf("Part 1: %d\n", ans)
}
