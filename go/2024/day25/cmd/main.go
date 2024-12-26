package main

import (
	"fmt"
	"helpers"
)

func main() {
	file := helpers.ReadFile("../input.txt")

	locks := make(map[int][5]int, 0)
	keys := make(map[int][5]int, 0)

	current := 0
	val := [5]int{}
	lock, key := false, false
	for idx, row := range file {
		if idx%8 == 0 {
			if row == "#####" {
				lock = true
			}

			if row == "....." {
				key = true
			}
		}

		if row == "" {
			if lock {
				locks[current] = val
			}

			if key {
				keys[current] = val
			}

			val = [5]int{}
			current++
			lock = false
			key = false
			continue
		}

		for i := 0; i < len(row); i++ {
			if row[i] == '#' {
				val[i] += 1
			}
		}
	}

	ansPartOne := 0
	for _, key := range keys {
		for _, lock := range locks {
			unlocked := true
			// tmp := [5]int{}
			for i := 0; i < len(key); i++ {
				// tmp[i] = key[i] + lock[i]
				if key[i]+lock[i] > 7 {
					unlocked = false
					continue
				}
			}

			if unlocked {
				ansPartOne++
			}
			// fmt.Printf("Key:%d Unlocks:%d [%v]:[%d]\n", key, lock, unlocked, tmp)
		}
	}

	fmt.Printf("Part 1: %d\n", ansPartOne)
}
