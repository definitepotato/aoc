package main

import (
	"fmt"
	"helpers"
	"strconv"
)

func nextEmptyIdx(fs []string) int {
	for idx := 0; idx < len(fs); idx++ {
		if fs[idx] == "." {
			return idx
		}
	}
	return -1
}

func moveFile(fs []string) {
	// iterate filesystem in reverse
	for idx := len(fs) - 1; idx >= 0; idx-- {
		if fs[idx] == "." {
			continue
		}
		nextEmptyBlock := nextEmptyIdx(fs)
		fs[nextEmptyBlock] = fs[idx]
		fs[idx] = "."
	}
}

func checksum(fs []string) int {
	sum := 0
	for idx := 1; idx < len(fs); idx++ {
		n, _ := strconv.Atoi(fs[idx])
		sum += (idx - 1) * n
	}

	return sum
}

func main() {
	file := helpers.ReadFile("../input.txt")

	eFS := []string{}
	for _, v := range file {
		// each file on disk has an ID number based on the order
		// of the files as they appear before they are rearranged
		// start with ID 0, when expanded they will appear on the
		// filesystem as this ID number
		fileId := 0

		// iterate length of the filesystem blocks
		for i := 0; i < len(v); i += 2 {
			// convert the file size we find at the index to integer
			files, _ := strconv.Atoi(string(v[i]))

			// iterate the number of times equal to the number found at
			// the index and insert the fileId number that many times
			for fileCount := 0; fileCount < files; fileCount++ {
				eFS = append(eFS, strconv.Itoa(fileId))
			}

			// the last block is a file without empty space beyond
			// it, we want to prevent index out of range so check
			// for that boundary
			if i < len(v)-1 {
				// convert the number of empty blocks we find at the index to integer
				blocks, _ := strconv.Atoi(string(v[i+1]))

				// iterate the number of times equal to the number found at
				// the index and insert the "." character to signal empty space
				for blockCount := 0; blockCount < blocks; blockCount++ {
					eFS = append(eFS, ".")
				}
			}

			// we're moving to the next file in the next iteration
			// so increment fileId here
			fileId += 1
		}
	}

	moveFile(eFS)
	fmt.Println(checksum(eFS))
}
