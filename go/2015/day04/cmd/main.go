package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
)

func hash(s string) string {
	h := md5.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}

func main() {
	found5 := false

	for i := 0; i < math.MaxInt; i++ {
		input := fmt.Sprintf("%s%d", "bgvyzdsv", i)

		h := hash(input)

		if h[0:5] == "00000" {
			if !found5 {
				fmt.Printf("Part 1: %d (%s)\n", i, h)
			}

			found5 = true
		}

		if h[0:6] == "000000" {
			fmt.Printf("Part 2: %d (%s)\n", i, h)
			break
		}
	}
}
