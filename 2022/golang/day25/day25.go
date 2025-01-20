package day25

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type day25 struct {
	sum int
}

var digits = map[rune]int{
	'2': 2,
	'1': 1,
	'0': 0,
	'-': -1,
	'=': -2,
}

func parseLine(line string) (sum int) {
	length := len(line)

	for i, r := range line {
		sum += digits[r] * (int(math.Pow(5, float64(length-i-1))))
	}

	return
}

func convertToSnafu(sum int) string {
	snafu := []string{}

	remainders := []int{}
	for sum > 0 {
		quotient := sum / 5
		remainder := sum % 5
		remainders = append(remainders, remainder)
		sum = quotient
	}

	carry := 0

	for _, remainder := range remainders {
		withCarry := remainder + carry

		if withCarry > 2 {
			carry = 1
		} else {
			carry = 0
		}

		switch withCarry {
		case 3:
			snafu = append(snafu, "=")
		case 4:
			snafu = append(snafu, "-")
		default:
			snafu = append(snafu, fmt.Sprint(withCarry%5))
		}
	}

	if carry == 1 {
		snafu = append(snafu, "1")
	}

	slices.Reverse(snafu)

	return strings.Join(snafu, "")
}

func (d *day25) Part1() string {
	return convertToSnafu(d.sum)
}

func Parse(filename string) *day25 {
	data := utils.ReadLines(filename)

	sum := 0
	for _, line := range data {
		sum += parseLine(line)
	}

	return &day25{sum}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: SNAFU number to supply:", day.Part1())
}
