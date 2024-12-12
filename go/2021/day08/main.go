package main

/* Part 1 */
// 0: 6 segments (abcefg)
// 1: 2 segments (cf) *
// 2: 5 segments (acdeg)
// 3: 5 segments (acdfg)
// 4: 4 segments (bcdf) *
// 5: 5 segments (abdfg)
// 6: 6 segments (abdefg)
// 7: 3 segments (acf) *
// 8: 7 segments (abcdefg) *
// 9: 6 segments (abcdfg)

// 5 segments: 2, 3, 5
// 6 segments: 0, 6, 9

/* Part 2 */
// acedgfb: 8 (abcdefg)
// cdfbe: 5 (bcdef)
// gcdfa: 2 (acdfg)
// fbcad: 3 (abcdf)
// dab: 7 (abd)
// cefabd: 9 (abcdef)
// cdfgeb: 6 (bcdefg)
// eafb: 4 (abef)
// cagedb: 0 (abcdeg)
// ab: 1 (ab)

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"util"
)

func getOutputValues(input []string) []string {
	output := []string{}

	for i := 0; i < len(input); i++ {
		l := strings.Split(input[i], " | ")
		n := strings.Split(l[1], " ")
		for j := 0; j < len(n); j++ {
			output = append(output, n[j])
		}
	}

	return output
}

// TODO:
// - Code copied from Elizabeth, solve it yourself you lazy bum
func partTwo(input []string) int {
	ans := 0
	order := []string{"x", "x", "x", "x", "x", "x", "x"}
	letters := []string{"a", "b", "c", "d", "e", "f", "g"}
	for i := 0; i < len(input); i++ {
		line := input[i]
		parts := strings.Split(line, "|")
		signal := strings.Split(parts[0], " ")
		output := strings.Split(parts[1], " ")
		for x := 0; x < len(letters); x++ {
			six := 0
			two := 0
			five := 0
			four := 0
			three := 0
			seven := 0
			for n := 0; n < len(signal); n++ {
				letter := strings.Split(signal[n], "")
				for a := 0; a < len(letter); a++ {
					if letter[a] == letters[x] {
						if len(signal[n]) == 6 {
							six++
						} else if len(signal[n]) == 2 {
							two++
						} else if len(signal[n]) == 5 {
							five++
						} else if len(signal[n]) == 4 {
							four++
						} else if len(signal[n]) == 3 {
							three++
						} else {
							seven++
						}
					}
				}
			}
			if six == 3 && two == 0 && five == 3 && four == 0 && three == 1 && seven == 1 {
				order[0] = letters[x]
			}
			if six == 3 && two == 0 && five == 1 && four == 1 && three == 0 && seven == 1 {
				order[1] = letters[x]
			}
			if six == 2 && two == 1 && five == 2 && four == 1 && three == 1 && seven == 1 {
				order[2] = letters[x]
			}
			if six == 2 && two == 0 && five == 3 && four == 1 && three == 0 && seven == 1 {
				order[3] = letters[x]
			}
			if six == 2 && two == 0 && five == 1 && four == 0 && three == 0 && seven == 1 {
				order[4] = letters[x]
			}
			if six == 3 && two == 1 && five == 2 && four == 1 && three == 1 && seven == 1 {
				order[5] = letters[x]
			}
			if six == 3 && two == 0 && five == 3 && four == 0 && three == 0 && seven == 1 {
				order[6] = letters[x]
			}
		}
		zeroS := []string{}
		oneS := []string{}
		twoS := []string{}
		threeS := []string{}
		fourS := []string{}
		fiveS := []string{}
		sixS := []string{}
		sevenS := []string{}
		eightS := []string{}
		nineS := []string{}
		zeroS = append(zeroS, order[0], order[1], order[2], order[4], order[5], order[6])
		oneS = append(oneS, order[2], order[5])
		twoS = append(twoS, order[0], order[2], order[3], order[4], order[6])
		threeS = append(threeS, order[0], order[2], order[3], order[5], order[6])
		fourS = append(fourS, order[1], order[2], order[3], order[5])
		fiveS = append(fiveS, order[0], order[1], order[3], order[5], order[6])
		sixS = append(sixS, order[0], order[1], order[3], order[4], order[5], order[6])
		sevenS = append(sevenS, order[0], order[2], order[5])
		eightS = append(eightS, order[0], order[1], order[2], order[3], order[4], order[5], order[6])
		nineS = append(nineS, order[0], order[1], order[2], order[3], order[5], order[6])
		sort.Strings(zeroS)
		zeroString := strings.Join(zeroS, "")
		sort.Strings(oneS)
		oneString := strings.Join(oneS, "")
		sort.Strings(twoS)
		twoString := strings.Join(twoS, "")
		sort.Strings(threeS)
		threeString := strings.Join(threeS, "")
		sort.Strings(fourS)
		fourString := strings.Join(fourS, "")
		sort.Strings(fiveS)
		fiveString := strings.Join(fiveS, "")
		sort.Strings(sixS)
		sixString := strings.Join(sixS, "")
		sort.Strings(sevenS)
		sevenString := strings.Join(sevenS, "")
		sort.Strings(eightS)
		eightString := strings.Join(eightS, "")
		sort.Strings(nineS)
		nineString := strings.Join(nineS, "")
		// fmt.Println(output)
		deciphered := []string{}
		for i := 1; i < len(output); i++ {
			outputNumber := strings.Split(output[i], "")
			sort.Strings(outputNumber)
			outputNumberString := strings.Join(outputNumber, "")
			// fmt.Println("Just one number: ", outputNumber)
			if outputNumberString == zeroString {
				deciphered = append(deciphered, "0")
			} else if outputNumberString == oneString {
				deciphered = append(deciphered, "1")
			} else if outputNumberString == twoString {
				deciphered = append(deciphered, "2")
			} else if outputNumberString == threeString {
				deciphered = append(deciphered, "3")
			} else if outputNumberString == fourString {
				deciphered = append(deciphered, "4")
			} else if outputNumberString == fiveString {
				deciphered = append(deciphered, "5")
			} else if outputNumberString == sixString {
				deciphered = append(deciphered, "6")
			} else if outputNumberString == sevenString {
				deciphered = append(deciphered, "7")
			} else if outputNumberString == eightString {
				deciphered = append(deciphered, "8")
			} else if outputNumberString == nineString {
				deciphered = append(deciphered, "9")
			} else {
				fmt.Println("Somethin aint right")
			}
		}
		// fmt.Println(deciphered)
		ans += convert(strings.Join(deciphered, ""))
	}
	return ans

}

func convert(item string) int {
	result, _ := strconv.Atoi(item)
	return result
}

func main() {
	input := util.ReadFile("input.txt")
	outputValues := getOutputValues(input)
	count := 0

	for _, v := range outputValues {
		switch len(v) {
		case 2, 4, 3, 7:
			count++
		}
	}

	fmt.Printf("Part 1: %d\n", count)
	fmt.Printf("Part 2: %d\n", partTwo(input))
}
