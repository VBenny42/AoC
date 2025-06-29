package day23

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/VBenny42/AoC/2020/golang/utils"
	"github.com/VBenny42/AoC/2020/golang/utils/set"
)

type day23 struct {
	currentCup   int
	currentIndex int
	cups         []int
	maxCup       int
}

func (d *day23) crabMove() {
	removedCups := make([]int, 0, 3)
	for range 3 {
		nextIndex := (d.currentIndex + 1) % len(d.cups)
		removedCups = append(removedCups, d.cups[nextIndex])
		d.cups = append(d.cups[:nextIndex], d.cups[nextIndex+1:]...)
		if nextIndex < d.currentIndex {
			d.currentIndex--
		}
	}

	destinationCup := d.currentCup - 1
	if destinationCup < 1 {
		destinationCup = d.maxCup
	}

	removedSet := set.NewSet[int]()
	for _, cup := range removedCups {
		removedSet.Add(cup)
	}

	for removedSet.Contains(destinationCup) {
		destinationCup--
		if destinationCup < 1 {
			destinationCup = d.maxCup
		}
	}

	destinationIndex := -1
	for i, cup := range d.cups {
		if cup == destinationCup {
			destinationIndex = i + 1
			break
		}
	}

	if destinationIndex == -1 {
		panic(fmt.Sprintf("Destination cup %d not found in cups %v", destinationCup, d.cups))
	}

	newCups := make([]int, 0, len(d.cups)+3)
	newCups = append(newCups, d.cups[:destinationIndex]...)
	newCups = append(newCups, removedCups...)
	newCups = append(newCups, d.cups[destinationIndex:]...)
	d.cups = newCups

	if destinationIndex <= d.currentIndex {
		d.currentIndex += 3
	}
}

func (d *day23) Part1() string {
	for i := 0; i < 100; i++ {
		d.crabMove()
		d.currentIndex = (d.currentIndex + 1) % len(d.cups)
		d.currentCup = d.cups[d.currentIndex]
	}

	oneIndex := -1
	for i, cup := range d.cups {
		if cup == 1 {
			oneIndex = i
			break
		}
	}

	if oneIndex == -1 {
		panic("Cup 1 not found")
	}

	result := make([]string, 0, len(d.cups)-1)
	for i := 1; i < len(d.cups); i++ {
		cupIndex := (oneIndex + i) % len(d.cups)
		result = append(result, strconv.Itoa(d.cups[cupIndex]))
	}

	return strings.Join(result, "")
}

// Gemini did this :)
func (d *day23) Part2() int {
	// Create a linked list using an array where next[i] = next cup after cup i
	const totalCups = 1000000
	const moves = 10000000

	next := make([]int, totalCups+1) // 1-indexed

	// Initialize with original cups
	for i := 0; i < len(d.cups)-1; i++ {
		next[d.cups[i]] = d.cups[i+1]
	}

	// Connect original cups to new cups (10, 11, 12, ..., 1000000)
	if len(d.cups) < totalCups {
		next[d.cups[len(d.cups)-1]] = len(d.cups) + 1

		// Chain the new cups
		for i := len(d.cups) + 1; i < totalCups; i++ {
			next[i] = i + 1
		}

		// Connect last cup back to first
		next[totalCups] = d.cups[0]
	} else {
		// If we already have all cups, connect last to first
		next[d.cups[len(d.cups)-1]] = d.cups[0]
	}

	current := d.cups[0]

	for move := 0; move < moves; move++ {
		// Pick up three cups
		cup1 := next[current]
		cup2 := next[cup1]
		cup3 := next[cup2]

		// Remove the three cups from the circle
		next[current] = next[cup3]

		// Find destination
		destination := current - 1
		if destination < 1 {
			destination = totalCups
		}

		// Make sure destination is not one of the picked up cups
		for destination == cup1 || destination == cup2 || destination == cup3 {
			destination--
			if destination < 1 {
				destination = totalCups
			}
		}

		// Insert the three cups after destination
		next[cup3] = next[destination]
		next[destination] = cup1

		// Move to next current cup
		current = next[current]
	}

	// Find the two cups after cup 1 and return their product
	cup1After := next[1]
	cup2After := next[cup1After]

	return cup1After * cup2After
}

func Parse(filename string) *day23 {
	line := utils.ReadTrimmed(filename)
	cups := make([]int, 0, len(line))
	maxCup := 0

	for _, char := range line {
		cup := int(char - '0')
		cups = append(cups, cup)
		if cup > maxCup {
			maxCup = cup
		}
	}

	return &day23{
		currentCup:   cups[0],
		currentIndex: 0,
		cups:         cups,
		maxCup:       maxCup,
	}
}

func Solve(filename string) {
	part1 := Parse(filename)
	part2 := day23{
		currentCup:   part1.currentCup,
		currentIndex: part1.currentIndex,
		cups:         make([]int, len(part1.cups)),
		maxCup:       part1.maxCup,
	}
	copy(part2.cups, part1.cups)

	fmt.Println("ANSWER1: labels on the cups after cup `1`:", part1.Part1())
	fmt.Println("ANSWER2: product of the two cups after cup `1`:", part2.Part2())
}
