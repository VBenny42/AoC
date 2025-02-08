package day08

import (
	"fmt"
	"sort"
	"strings"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

type entry struct {
	uniqueDigits   []string
	appearedDigits []string
}

type day08 struct {
	entries []entry
}

// unique mapping lengths
const (
	one   = 2
	four  = 4
	seven = 3
	eight = 7
)

func removeIndices(s []string, indices ...int) (result []string) {
	m := make(map[int]struct{}, len(indices))
	for _, i := range indices {
		m[i] = struct{}{}
	}

	for i, v := range s {
		if _, ok := m[i]; !ok {
			result = append(result, v)
		}
	}

	return
}

func overlaps(s1, s2 string) bool {
	if len(s2) > len(s1) {
		s1, s2 = s2, s1
	}

	m := make(map[rune]struct{}, len(s1))
	for _, c := range s1 {
		m[c] = struct{}{}
	}

	for _, c := range s2 {
		if _, ok := m[c]; !ok {
			return false
		}
	}

	return true
}

func (d *day08) Part1() (sum int) {
	// need to sum the number of 1, 4, 7, 8 digits that appear
	for _, entry := range d.entries {
		for _, digit := range entry.appearedDigits {
			switch len(digit) {
			case one, four, seven, eight:
				sum++
			}
		}
	}
	return
}

func (d *day08) Part2() (sum int) {
	// figuring out what wires map to what
	var (
		indexToMapping  = make([]string, 10)
		indicesToRemove = make([]int, 0, 10)
	)
	for _, entry := range d.entries {
		indicesToRemove = indicesToRemove[:0]

		for i, digit := range entry.uniqueDigits {
			switch len(digit) {
			case one:
				indexToMapping[1] = digit
				indicesToRemove = append(indicesToRemove, i)
			case four:
				indexToMapping[4] = digit
				indicesToRemove = append(indicesToRemove, i)
			case seven:
				indexToMapping[7] = digit
				indicesToRemove = append(indicesToRemove, i)
			case eight:
				indexToMapping[8] = digit
				indicesToRemove = append(indicesToRemove, i)
			}
		}

		entry.uniqueDigits = removeIndices(entry.uniqueDigits, indicesToRemove...)
		indicesToRemove = indicesToRemove[:0]

		var zeroThreeNine []string

		for i, digit := range entry.uniqueDigits {
			if overlaps(digit, indexToMapping[1]) {
				zeroThreeNine = append(zeroThreeNine, digit)
				indicesToRemove = append(indicesToRemove, i)
			}
		}

		for i, digit := range zeroThreeNine {
			// This must be the 3 digit
			if len(digit) == 5 {
				indexToMapping[3] = digit
				zeroThreeNine = removeIndices(zeroThreeNine, i)
				break
			}
		}

		for i, digit := range zeroThreeNine {
			// Must be the 9 digit
			if overlaps(digit, indexToMapping[4]) {
				indexToMapping[9] = digit
				zeroThreeNine = removeIndices(zeroThreeNine, i)
				break
			}
		}

		// Only one digit left, must be the 0 digit
		indexToMapping[0] = zeroThreeNine[0]

		entry.uniqueDigits = removeIndices(entry.uniqueDigits, indicesToRemove...)

		for i, digit := range entry.uniqueDigits {
			if len(digit) == 6 {
				indexToMapping[6] = digit
				entry.uniqueDigits = removeIndices(entry.uniqueDigits, i)
				break
			}
		}

		for i, digit := range entry.uniqueDigits {
			if overlaps(digit, indexToMapping[9]) {
				indexToMapping[5] = digit
				entry.uniqueDigits = removeIndices(entry.uniqueDigits, i)
				break
			}
		}

		// Only one digit left, must be the 2 digit
		indexToMapping[2] = entry.uniqueDigits[0]

		// Now we have all the digits mapped to the correct index
		// Can figure out the sum
		var num int
		for _, digit := range entry.appearedDigits {
			for i, mapping := range indexToMapping {
				if digit == mapping {
					num *= 10
					num += i
				}
			}
		}

		sum += num
	}

	return
}

func Parse(filename string) *day08 {
	var (
		data    = utils.ReadLines(filename)
		entries = make([]entry, len(data))
	)

	sortString := func(s string) string {
		chars := strings.Split(s, "")
		sort.Strings(chars)
		return strings.Join(chars, "")
	}

	for i, line := range data {
		var (
			left, right, found = strings.Cut(line, "|")
			leftFields         = strings.Fields(left)
			rightFields        = strings.Fields(right)
		)
		if found == false || len(leftFields) != 10 || len(rightFields) != 4 {
			panic("Invalid input")
		}

		entries[i].uniqueDigits = make([]string, 10)
		for j, field := range leftFields {
			entries[i].uniqueDigits[j] = sortString(field)
		}

		entries[i].appearedDigits = make([]string, 4)
		for j, field := range rightFields {
			entries[i].appearedDigits[j] = sortString(field)
		}
	}

	return &day08{entries: entries}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: times that '1', '4', '7' or '8' appear:", day.Part1())
	fmt.Println("ANSWER2: sum of the actual numbers:", day.Part2())
}
