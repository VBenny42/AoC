package day24

import (
	"fmt"
	"os"
	"strings"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

const (
	w = iota
	x
	y
	z
)

type opcode int

const (
	add opcode = iota
	mul
	div
	mod
	eql
)

type instruction struct {
	opcode opcode
	a, b   string
}

func (o *opcode) perform(a, b int) int {
	switch *o {
	case add:
		return a + b
	case mul:
		return a * b
	case div:
		// I think go performs integer division by default
		// truncating towards zero
		return a / b
	case mod:
		return a % b
	case eql:
		if a == b {
			return 1
		}
		return 0
	default:
		panic("unknown opcode")
	}
}

type alu struct {
	instructions []instruction
}

type day24 struct {
	alus []alu
}

func (a *alu) writeToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, instr := range a.instructions {
		file.WriteString(instr.String() + "\n")
	}

	return nil
}

func (d *day24) SplitAlus() {
	for i, alu := range d.alus {
		err := alu.writeToFile(fmt.Sprintf("scratch/alu%d.txt", i))
		if err != nil {
			panic(err)
		}
	}

	return
}

func Parse(filename string) *day24 {
	var (
		data    = utils.ReadLines(filename)
		alus    = make([]alu, 0, 14)
		current alu
	)

	for i := 1; i < len(data); i++ {
		fields := strings.Fields(data[i])

		// inp line, new alu
		if len(fields) == 2 {
			alus = append(alus, current)
			current = alu{}
			continue
		}

		var instr instruction

		switch fields[0] {
		case "mul":
			instr.opcode = mul
		case "add":
			instr.opcode = add
		case "div":
			instr.opcode = div
		case "mod":
			instr.opcode = mod
		case "eql":
			instr.opcode = eql
		default:
			panic("unknown opcode")
		}

		instr.a = fields[1]
		instr.b = fields[2]

		current.instructions = append(current.instructions, instr)
	}

	alus = append(alus, current)

	if len(alus) != 14 {
		panic("invalid number of alus, got: " + fmt.Sprint(len(alus)))
	}

	return &day24{alus: alus}
}

// Solved by hand, yucky
func Solve(filename string) {
	fmt.Println("ANSWER1: largest model number accepted by MONAD: 99911993949684")
	fmt.Println("ANSWER2: smallest model number accepted by MONAD: 62911941716111")
}
