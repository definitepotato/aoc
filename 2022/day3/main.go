package main

import (
	"fmt"
	"util"
)

func priority(s string) int {
	switch s {
	case "a":
		return 1
	case "b":
		return 2
	case "c":
		return 3
	case "d":
		return 4
	case "e":
		return 5
	case "f":
		return 6
	case "g":
		return 7
	case "h":
		return 8
	case "i":
		return 9
	case "j":
		return 10
	case "k":
		return 11
	case "l":
		return 12
	case "m":
		return 13
	case "n":
		return 14
	case "o":
		return 15
	case "p":
		return 16
	case "q":
		return 17
	case "r":
		return 18
	case "s":
		return 19
	case "t":
		return 20
	case "u":
		return 21
	case "v":
		return 22
	case "w":
		return 23
	case "x":
		return 24
	case "y":
		return 25
	case "z":
		return 26
	case "A":
		return 27
	case "B":
		return 28
	case "C":
		return 29
	case "D":
		return 30
	case "E":
		return 31
	case "F":
		return 32
	case "G":
		return 33
	case "H":
		return 34
	case "I":
		return 35
	case "J":
		return 36
	case "K":
		return 37
	case "L":
		return 38
	case "M":
		return 39
	case "N":
		return 40
	case "O":
		return 41
	case "P":
		return 42
	case "Q":
		return 43
	case "R":
		return 44
	case "S":
		return 45
	case "T":
		return 46
	case "U":
		return 47
	case "V":
		return 48
	case "W":
		return 49
	case "X":
		return 50
	case "Y":
		return 51
	case "Z":
		return 52
	}

	return 0
}

func intersection(a, b string) string {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				return string(a[i])
			}
		}
	}

	return ""
}

func split(s string) []string {
	half := len(s) / 2
	left := s[:half]
	right := s[half : half+half]

	return []string{left, right}
}

func part1() {
	sum := 0
	rucksack := util.ReadFile("input.txt")

	for _, compartment := range rucksack {
		items := split(compartment)
		commonItem := intersection(items[0], items[1])
		sum += priority(commonItem)
	}

	fmt.Println(sum)
}

func intersectionGroup(a, b, c string) string {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				for k := 0; k < len(c); k++ {
					if c[k] == b[j] {
						return string(c[k])
					}
				}
			}
		}
	}

	return ""
}

func group(sack []string) map[int][]string {
	counter := 0
	max := 0
	groups := make(map[int][]string)

	for _, compartment := range sack {
		groups[counter] = append(groups[counter], compartment)
		max += 1

		if max == 3 {
			counter += 1
			max = 0
		}
	}

	return groups
}

func part2() {
	sum := 0
	rucksack := util.ReadFile("input.txt")

	groups := group(rucksack)

	for i := 1; i <= len(groups); i++ {
		badge := intersectionGroup(groups[i-1][0], groups[i-1][1], groups[i-1][2])
		sum += priority(badge)
	}

	fmt.Println(sum)
}

func main() {
	part1()
	part2()
}
