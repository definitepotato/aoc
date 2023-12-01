package main

import (
	"fmt"
	"helpers"
	"strconv"
)

func stoi(s string) (int, error) {
	r, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return r, nil
}

func IsNum(x string) (int, bool) {
	var num int

	if num, err := stoi(x); err != nil {
		return num, false
	}

	return num, true
}

func SeekFromLeft(s string) string {
	for i := 0; i < len(s); i++ {
		c := string(s[i])

		if _, n := IsNum(c); n {
			return c
		}
	}

	return ""
}

func SeekFromRight(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		c := string(s[i])

		if _, n := IsNum(c); n {
			return c
		}
	}

	return ""
}

func SeekWordsAsNum(s string) (string, string) {
	n := []string{}

	for i := 0; i < len(s); i++ {
		c := string(s[i])

		if _, ok := IsNum(c); ok {
			num, _ := stoi(c)
			n = append(n, fmt.Sprint(num))
			continue
		}

		if s[i] == 'o' {
			if i+3 <= len(s) {
				if s[i:i+3] == "one" {
					n = append(n, "1")
					continue
				}
			}
		}

		if s[i] == 't' {
			if i+3 <= len(s) {
				if s[i:i+3] == "two" {
					n = append(n, "2")
					continue
				}
			}

			if i+5 <= len(s) {
				if s[i:i+5] == "three" {
					n = append(n, "3")
					continue
				}
			}
		}

		if s[i] == 'f' {
			if i+4 <= len(s) {
				if s[i:i+4] == "four" {
					n = append(n, "4")
					continue
				}
			}

			if i+4 <= len(s) {
				if s[i:i+4] == "five" {
					n = append(n, "5")
					continue
				}
			}
		}

		if s[i] == 's' {
			if i+3 <= len(s) {
				if s[i:i+3] == "six" {
					n = append(n, "6")
					continue
				}
			}

			if i+5 <= len(s) {
				if s[i:i+5] == "seven" {
					n = append(n, "7")
					continue
				}
			}
		}

		if s[i] == 'e' {
			if i+5 <= len(s) {
				if s[i:i+5] == "eight" {
					n = append(n, "8")
					continue
				}
			}
		}

		if s[i] == 'n' {
			if i+4 <= len(s) {
				if s[i:i+4] == "nine" {
					n = append(n, "9")
					continue
				}
			}
		}
	}

	first := n[0]
	last := "0"

	if len(n) > 1 {
		last = n[len(n)-1]
	}

	return first, last
}

func main() {
	input := helpers.ReadFile("input.txt")

	// Part 1.
	var sumPt1 int
	calibrationsPt1 := []int{}
	for _, row := range input {
		left := SeekFromLeft(row)
		right := SeekFromRight(row)

		n, _ := stoi(left + right)
		calibrationsPt1 = append(calibrationsPt1, n)
	}

	for i := 0; i < len(calibrationsPt1); i++ {
		sumPt1 += calibrationsPt1[i]
	}
	fmt.Println(sumPt1)

	// Part 2.
	var sumPt2 int
	calibrationsPt2 := []int{}
	for _, row := range input {
		first, last := SeekWordsAsNum(row)

		if last == "0" {
			fmt.Println(last)
			n, _ := stoi(first + first)
			calibrationsPt2 = append(calibrationsPt2, n)
		} else {
			n, _ := stoi(first + last)
			calibrationsPt2 = append(calibrationsPt2, n)
		}
	}

	for i := 0; i < len(calibrationsPt2); i++ {
		sumPt2 += calibrationsPt2[i]
	}
	fmt.Println(sumPt2)
}
