package day05

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type instruction struct {
	moveAmount int
	srcStack   int
	dstStack   int
}

type day05 struct {
	Stacks       [][]rune
	Instructions []instruction
}

func parseCrates(line string, numStacks int) []rune {
	step := 4

	crates := make([]rune, numStacks)

	for i := 0; i < len(line); i += step {
		end := i + step
		if end > len(line) {
			end = len(line)
		}

		if i+1 < len(line) && line[i+1] != ' ' {
			crates[i/step] = rune(line[i+1])
		}
	}

	return crates
}

func (d *day05) Part1() string {
	var srcLen int
	var movedElements []rune

	for _, instruction := range d.Instructions {
		srcStack := d.Stacks[instruction.srcStack]
		dstStack := &d.Stacks[instruction.dstStack]

		srcLen = len(srcStack)
		movedElements = srcStack[srcLen-instruction.moveAmount:]
		slices.Reverse(movedElements)

		d.Stacks[instruction.srcStack] = srcStack[:srcLen-instruction.moveAmount]
		*dstStack = append(*dstStack, movedElements...)
	}

	topVals := strings.Builder{}

	for _, stack := range d.Stacks {
		topVals.WriteRune(stack[len(stack)-1])
	}

	return topVals.String()
}

func (d *day05) Part2() string {
	var srcLen int
	var movedElements []rune

	for _, instruction := range d.Instructions {
		srcStack := d.Stacks[instruction.srcStack]
		dstStack := &d.Stacks[instruction.dstStack]

		srcLen = len(srcStack)
		movedElements = srcStack[srcLen-instruction.moveAmount:]

		d.Stacks[instruction.srcStack] = srcStack[:srcLen-instruction.moveAmount]
		*dstStack = append(*dstStack, movedElements...)
	}

	topVals := strings.Builder{}

	for _, stack := range d.Stacks {
		topVals.WriteRune(stack[len(stack)-1])
	}

	return topVals.String()
}

func Parse(filename string) *day05 {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.Trim(string(file), "\n"), "\n")

	stackIndex := 0

	for i, line := range lines {
		if line == "" {
			stackIndex = i - 1
			break
		}
	}

	numStacks := len(strings.Fields(lines[stackIndex]))
	stacks := make([][]rune, numStacks)

	for i := 0; i < stackIndex; i++ {
		crates := parseCrates(lines[i], numStacks)
		for j, crate := range crates {
			if crate == rune(0) {
				continue
			}
			stacks[j] = append(stacks[j], crate)
		}
	}

	for i := range stacks {
		slices.Reverse(stacks[i])
	}

	instructions := make([]instruction, len(lines)-(stackIndex+2))

	for i, line := range lines[stackIndex+2:] {
		fmt.Sscanf(line, "move %d from %d to %d",
			&instructions[i].moveAmount,
			&instructions[i].srcStack,
			&instructions[i].dstStack)
		instructions[i].srcStack--
		instructions[i].dstStack--
	}

	return &day05{stacks, instructions}
}

func Solve(filename string) {
	fmt.Println("ANSWER1: top crates on each stack:", Parse(filename).Part1())
	fmt.Println("ANSWER2: top crates on each stack:", Parse(filename).Part2())
}
