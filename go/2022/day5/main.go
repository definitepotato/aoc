package main

import (
	"fmt"
	"strings"
	"util"
)

type Stack map[int][]string

func (s Stack) IsEmpty(idx int) bool {
	return len(s[idx]) == 0
}

func (s Stack) Push(idx int, str string) {
	s[idx] = append(s[idx], str)
}

func (s Stack) Pop(idx int) (string, bool) {
	if s.IsEmpty(idx) {
		return "", false
	} else {
		index := len(s[idx]) - 1   // Get the index of the top most element.
		element := (s[idx])[index] // Index into the slice and obtain the element.
		s[idx] = (s[idx])[:index]  // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s Stack) TopCrates() []string {
	crates := []string{}

	for i := 0; i < len(s); i++ {
		crate, err := s.Pop(i)
		if !err {
			panic(err)
		}
		crates = append(crates, crate)
	}

	return crates
}

func ProdStack() Stack {
	stack := make(Stack)

	// Stack 1
	stack.Push(0, "F")
	stack.Push(0, "D")
	stack.Push(0, "B")
	stack.Push(0, "Z")
	stack.Push(0, "T")
	stack.Push(0, "J")
	stack.Push(0, "R")
	stack.Push(0, "N")

	// Stack 2
	stack.Push(1, "R")
	stack.Push(1, "S")
	stack.Push(1, "N")
	stack.Push(1, "J")
	stack.Push(1, "H")

	// Stack 3
	stack.Push(2, "C")
	stack.Push(2, "R")
	stack.Push(2, "N")
	stack.Push(2, "J")
	stack.Push(2, "G")
	stack.Push(2, "Z")
	stack.Push(2, "F")
	stack.Push(2, "Q")

	// Stack 4
	stack.Push(3, "F")
	stack.Push(3, "V")
	stack.Push(3, "N")
	stack.Push(3, "G")
	stack.Push(3, "R")
	stack.Push(3, "T")
	stack.Push(3, "Q")

	// Stack 5
	stack.Push(4, "L")
	stack.Push(4, "T")
	stack.Push(4, "Q")
	stack.Push(4, "F")

	// Stack 6
	stack.Push(5, "Q")
	stack.Push(5, "C")
	stack.Push(5, "W")
	stack.Push(5, "Z")
	stack.Push(5, "B")
	stack.Push(5, "R")
	stack.Push(5, "G")
	stack.Push(5, "N")

	// Stack 7
	stack.Push(6, "F")
	stack.Push(6, "C")
	stack.Push(6, "L")
	stack.Push(6, "S")
	stack.Push(6, "N")
	stack.Push(6, "H")
	stack.Push(6, "M")

	// Stack 8
	stack.Push(7, "D")
	stack.Push(7, "N")
	stack.Push(7, "Q")
	stack.Push(7, "M")
	stack.Push(7, "T")
	stack.Push(7, "J")

	// Stack 9
	stack.Push(8, "P")
	stack.Push(8, "G")
	stack.Push(8, "S")

	return stack
}

func DevStack() Stack {
	stack := make(Stack)

	// Stack 1
	stack.Push(0, "Z")
	stack.Push(0, "N")

	// Stack 2
	stack.Push(1, "M")
	stack.Push(1, "C")
	stack.Push(1, "D")

	// Stack 3
	stack.Push(2, "P")

	return stack
}

func part1() {
	stack := ProdStack()
	fmt.Println(stack)

	instructions := util.ReadFile("input.txt")
	for _, instruction := range instructions {
		move := strings.Split(instruction, " ")

		// move[1]: this many crates
		// move[3]: from here
		// move[5]: to here
		for i := 0; i < util.Stoi(move[1]); i++ {
			crate, err := stack.Pop(util.Stoi(move[3]) - 1)
			if !err {
				panic(err)
			}

			stack.Push(util.Stoi(move[5])-1, crate)
		}
	}

	fmt.Println(stack.TopCrates())
}

func part2() {
	stack := ProdStack()
	fmt.Println(stack)

	instructions := util.ReadFile("input.txt")
	for _, instruction := range instructions {
		move := strings.Split(instruction, " ")

		// move[1]: this many crates
		// move[3]: from here
		// move[5]: to here
		if util.Stoi(move[1]) == 1 {
			crate, err := stack.Pop(util.Stoi(move[3]) - 1)
			if !err {
				panic(err)
			}
			stack.Push(util.Stoi(move[5])-1, crate)
		}

		// Keep the order when moving more than 1 crate.
		if util.Stoi(move[1]) > 1 {
			var holding []string
			for i := 0; i < util.Stoi(move[1]); i++ {
				crate, err := stack.Pop(util.Stoi(move[3]) - 1)
				if !err {
					panic(err)
				}
				holding = append(holding, crate)
			}

			holdingReversed := util.Reverse(holding) // Reverse to keep original order.
			for i := 0; i < len(holdingReversed); i++ {
				stack.Push(util.Stoi(move[5])-1, holdingReversed[i])
			}
		}
	}

	fmt.Println(stack.TopCrates())
}

func main() {
	part1()
	part2()
}
