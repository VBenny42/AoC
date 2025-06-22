package day14

import (
	"fmt"

	"github.com/VBenny42/AoC/2020/golang/utils"
)

type day14 struct {
	lines  []string
	memory map[int]int
}

func (d *day14) Part1() (sum int) {
	var (
		mask        string
		addr, value int
	)

	for _, line := range d.lines {
		if line[:4] == "mask" {
			mask = line[7:]
			continue
		} else {
			fmt.Sscanf(line, "mem[%d] = %d", &addr, &value)

			for i := 0; i < len(mask); i++ {
				if mask[len(mask)-1-i] == '1' {
					value |= (1 << i)
				} else if mask[len(mask)-1-i] == '0' {
					value &= ^(1 << i)
				}
			}

			d.memory[addr] = value
		}
	}

	for _, v := range d.memory {
		sum += v
	}

	return
}

func (d *day14) Part2() (sum int) {
	d.memory = make(map[int]int) // Reset memory for Part 2

	var (
		mask        string
		addr, value int
	)

	for _, line := range d.lines {
		if line[:4] == "mask" {
			mask = line[7:]
			continue
		} else {
			fmt.Sscanf(line, "mem[%d] = %d", &addr, &value)

			for i := 0; i < len(mask); i++ {
				if mask[len(mask)-1-i] == '1' {
					addr |= (1 << i)
				}
			}

			var xPositions []int
			for i := 0; i < len(mask); i++ {
				if mask[len(mask)-1-i] == 'X' {
					xPositions = append(xPositions, i)
				}
			}

			numCombinations := 1 << len(xPositions) // 2^n
			for combo := 0; combo < numCombinations; combo++ {
				currentAddr := addr

				for j, pos := range xPositions {
					if (combo>>j)&1 == 1 {
						currentAddr |= (1 << pos) // Set bit to 1
					} else {
						currentAddr &= ^(1 << pos) // Set bit to 0
					}
				}

				d.memory[currentAddr] = value
			}
		}
	}

	for _, v := range d.memory {
		sum += v
	}

	return
}

func Parse(filename string) *day14 {
	return &day14{
		lines:  utils.ReadLines(filename),
		memory: make(map[int]int),
	}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: sum of values in memory after initialization:", day.Part1())
	fmt.Println("ANSWER2: sum of values in memory after initialization with floating vals:", day.Part2())
}
