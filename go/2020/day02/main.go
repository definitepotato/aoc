package main

import (
	"fmt"
	"strings"
	"util"
)

func countLetters(l, p string) int {
	c := 0
	for _, letter := range p {
		if l == string(letter) {
			c++
		}
	}

	return c
}

func isValid(password string) bool {
	policy := strings.Split(password, ":")[0]
	passwd := strings.Split(password, ": ")[1]

	letter := strings.Split(policy, " ")[1]
	nRange := strings.Split(policy, " ")[0]
	nLow := strings.Split(nRange, "-")[0]
	nHigh := strings.Split(nRange, "-")[1]

	nLetters := countLetters(letter, passwd)

	if nLetters <= util.Stoi(nHigh) && nLetters >= util.Stoi(nLow) {
		return true
	}

	return false
}

func isValidPosition(password string) bool {
	policy := strings.Split(password, ":")[0]
	passwd := strings.Split(password, ": ")[1]

	letter := strings.Split(policy, " ")[1]
	nRange := strings.Split(policy, " ")[0]
	nLow := strings.Split(nRange, "-")[0]
	nHigh := strings.Split(nRange, "-")[1]

	if string(passwd[util.Stoi(nLow)-1]) == letter && string(passwd[util.Stoi(nHigh)-1]) == letter {
		return false
	}

	if string(passwd[util.Stoi(nLow)-1]) == letter || string(passwd[util.Stoi(nHigh)-1]) == letter {
		return true
	}

	return false
}

func part1() {
	passwords := util.ReadFile("input.txt")
	nTrue := 0

	for _, password := range passwords {
		if isValid(password) {
			nTrue++
		}
	}

	fmt.Println(nTrue)
}

func part2() {
	passwords := util.ReadFile("input.txt")
	nTrue := 0

	for _, password := range passwords {
		if isValidPosition(password) {
			nTrue++
		}
	}

	fmt.Println(nTrue)
}

func main() {
	part1()
	part2()
}
