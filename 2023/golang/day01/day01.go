package day01

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type day01 struct {
	lines []string
}

func (d *day01) Part1() (sum int) {
	pattern := regexp.MustCompile(`\d`)

	for _, line := range d.lines {
		matches := pattern.FindAllString(line, -1)
		if len(matches) == 0 {
			panic("No matches found")
		}
		first := matches[0]
		last := matches[len(matches)-1]
		sum += 10*utils.Atoi(first) + utils.Atoi(last)
	}

	return
}

func (d *day01) Part2() (sum int) {
	digitsMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for _, line := range d.lines {
		var firstDigit, lastDigit int
		firstIndex, lastIndex := len(line), -1

		for word, digit := range digitsMap {
			digitIndex := strings.Index(line, strconv.Itoa(digit))
			if digitIndex != -1 && digitIndex < firstIndex {
				firstIndex = digitIndex
				firstDigit = digit
			}
			if digitIndex != -1 && digitIndex > lastIndex {
				lastIndex = digitIndex
				lastDigit = digit
			}

			// Check word form
			for i := 0; i <= len(line)-len(word); i++ {
				if line[i:i+len(word)] == word {
					if i < firstIndex {
						firstIndex = i
						firstDigit = digit
					}
					if i > lastIndex {
						lastIndex = i
						lastDigit = digit
					}
				}
			}
		}

		sum += 10*firstDigit + lastDigit
	}

	return
}

func Parse(filename string) *day01 {
	data := utils.ReadLines(filename)

	return &day01{lines: data}
}

func Solve(filename string) {
	d := Parse(filename)

	fmt.Println("ANSWER1: sum of calibration values:", d.Part1())
	fmt.Println("ANSWER2: sum of calibration values with digits as words:", d.Part2())
}
