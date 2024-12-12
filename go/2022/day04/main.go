package main

import (
	"fmt"
	"strings"
	"util"
)

func assignment(s string) []string {
	return strings.Split(s, ",")
}

func overlapping(s []string) bool {
	left := strings.Split(s[0], "-")
	right := strings.Split(s[1], "-")

	if util.Stoi(left[0]) <= util.Stoi(right[0]) && util.Stoi(left[1]) >= util.Stoi(right[1]) {
		return true
	}

	if util.Stoi(right[0]) <= util.Stoi(left[0]) && util.Stoi(right[1]) >= util.Stoi(left[1]) {
		return true
	}

	return false
}

func anyOverlap(s []string) bool {
	left := strings.Split(s[0], "-")
	right := strings.Split(s[1], "-")

	if util.Stoi(left[1]) >= util.Stoi(right[0]) && util.Stoi(left[0]) <= util.Stoi(right[1]) {
		return true
	}

	return false
}

func part1() {
	overlappingAssignments := 0

	assignments := util.ReadFile("input.txt")
	for _, a := range assignments {
		if overlapping((assignment(a))) {
			overlappingAssignments += 1
		}
	}

	fmt.Println(overlappingAssignments)
}

func part2() {
	overlappingAssignments := 0

	assignments := util.ReadFile("input.txt")
	for _, a := range assignments {
		if anyOverlap((assignment(a))) {
			overlappingAssignments += 1
		}
	}

	fmt.Println(overlappingAssignments)
}

func main() {
	part1()
	part2()
}
