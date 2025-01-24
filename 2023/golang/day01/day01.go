package day01

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	// "sync"
	// "sync/atomic"

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
		"zero":  0,
	}
	for i := 0; i <= 9; i++ {
		digitsMap[strconv.Itoa(i)] = i
	}

	// Goroutines shave off about 4ms, but I don't think it's worth it

	// var (
	// 	wg        sync.WaitGroup
	// 	sumAtomic atomic.Int32
	// )
	// wg.Add(len(d.lines))

	for _, line := range d.lines {
		// go func(line string) {
		// defer wg.Done()
		var firstDigit, lastDigit int

		index := 0

		for index < len(line) {
			for word, digit := range digitsMap {
				if strings.HasPrefix(line[index:], word) {
					if firstDigit == 0 {
						firstDigit = digit
					}
					lastDigit = digit
					break
				}
			}
			index++
		}

		// sumAtomic.Add(int32(10*firstDigit + lastDigit))
		sum += 10*firstDigit + lastDigit
		// }(line)
	}

	// wg.Wait()

	// return int(sumAtomic.Load())
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
