package main

import (
	"fmt"
	"helpers"
)

func Hash(char rune, value int) int {
	value += int(char)  // inc current value by ascii code value of char
	value = value * 17  // set value to itself * 17
	value = value % 256 // set value to remainder of dividing itself by 256

	return value
}

func main() {
	input := helpers.ReadFile("input.txt")

	// Part 1.
	value := 0
	currentStepValue := 0
	for _, char := range input[0] {
		if char == 44 {
			value += currentStepValue
			currentStepValue = 0
			continue
		}

		currentStepValue = Hash(char, currentStepValue)
	}

	fmt.Println("Part 1: ", value)
}
