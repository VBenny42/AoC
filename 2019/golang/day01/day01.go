package day01

import (
	"fmt"

	"github.com/VBenny42/AoC/2019/golang/utils"
)

type day01 struct {
	masses []int
}

func calculateFuel(mass int) int {
	return mass/3 - 2
}

func (d *day01) Part1() (sum int) {
	for _, mass := range d.masses {
		sum += calculateFuel(mass)
	}

	return
}

func (d *day01) Part2() (sum int) {
	for _, mass := range d.masses {
		fuel := calculateFuel(mass)
		for fuel > 0 {
			sum += fuel
			fuel = calculateFuel(fuel)
		}
	}

	return
}

func Parse(filename string) *day01 {
	lines := utils.ReadLines(filename)
	return &day01{masses: utils.MapSlices(lines, utils.Atoi)}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: sum of the fuel requirements for all the modules:", day.Part1())
	fmt.Println("ANSWER2: sum of the fuel requirements for all the modules including fuel for the fuel:", day.Part2())
}
