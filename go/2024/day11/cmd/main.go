package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Stones      = "0 1 10 99 999" // [1 blink] 1 2024 1 0 9 9 2021976
	TestInput   = "125 17"        // [6 blinks] 2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2
	PuzzleInput = "510613 358 84 40702 4373582 2 0 1584"
)

func sliceStringToInt(in []string) []int {
	var out []int

	for _, v := range in {
		i, _ := strconv.Atoi(v)
		out = append(out, i)
	}

	return out
}

func insert(in []int, idx int, val int) []int {
	if idx < 0 || idx > len(in) {
		fmt.Println("out of range")
		return in
	}

	// insert `val` in new slice at `idx`
	out := append(in[:idx], append([]int{val}, in[idx:]...)...)
	return out
}

func countStones(in string) int {
	inputString := strings.Split(in, " ")
	inputSlice := sliceStringToInt(inputString)

	for c := 0; c < 25; c++ {
		var newStones []int
		for _, stone := range inputSlice {
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if len(strconv.Itoa(stone))%2 == 0 {
				ss := strconv.Itoa(stone) // convert to string for splitting
				left := ss[:len(ss)/2]    // left half of number
				right := ss[len(ss)/2:]   // right half of number

				stoneLeft, _ := strconv.Atoi(left)   // convert back to number
				stoneRight, _ := strconv.Atoi(right) // convert back to number

				newStones = append(newStones, stoneLeft)
				newStones = append(newStones, stoneRight)
			} else {
				newStones = append(newStones, stone*2024)
			}
		}

		inputSlice = newStones
	}

	return len(inputSlice)
}

// Part 2 starts here.
// Not my code, I couldn't figure this out but wanted to learn so
// I copied most of the solve from: https://kyle.so/writing/aoc-2024
func blinkOnce(stone int) []int {
	if stone == 0 {
		return []int{1}
	}

	// convert to string to check digits
	s := strconv.Itoa(stone)

	if len(s)%2 == 0 {
		left := s[:len(s)/2]  // left half of number
		right := s[len(s)/2:] // right half of number

		stoneLeft, _ := strconv.Atoi(left)   // convert back to number
		stoneRight, _ := strconv.Atoi(right) // convert back to number

		return []int{stoneLeft, stoneRight}
	}

	return []int{stone * 2024}
}

func blinkAllStones(stoneCount map[int]int) map[int]int {
	nc := make(map[int]int)

	for s, c := range stoneCount {
		// get new stones from transformation
		ns := blinkOnce(s)
		// add to counts, multiplied by how many of oritinal stones we had
		for _, n := range ns {
			nc[n] += c
		}
	}

	return nc
}

func sumStones(stoneCount map[int]int) int {
	total := 0
	for _, c := range stoneCount {
		total += c
	}
	return total
}

func countStonesFaster(in string) int {
	inputString := strings.Split(in, " ")
	inputSlice := sliceStringToInt(inputString)

	// count freq of each stone val
	freq := make(map[int]int)
	for _, stone := range inputSlice {
		freq[stone]++
	}

	for c := 0; c < 75; c++ {
		freq = blinkAllStones(freq)
	}

	return sumStones(freq)
}

func main() {
	fmt.Printf("Part 1: %d\n", countStones(PuzzleInput))
	fmt.Printf("Part 2: %d\n", countStonesFaster(PuzzleInput))
}
