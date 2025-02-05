package day03

import (
	"fmt"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

type day03 struct {
	numbers [][]int
}

func (d *day03) Part1() int {
	var (
		gamma, epsilon []int
		width          = len(d.numbers[0])
		length         = len(d.numbers)
		mostCommon     int
	)

	for i := 0; i < width; i++ {
		var count int

		for _, number := range d.numbers {
			count += number[i]
		}

		if count > length/2 {
			// most common character is 1
			mostCommon = 1
		} else {
			// most common character is 0
			mostCommon = 0
		}
		gamma = append(gamma, mostCommon)
		epsilon = append(epsilon, 1-mostCommon)
	}

	var gammaNumber, epsilonNumber int

	for i := 0; i < width; i++ {
		gammaNumber += gamma[i] << (len(gamma) - i - 1)
		epsilonNumber += epsilon[i] << (len(epsilon) - i - 1)
	}

	return gammaNumber * epsilonNumber
}

func (d *day03) Part2() int {
	var (
		oxygen, co2 int
		width       = len(d.numbers[0])
	)

	numbersToCheck := make([][]int, len(d.numbers))
	copy(numbersToCheck, d.numbers)

	for i := 0; i < width; i++ {
		if len(numbersToCheck) == 1 {
			break
		}

		var (
			count               int
			tmpNumbers          [][]int
			mostCommonCharacter int
		)

		for _, number := range numbersToCheck {
			count += number[i]
		}

		// go shenanigans
		// integer division rounding down
		if count*2 >= len(numbersToCheck) {
			mostCommonCharacter = 1
		} else {
			mostCommonCharacter = 0
		}
		for _, number := range numbersToCheck {
			if number[i] == mostCommonCharacter {
				tmpNumbers = append(tmpNumbers, number)
			}
		}

		numbersToCheck = tmpNumbers
	}

	for i := 0; i < len(numbersToCheck[0]); i++ {
		oxygen += numbersToCheck[0][i] << (len(numbersToCheck[0]) - i - 1)
	}

	numbersToCheck = make([][]int, len(d.numbers))
	copy(numbersToCheck, d.numbers)

	for i := 0; i < width; i++ {
		if len(numbersToCheck) == 1 {
			break
		}

		var (
			count                int
			tmpNumbers           [][]int
			leastCommonCharacter int
		)

		for _, number := range numbersToCheck {
			count += number[i]
		}

		if count*2 >= len(numbersToCheck) {
			leastCommonCharacter = 0
		} else {
			leastCommonCharacter = 1
		}
		for _, number := range numbersToCheck {
			if number[i] == leastCommonCharacter {
				tmpNumbers = append(tmpNumbers, number)
			}
		}

		numbersToCheck = tmpNumbers
	}

	for i := 0; i < len(numbersToCheck[0]); i++ {
		co2 += numbersToCheck[0][i] << (len(numbersToCheck[0]) - i - 1)
	}

	return oxygen * co2
}

func Parse(filename string) *day03 {
	var (
		data    = utils.ReadLines(filename)
		numbers = make([][]int, len(data))
	)

	for i, line := range data {
		numbers[i] = make([]int, len(line))
		for j, char := range line {
			numbers[i][j] = int(char - '0')
		}
	}

	return &day03{
		numbers: numbers,
	}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: power consumption of submarine:", day.Part1())
	fmt.Println("ANSWER2: life support rating of submarine:", day.Part2())
}
