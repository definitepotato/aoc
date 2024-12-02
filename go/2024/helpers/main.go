package helpers

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
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

func Stoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func SliceStringToInt(input []string) []int {
	var newInput []int

	for _, v := range input {
		newInput = append(newInput, Stoi(v))
	}

	return newInput
}

func SortString(input string) string {
	s := strings.Split(input, "")
	sort.Strings(s)

	return strings.Join(s, "")
}

func Reverse(input []string) []string {
	reversed := []string{}

	for v := range input {
		x := input[len(input)-1-v]
		reversed = append(reversed, x)
	}
	return reversed
}

// Returns unique items in a slice.
func Unique(slice []string) []string {
	// create a map with all the values as key
	uniqMap := make(map[string]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	// turn the map keys into a slice
	uniqSlice := make([]string, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}
