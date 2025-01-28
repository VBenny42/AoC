package day12

import (
	"fmt"
	"slices"
	"strings"
	"sync"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type (
	condition int
	row       struct {
		springs []condition
		groups  []int
	}
)

//go:generate stringer -type=condition
const (
	operational condition = iota
	damaged
	unknown
)

type day12 struct {
	rows []row
}

// borrowed from
// https://github.com/ConcurrentCrab/AoC/blob/aba36645b18566bbb7028437a0929d4b6af0e5f5/solutions/12-1.go
func possibleArrangements(r row) (possible int) {
	type state struct {
		groupIdx          int
		groupSize         int
		wantWorkingSpring bool
	}

	var (
		currStates = map[state]int{{0, 0, false}: 1}
		tempStates = map[state]int{}
	)

	for _, spring := range r.springs {
		for curr, count := range currStates {
			temp := curr

			switch {
			// Adding damaged spring to a group
			case (spring == damaged || spring == unknown) &&
				temp.groupIdx < len(r.groups) &&
				!temp.wantWorkingSpring:

				if spring == unknown && temp.groupSize == 0 {
					tempStates[curr] += count
				}

				// Add the damaged spring to the group
				temp.groupSize++

				// Move to the next group, looking for a working spring
				if temp.groupSize == r.groups[temp.groupIdx] {
					temp.groupIdx, temp.groupSize = temp.groupIdx+1, 0
					temp.wantWorkingSpring = true
				}

				tempStates[temp] += count

			// If not in a group, we can mark ? as operational, or . as operational
			case (spring == operational || spring == unknown) &&
				temp.groupSize == 0:

				temp.wantWorkingSpring = false
				tempStates[temp] += count
			}
		}

		currStates, tempStates = tempStates, currStates
		clear(tempStates)
	}

	for state, count := range currStates {
		if state.groupIdx == len(r.groups) {
			possible += count
		}
	}

	return
}

func (d *day12) Part1() (sum int) {
	for _, r := range d.rows {
		sum += possibleArrangements(r)
	}
	return
}

func (d *day12) Part2() (sum int) {
	var (
		wg sync.WaitGroup
		ch = make(chan int)
	)

	for _, r := range d.rows {
		wg.Add(1)
		go func(r row) {
			defer wg.Done()

			r.springs = append(r.springs, unknown)
			repeatedRow := row{
				springs: slices.Repeat(r.springs, 5),
				groups:  slices.Repeat(r.groups, 5),
			}
			repeatedRow.springs = repeatedRow.springs[:len(repeatedRow.springs)-1]

			ch <- possibleArrangements(repeatedRow)
		}(r)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for n := range ch {
		sum += n
	}

	return
}

func Parse(filename string) *day12 {
	var (
		data = utils.ReadLines(filename)
		rows = make([]row, len(data))
	)

	for i, line := range data {
		left, right, _ := strings.Cut(line, " ")

		rows[i].springs = make([]condition, len(left))
		for j, c := range left {
			switch c {
			case '.':
				rows[i].springs[j] = operational
			case '#':
				rows[i].springs[j] = damaged
			case '?':
				rows[i].springs[j] = unknown
			default:
				panic("invalid condition")
			}
		}

		damaged := strings.Split(right, ",")
		rows[i].groups = make([]int, len(damaged))
		for j, d := range damaged {
			rows[i].groups[j] = utils.Atoi(d)
		}
	}

	return &day12{rows: rows}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: sum of possible arrangements:", day.Part1())
	fmt.Println("ANSWER1: sum of possible unfolded arrangements:", day.Part2())
}
