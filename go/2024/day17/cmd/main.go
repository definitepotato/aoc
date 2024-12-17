package main

import (
	"fmt"
)

type Registers struct {
	A      int
	B      int
	C      int
	InsPtr int
}

// combo operand
func (r *Registers) combo(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return r.A
	case 5:
		return r.B
	case 6:
		return r.C
	case 7:
		return 0
	default:
		return 0
	}
}

// opcode 0 (adv)
func (r *Registers) adv(operand int) {
	r.A >>= r.combo(operand)
	r.InsPtr += 2
}

// opcode 1 (bxl)
func (r *Registers) bxl(operand int) {
	r.B ^= operand
	r.InsPtr += 2
}

// opcode 2 (bst)
func (r *Registers) bst(operand int) {
	r.B = r.combo(operand) % 8
	r.InsPtr += 2
}

// opcode 3 (jnz)
func (r *Registers) jnz(operand int) {
	if r.A == 0 {
		r.InsPtr += 2
		return
	}
	r.InsPtr = operand
}

// opcode 4 (bxc)
func (r *Registers) bxc(operand int) {
	_ = operand // ignore operand

	r.B ^= r.C
	r.InsPtr += 2
}

// opcode 5 (out)
func (r *Registers) out(operand int) int {
	r.InsPtr += 2
	return r.combo(operand) % 8
}

// opcode 6 (bdv)
func (r *Registers) bdv(operand int) {
	r.B = r.A >> r.combo(operand)
	r.InsPtr += 2
}

// opcode 7 (cdv)
func (r *Registers) cdv(operand int) {
	r.C = r.A >> r.combo(operand)
	r.InsPtr += 2
}

func main() {
	cpu := Registers{
		A:      729,
		B:      0,
		C:      0,
		InsPtr: 0,
	}
	program := []int{0, 1, 5, 4, 3, 0}

	var ans []int
	for cpu.InsPtr+1 < len(program) {
		switch program[cpu.InsPtr] {
		case 0:
			cpu.adv(program[cpu.InsPtr+1])
			continue
		case 1:
			cpu.bxl(program[cpu.InsPtr+1])
			continue
		case 2:
			cpu.bst(program[cpu.InsPtr+1])
			continue
		case 3:
			cpu.jnz(program[cpu.InsPtr+1])
			continue
		case 4:
			cpu.bxc(program[cpu.InsPtr+1])
			continue
		case 5:
			out := cpu.out(program[cpu.InsPtr+1])
			ans = append(ans, out)
			continue
		case 6:
			cpu.bdv(program[cpu.InsPtr+1])
			continue
		case 7:
			cpu.cdv(program[cpu.InsPtr+1])
			continue
		}
	}

	fmt.Printf("Part 1:%d\n", ans)
}
