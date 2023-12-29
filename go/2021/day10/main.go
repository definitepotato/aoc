package main

import (
	"fmt"
	"sort"
	"sync"
	"util"
)

type Stack struct {
	item []string
	lock sync.RWMutex
}

// Check if stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.item) == 0
}

// Push a new value onto the stack.
func (s *Stack) Push(str string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.item = append(s.item, str)
}

// Remove and return top element of stack.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		s.lock.Lock()
		defer s.lock.Unlock()
		index := len(s.item) - 1 // Get the index of the last element
		element := s.item[index] // Index into the slice and obtain the element
		s.item = s.item[:index]  // Remove it from the stack by slicing it off
		return element, true
	}
}

func (s *Stack) Last() string {
	if s.IsEmpty() {
		return ""
	} else {
		s.lock.Lock()
		defer s.lock.Unlock()
		index := len(s.item) - 1 // Get the index of the last element
		element := s.item[index] // Index into the slice and obtain the element
		return element
	}
}

// Handle incomplete lines.
func FixSyntax(s string) []string {
	var last string
	expecting := &Stack{}

	for i := 0; i < len(s); i++ { // elements of s

		// Skip corrupted lines.
		// Remove completed brackets.
		// Only incompleted brackets should remain.
		switch string(s[i]) {
		case "{", "[", "(", "<":
			expecting.Push(string(s[i]))
		case "}":
			last = expecting.Last()
			if last == "{" {
				_, _ = expecting.Pop()
			} else {
				continue
			}
		case "]":
			last = expecting.Last()
			if last == "[" {
				_, _ = expecting.Pop()
			} else {
				continue
			}
		case ")":
			last = expecting.Last()
			if last == "(" {
				_, _ = expecting.Pop()
			} else {
				continue
			}
		case ">":
			last = expecting.Last()
			if last == "<" {
				_, _ = expecting.Pop()
			} else {
				continue
			}
		}
	}

	// For each incomplete open bracket collected, append it's closing bracket.
	final := []string{}
	for _, v := range expecting.item {
		switch v {
		case "{":
			final = append(final, "}")
		case "[":
			final = append(final, "]")
		case "(":
			final = append(final, ")")
		case "<":
			final = append(final, ">")
		}
	}

	return util.Reverse(final)
}

// Handle corrupted lines.
func CheckSyntax(s string) int {
	var last string
	expecting := &Stack{}

	for i := 0; i < len(s); i++ { // elements of s

		// Skip completed or incompleted lines.
		// Corrupted lines have bad closing bracket.
		// Return point value when bad closing bracket is found.
		switch string(s[i]) {
		case "{", "[", "(", "<":
			expecting.Push(string(s[i]))
		case "}":
			last = expecting.Last()
			if last == "{" {
				_, _ = expecting.Pop()
				continue
			} else {
				return 1197
			}
		case "]":
			last = expecting.Last()
			if last == "[" {
				_, _ = expecting.Pop()
				continue
			} else {
				return 57
			}
		case ")":
			last = expecting.Last()
			if last == "(" {
				_, _ = expecting.Pop()
				continue
			} else {
				return 3
			}
		case ">":
			last = expecting.Last()
			if last == "<" {
				_, _ = expecting.Pop()
				continue
			} else {
				return 25137
			}
		}
	}
	return 0
}

func Part1(input []string) int {
	var points int = 0

	for i := 0; i < len(input); i++ { // row
		points += CheckSyntax(input[i])
	}

	return points
}

func Part2(input []string) int {
	var points int = 0
	var scores []int

	for i := 0; i < len(input); i++ { // row
		if res := CheckSyntax(input[i]); res == 0 {
			res := FixSyntax(input[i])
			for _, v := range res {
				points *= 5

				if v == "}" {
					points += 3
				}

				if v == "]" {
					points += 2
				}

				if v == ")" {
					points += 1
				}

				if v == ">" {
					points += 4
				}
			}
			scores = append(scores, points)
			points = 0
		}
	}

	sort.Ints(scores)
	z := (len(scores) - 1) / 2 // Get middle score
	return scores[z]
}

func main() {
	input := util.ReadFile("input.txt")

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
