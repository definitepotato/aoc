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

	for _, line := range lines {
		dir := line[0]
		distance, _ := strconv.Atoi(line[1:])

		if dir == 'R' {
			dial += distance
		}

		if dir == 'L' {
			start = (100 - start) % 100
			dial -= distance
		}

		if dial%100 == 0 {
			part1 += 1
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
}
