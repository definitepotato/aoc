package day2

import (
	"fmt"
	"strconv"
	"strings"
	"util"
)

func ProcessOpcodes(opcodes string, pos int) string {
	s := strings.Split(opcodes, ",")

	if util.Stoi(s[pos]) == 99 {
		return strings.Join(s, ",")
	}

	num0 := util.Stoi(s[pos])
	num1 := util.Stoi(s[pos+1])
	num2 := util.Stoi(s[pos+2])
	resultPos := util.Stoi(s[pos+3])

	// Add opcode.
	if num0 == 1 {
		result := util.Stoi(s[num1]) + util.Stoi(s[num2])
		s[resultPos] = strconv.Itoa(result)
	}

	// Multiply opcode.
	if num0 == 2 {
		result := util.Stoi(s[num1]) * util.Stoi(s[num2])
		s[resultPos] = strconv.Itoa(result)
	}

	return ProcessOpcodes(strings.Join(s, ","), pos+4)
}

func main() {
	opcodes := util.ReadFile("input.txt")

	// Restore original state before the last computer caught fire.
	opcodes = strings.Split(opcodes[0], ",")
	opcodes[1] = "12"
	opcodes[2] = "2"
	originalOpcodes := strings.Join(opcodes, ",")

	// Part one.
	result := ProcessOpcodes(originalOpcodes, 0)
	fmt.Println(result)
}
