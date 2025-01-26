package day06

import (
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type race struct {
	duration int
	record   int
}

type day06 struct {
	races    []race
	longRace race
}

func (r *race) waysToBeatRecord() (ways int) {
	// Skip 0ms and skip record time,
	// boat can't move if button's held whole time
	for i := 1; i < r.record; i++ {
		distance := i * (r.duration - i)
		if distance > r.record {
			ways++
		}
	}

	return
}

func (r *race) rangeWaysToBeatRecord() int {
	// binary search to find low and high bounds where distance > record
	willBeatRecord := func(holdTime int) bool {
		return holdTime*(r.duration-holdTime) > r.record
	}

	// Find lower bound
	low, high := 1, r.duration/2
	for low < high {
		mid := low + (high-low)/2
		if willBeatRecord(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	lowerBound := low

	// Find upper bound
	low, high = r.duration/2, r.duration
	for low < high {
		mid := low + (high-low+1)/2
		if willBeatRecord(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	upperBound := low

	return upperBound - lowerBound + 1
}

func (d *day06) Part1() (product int) {
	product = 1

	for _, race := range d.races {
		product *= race.rangeWaysToBeatRecord()
	}

	return
}

func (d *day06) Part2() int {
	return d.longRace.rangeWaysToBeatRecord()
}

func Parse(filename string) *day06 {
	data := utils.ReadLines(filename)

	durations := strings.Fields(
		strings.Split(data[0], ":")[1],
	)
	records := strings.Fields(
		strings.Split(data[1], ":")[1],
	)

	races := make([]race, len(durations))
	for i := range durations {
		races[i].duration = utils.Atoi(durations[i])
		races[i].record = utils.Atoi(records[i])
	}

	longRace := race{
		duration: utils.Atoi(strings.Join(durations, "")),
		record:   utils.Atoi(strings.Join(records, "")),
	}

	return &day06{races, longRace}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: product of ways to beat record for each race:", day.Part1())
	fmt.Println("ANSWER2: ways to beat record for long race:", day.Part2())
}
