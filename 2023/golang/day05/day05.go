package day05

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type mapRange struct {
	destStart int
	srcStart  int
	length    int
}

type mapIdx int

const (
	seedToSoil mapIdx = iota
	soilToFertilizer
	fertilizerToWater
	waterToLight
	lightToTemp
	tempToHumidity
	humidityToLocation
)

type day05 struct {
	seeds []int
	maps  [7][]mapRange
}

func getMappedValue(maps []mapRange, value int) int {
	for _, m := range maps {
		if value >= m.srcStart && value < m.srcStart+m.length {
			return m.destStart + value - m.srcStart
		}
	}

	return value
}

func (d *day05) getLocationFromSeed(seed int) int {
	value := seed
	value = getMappedValue(d.maps[seedToSoil], value)
	value = getMappedValue(d.maps[soilToFertilizer], value)
	value = getMappedValue(d.maps[fertilizerToWater], value)
	value = getMappedValue(d.maps[waterToLight], value)
	value = getMappedValue(d.maps[lightToTemp], value)
	value = getMappedValue(d.maps[tempToHumidity], value)
	value = getMappedValue(d.maps[humidityToLocation], value)

	return value
}

func (d *day05) Part1() int {
	minLocation := d.getLocationFromSeed(d.seeds[0])
	for _, seed := range d.seeds[1:] {
		minLocation = min(minLocation, d.getLocationFromSeed(seed))
	}

	return minLocation
}

// Can't do brute force, need to use range info somehow
func (d *day05) Part2Other() int {
	minLocation := d.getLocationFromSeed(d.seeds[0])
	for i := 0; i < len(d.seeds); i += 2 {
		for j := d.seeds[i]; j < d.seeds[i]+d.seeds[i+1]; j++ {
			minLocation = min(minLocation, d.getLocationFromSeed(j))
		}
	}

	return minLocation
}

// Borrowed from
// https://github.com/umuterenornek/adventofcode-rust/blob/2023/src/bin/05.rs
func (d *day05) Part2() int {
	var (
		locationRanges ranges
		rangesToAdd    ranges
	)

	for i := 0; i < len(d.seeds); i += 2 {
		locationRanges.union(ranges{[]rangeType{{
			d.seeds[i],
			d.seeds[i] + d.seeds[i+1],
		}}})
	}

	for _, maps := range d.maps {
		for _, m := range maps {
			var (
				diff     = m.destStart - m.srcStart
				srcRange = ranges{[]rangeType{{
					m.srcStart,
					m.srcStart + m.length,
				}}}
				rangeDiff = srcRange.intersection(locationRanges)
			)

			for _, r := range rangeDiff.ranges {
				rangesToAdd.union(ranges{[]rangeType{{
					r.start + diff,
					r.end + diff,
				}}})
			}

			locationRanges = locationRanges.difference(rangeDiff)
		}

		locationRanges.union(rangesToAdd)
		rangesToAdd = ranges{[]rangeType{}}
	}

	return locationRanges.ranges[0].start
}

func Parse(filename string) *day05 {
	data := utils.ReadLines(filename)

	var (
		seedStrings = strings.Fields(
			strings.Split(data[0], ": ")[1],
		)
		seeds = make([]int, len(seedStrings))
	)

	for i, seedString := range seedStrings {
		seeds[i] = utils.Atoi(seedString)
	}

	var (
		maps = [7][]mapRange{}
		idx  int
	)

	for _, line := range data[2:] {
		if line == "" {
			idx++
			continue
		}

		if unicode.IsDigit(rune(line[0])) {
			parts := strings.Fields(line)
			var (
				destStart = utils.Atoi(parts[0])
				srcStart  = utils.Atoi(parts[1])
				length    = utils.Atoi(parts[2])
			)

			maps[idx] = append(maps[idx], mapRange{destStart, srcStart, length})
		}

	}

	return &day05{seeds, maps}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: lowest location number:", day.Part1())
	fmt.Println("ANSWER2: lowest location number with range:", day.Part2())
}
