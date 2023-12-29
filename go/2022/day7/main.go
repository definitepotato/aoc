package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"util"
)

func SplitCommand(s string) []string {
	return strings.Split(s, " ")[1:]
}

func part1() {
	commands := util.ReadFile("input.txt")

	var stack []int
	var sum int

	for _, command := range commands {
		if strings.Contains(command, "$ cd") {
			c := SplitCommand(command)
			if c[1] != ".." {
				stack = append(stack, 0) // cd into new dir
			}

			if c[1] == ".." {
				dirSize := stack[len(stack)-1] // pop last dir to cd out

				if dirSize <= 100000 {
					sum += dirSize
				}

				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
					stack[len(stack)-1] += dirSize
				}
			}

		}

		fs := strings.Fields(command)
		fileSize := util.Stoi(fs[0])
		stack[len(stack)-1] += fileSize
	}

	fmt.Println(sum)
}

func part2() {
	sc := bufio.NewScanner(os.Stdin)

	var stack, sizes []int
	popd := func() { // cd out of dir, compute dirSize
		dirSize := stack[len(stack)-1]
		sizes = append(sizes, dirSize)
		if stack = stack[:len(stack)-1]; len(stack) > 0 {
			stack[len(stack)-1] += dirSize
		}
	}

	for sc.Scan() {
		switch line := sc.Text(); line[:4] {
		case "$ cd":
			if path := line[5:]; path != ".." {
				stack = append(stack, 0) // cd into new dir
			} else {
				popd()
			}
		case "$ ls", "dir ": // Nothing to do, next line
		default: // file
			fs := strings.Fields(line)
			fileSize, _ := strconv.Atoi(fs[0])
			stack[len(stack)-1] += fileSize
		}
	}

	for len(stack) > 0 {
		popd()
	}

	sort.Ints(sizes)

	used := sizes[len(sizes)-1]
	free := 70000000 - used

	i := sort.Search(len(sizes), func(i int) bool {
		return free+sizes[i] >= 30000000
	})

	fmt.Println(sizes[i])
}

func main() {
	part1()
	part2()
}
