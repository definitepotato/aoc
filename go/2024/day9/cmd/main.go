package main

import (
	"fmt"
	"helpers"
	"strconv"
)

func checksum(fs []string) int {
	sum := 0
	for idx := 0; idx < len(fs); idx++ {
		n, _ := strconv.Atoi(fs[idx])
		sum += idx * n
	}

	return sum
}

func nextEmptyIdx(fs []string, upto int) int {
	for idx := 0; idx < len(fs); idx++ {
		if idx >= upto {
			return -1
		}

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

		nextEmptyBlock := nextEmptyIdx(fs, idx)
		// if -1 we couldn't find an empty block large enough to fit the file
		// move on to the next file
		if nextEmptyBlock == -1 {
			continue
		}

		fs[nextEmptyBlock] = fs[idx]
		fs[idx] = "."
	}
}

func fileLen(fs []string, filename string) int {
	length := 0
	for _, val := range fs {
		if val == filename {
			length += 1
		}
	}

	return length
}

func blockLen(fs []string, idx int) int {
	length := 0
	for i := idx; i < len(fs); i++ {
		if fs[i] != "." {
			return length
		}
		length += 1
	}

	return length
}

func nextEmptyIdxWithLen(fs []string, l int, upto int) int {
	for i := 0; i < len(fs); i++ {
		if i >= upto {
			return -1
		}

		length := blockLen(fs, i)
		if length >= l {
			return i
		}
	}
	return -1
}

func moveFileV2(fs []string) {
	// iterate file system in reverse
	for idx := len(fs) - 1; idx >= 0; idx-- {
		if fs[idx] == "." {
			continue
		}

		// get the length of the current file
		fileLength := fileLen(fs, fs[idx])
		emptyBlockIdx := nextEmptyIdxWithLen(fs, fileLength, idx)

		// if -1 we couldn't find an empty block large enough to fit the file
		// move on to the next file
		if emptyBlockIdx == -1 {
			continue
		}

		// move the file index by index swapping with empty indexes
		for r := 0; r < fileLength; r++ {
			// fmt.Printf("R%d: %s@%d => %s@%d\n", r, fs[idx-r], idx-r, fs[emptyBlockIdx+r], emptyBlockIdx+r)
			fs[emptyBlockIdx+r] = fs[idx-r]
			fs[idx-r] = "."
		}

		idx -= fileLength - 1
	}
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

	eFS2 := make([]string, len(eFS))
	copy(eFS2, eFS)

	moveFile(eFS)
	fmt.Printf("Part 1: %d\n", checksum(eFS))

	moveFileV2(eFS2)
	fmt.Printf("Part 2: %d\n", checksum(eFS2))
}
