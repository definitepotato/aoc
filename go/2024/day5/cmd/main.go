// Shamelessly copied from the internet. This is not my code but wanted to learn
// how this worked. It's pretty damn clean. I failed today.

package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("../input.txt")

	// [0] => rules, [1] => pages
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	// rules for sorting
	cmp := func(a, b string) int {
		for _, s := range strings.Split(split[0], "\n") {
			if s := strings.Split(s, "|"); s[0] == a && s[1] == b {
				return -1
			}
		}
		return 0
	}

	run := func(sorted bool) (r int) {
		for _, s := range strings.Split(split[1], "\n") {
			// check pages are sorted
			if s := strings.Split(s, ","); slices.IsSortedFunc(s, cmp) == sorted {
				// if not sorted then sort according to rules
				slices.SortFunc(s, cmp)

				n, _ := strconv.Atoi(s[len(s)/2]) // get middle number of pages as int
				r += n
			}
		}
		return r
	}

	fmt.Printf("Part 1: %d\n", run(true))
	fmt.Printf("Part 2: %d\n", run(false))
}
