package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadFile(filename string) []string {
	var lines []string

	readFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines
}

func part2() {
	lines := ReadFile("input.txt")
	dial := 50
	start := dial
	ans := 0

	for _, line := range lines {
		dir := -1 // left
		if line[0] == 'R' {
			dir = 1
		}
		distance, _ := strconv.Atoi(line[1:])

		turns := distance / 100
		ans += turns

		dial += dir * (distance % 100)

		if start != 0 && (dial < 0 || dial > 100) {
			ans += 1
		}

		dial %= 100
		if dial < 0 {
			dial += 100
		}

		if start != 0 && dial == 0 {
			ans += 1
		}
	}

	fmt.Printf("Part 2: %d\n", ans)
}

func part1() {
	lines := ReadFile("input.txt")
	dial := 50
	ans := 0

	for _, line := range lines {
		dir := -1 // left
		if line[0] == 'R' {
			dir = 1
		}
		distance, _ := strconv.Atoi(line[1:])

		dial += dir * distance
		dial %= 100
		if dial == 0 {
			ans += 1
		}
	}

	fmt.Printf("Part 1: %d\n", ans)
}

func main() {
	part1()
	part2()
}
