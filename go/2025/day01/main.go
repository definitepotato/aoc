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

func main() {
	lines := ReadFile("input.txt")

	dial := 50
	start := dial

	part1 := 0
	part2 := 0

	for _, line := range lines {
		dir := line[0]
		distance, _ := strconv.Atoi(line[1:])

		turns := distance / 100
		part2 += turns

		if dir == 'R' {
			if dial+distance%100 >= 100 {
				part2 += 1
			}
			dial += distance
		}

		if dir == 'L' {
			if dial > 0 && (dial-distance%100) <= 0 {
				part2 += 1
			}
			start = (100 - start) % 100
			dial -= distance
		}

		dial = dial % 100
		if dial == 0 {
			part1 += 1
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
