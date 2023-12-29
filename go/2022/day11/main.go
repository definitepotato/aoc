package main

import (
	"fmt"
	"sort"
)

type Monkey struct {
	Items                   []int
	Operation               func(int) int
	TestDivisibleBy         int
	TrueMonkey, FalseMonkey int
}

// faster to manually type this than write a parser (and potentially debug)
func ProdInput() []Monkey {
	return []Monkey{
		{
			Items: []int{77, 69, 76, 77, 50, 58},
			Operation: func(old int) int {
				return old * 11
			},
			TestDivisibleBy: 5,
			TrueMonkey:      1,
			FalseMonkey:     5,
		},
		{
			Items: []int{75, 70, 82, 83, 96, 64, 62},
			Operation: func(old int) int {
				return old + 8
			},
			TestDivisibleBy: 17,
			TrueMonkey:      5,
			FalseMonkey:     6,
		},
		{
			Items: []int{53},
			Operation: func(old int) int {
				return old * 3
			},
			TestDivisibleBy: 2,
			TrueMonkey:      0,
			FalseMonkey:     7,
		},
		{
			Items: []int{85, 64, 93, 64, 99},
			Operation: func(old int) int {
				return old + 4
			},
			TestDivisibleBy: 7,
			TrueMonkey:      7,
			FalseMonkey:     2,
		},
		{
			Items: []int{61, 92, 71},
			Operation: func(old int) int {
				return old * old
			},
			TestDivisibleBy: 3,
			TrueMonkey:      2,
			FalseMonkey:     3,
		},
		{
			Items: []int{79, 73, 50, 90},
			Operation: func(old int) int {
				return old + 2
			},
			TestDivisibleBy: 11,
			TrueMonkey:      4,
			FalseMonkey:     6,
		},
		{
			Items: []int{50, 89},
			Operation: func(old int) int {
				return old + 3
			},
			TestDivisibleBy: 13,
			TrueMonkey:      4,
			FalseMonkey:     3,
		},
		{
			Items: []int{83, 56, 64, 58, 93, 91, 56, 65},
			Operation: func(old int) int {
				return old + 5
			},
			TestDivisibleBy: 19,
			TrueMonkey:      1,
			FalseMonkey:     0,
		},
	}
}

func DevInput() []Monkey {
	return []Monkey{
		{
			Items: []int{79, 98},
			Operation: func(num int) int {
				return num * 19
			},
			TestDivisibleBy: 23,
			TrueMonkey:      2,
			FalseMonkey:     3,
		},
		{
			Items: []int{54, 65, 75, 74},
			Operation: func(num int) int {
				return num + 6
			},
			TestDivisibleBy: 19,
			TrueMonkey:      2,
			FalseMonkey:     0,
		},
		{
			Items: []int{79, 60, 97},
			Operation: func(num int) int {
				return num * num
			},
			TestDivisibleBy: 13,
			TrueMonkey:      1,
			FalseMonkey:     3,
		},
		{
			Items: []int{74},
			Operation: func(num int) int {
				return num + 3
			},
			TestDivisibleBy: 17,
			TrueMonkey:      0,
			FalseMonkey:     1,
		},
	}
}

// oh my god i figured out a math-y remainder theorem-y thing myself!
func part2() {
	monkeys := ProdInput()

	// the worry levels will always increase now that they're not being divided
	// by 3, and we care about remainders because that's what all the tests are
	// BUT we can't just mod by any monkey's testBy number, because they're all
	// throwing the items around,
	// so find a shared common denominator that can be used to keep the numbers
	// under overflow
	bigMod := 1
	for _, m := range monkeys {
		bigMod *= m.TestDivisibleBy
	}

	inspectedCounts := make([]int, len(monkeys))
	for round := 0; round < 10000; round++ {

		for i, monkey := range monkeys {
			for _, item := range monkey.Items {
				newItemVal := monkey.Operation(item)
				newItemVal %= bigMod

				if newItemVal%monkey.TestDivisibleBy == 0 {
					monkeys[monkey.TrueMonkey].Items = append(
						monkeys[monkey.TrueMonkey].Items, newItemVal)
				} else {
					monkeys[monkey.FalseMonkey].Items = append(
						monkeys[monkey.FalseMonkey].Items, newItemVal)
				}

			}
			inspectedCounts[i] += len(monkey.Items)

			// empty out this monkey's items
			monkeys[i].Items = []int{}
		}
	}

	sort.Ints(inspectedCounts)
	counts := inspectedCounts[len(inspectedCounts)-1] * inspectedCounts[len(inspectedCounts)-2]
	fmt.Println(counts)
}

func part1() {
	monkeys := ProdInput()

	inspectedCounts := make([]int, len(monkeys))
	for round := 0; round < 20; round++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.Items {
				newItemVal := monkey.Operation(item) / 3

				if newItemVal%monkey.TestDivisibleBy == 0 {
					monkeys[monkey.TrueMonkey].Items = append(
						monkeys[monkey.TrueMonkey].Items, newItemVal)
				} else {
					monkeys[monkey.FalseMonkey].Items = append(
						monkeys[monkey.FalseMonkey].Items, newItemVal)
				}

			}
			inspectedCounts[i] += len(monkey.Items)

			// empty out this monkey's items
			monkeys[i].Items = []int{}
		}
	}

	sort.Ints(inspectedCounts)
	counts := inspectedCounts[len(inspectedCounts)-1] * inspectedCounts[len(inspectedCounts)-2]
	fmt.Println(counts)
}

func main() {
	part1()
	part2()
}
