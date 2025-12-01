package day07

import (
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2019/golang/day05"
	"github.com/VBenny42/AoC/2019/golang/utils"
	"github.com/mowshon/iterium"
)

type day07 struct {
	program []int
}

func (d *day07) Amplify(phaseSettings []int, inputSignal int) (outputSignal int) {
	for _, phase := range phaseSettings {
		programCopy := make([]int, len(d.program))
		copy(programCopy, d.program)
		outputs := day05.ProcessDay05(programCopy, []int{phase, outputSignal})
		outputSignal = outputs[0]
	}

	return
}

func (d *day07) AmplifyFeedback(phaseSettings []int) int {
	// Create 5 amplifier instances
	amplifiers := make([]*AmplifierState, 5)
	for i := 0; i < 5; i++ {
		programCopy := make([]int, len(d.program))
		copy(programCopy, d.program)
		amplifiers[i] = &AmplifierState{
			program:    programCopy,
			ip:         0,
			inputQueue: []int{phaseSettings[i]}, // Initialize with phase
			halted:     false,
		}
	}

	signal := 0
	lastOutput := 0

	// Keep running until amplifier E halts
	for !amplifiers[4].halted {
		for i := 0; i < 5; i++ {
			// Add current signal as input
			amplifiers[i].inputQueue = append(amplifiers[i].inputQueue, signal)

			// Run amplifier until it outputs or halts
			output, halted := runUntilOutputOrHalt(amplifiers[i])

			if !halted {
				signal = output
				if i == 4 { // Amplifier E
					lastOutput = output
				}
			}
		}
	}

	return lastOutput
}

func (d *day07) Part1() (maxOutput int) {
	permutations := iterium.Permutations([]int{0, 1, 2, 3, 4}, 5)

	for {
		perm, err := permutations.Next()
		if err != nil {
			break
		}
		output := d.Amplify(perm, 0)
		if output > maxOutput {
			maxOutput = output
		}
	}

	return maxOutput
}

func (d *day07) Part2() (maxOutput int) {
	permutations := iterium.Permutations([]int{5, 6, 7, 8, 9}, 5)

	for {
		perm, err := permutations.Next()
		if err != nil {
			break
		}
		output := d.Amplify(perm, 0)
		if output > maxOutput {
			maxOutput = output
		}
	}

	return
}

func Parse(filename string) *day07 {
	line := utils.ReadTrimmed(filename)
	return &day07{
		program: utils.MapSlices(
			strings.Split(line, ","),
			utils.Atoi,
		),
	}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: highest signal that can be sent:", day.Part1())
	fmt.Println("ANSWER2: highest signal with feedback loop:", day.Part2())
}
