package day13

import (
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2020/golang/utils"
)

type busInfo struct {
	id     int
	offset int
}

type day13 struct {
	earliestDeparture int
	busIDs            []int
	busSchedule       []busInfo
}

func (d *day13) Part1() int {
	var (
		minWaitTime = int(^uint(0) >> 1) // Max int
		bestBusID   int
	)

	for _, id := range d.busIDs {
		waitTime := id - (d.earliestDeparture % id)
		if waitTime < minWaitTime {
			minWaitTime = waitTime
			bestBusID = id
		}
	}

	return bestBusID * minWaitTime
}

// Uses the Chinese Remainder Theorem to find the earliest timestamp
func (d *day13) Part2() int {
	if len(d.busSchedule) == 0 {
		return 0
	}

	// Start with the first bus
	time := 0
	step := d.busSchedule[0].id

	// Process each subsequent bus
	for i := 1; i < len(d.busSchedule); i++ {
		bus := d.busSchedule[i]

		// Find a time where this bus constraint is satisfied
		for (time+bus.offset)%bus.id != 0 {
			time += step
		}

		// Update step to be LCM of current step and this bus ID
		// Since all bus IDs are prime in AoC, LCM is just the product
		step *= bus.id
	}

	return time
}

func Parse(filename string) *day13 {
	var (
		lines = utils.ReadLines(filename)
		buses = strings.Split(lines[1], ",")
		day   day13
	)

	day.earliestDeparture = utils.Atoi(lines[0])

	for _, id := range buses {
		if id != "x" {
			day.busIDs = append(day.busIDs, utils.Atoi(id))
		}
	}

	for i, id := range buses {
		if id != "x" {
			day.busSchedule = append(day.busSchedule, busInfo{
				id:     utils.Atoi(id),
				offset: i,
			})
		}
	}

	return &day
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: product of earliest bus ID and wait time:", day.Part1())
	fmt.Println("ANSWER2: earliest timestamp where all bus IDs depart at their respective offsets:", day.Part2())
}
