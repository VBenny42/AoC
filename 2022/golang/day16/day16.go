package day16

import (
	"fmt"

	"gonum.org/v1/gonum/stat/combin"
)

type state1 struct {
	valve      int
	openValves uint
	pressure   int
}

type cacheKey1 struct {
	valve      int
	openValves uint
}

func (d *day16) getMaxPressure(timeRemaining int) []state1 {
	states := []state1{{valve: d.findValve("AA"), openValves: 0, pressure: 0}}
	bestChoices := make(map[cacheKey1]int)

	for minute := 1; minute < timeRemaining; minute++ {
		var newStates []state1

		for _, state := range states {
			key := cacheKey1{valve: state.valve, openValves: state.openValves}
			if bestPressure, ok := bestChoices[key]; ok && bestPressure >= state.pressure {
				continue
			}
			bestChoices[key] = state.pressure

			valve := d.valves[state.valve]
			flowRate := valve.flowRate
			mask := uint(1) << state.valve
			if state.openValves&mask == 0 && flowRate > 0 {
				newStates = append(newStates,
					state1{
						valve:      state.valve,
						openValves: state.openValves | mask,
						pressure:   state.pressure + (timeRemaining-minute)*flowRate,
					})
			}

			for _, tunnel := range valve.tunnels {
				newStates = append(newStates,
					state1{
						valve:      tunnel,
						openValves: state.openValves,
						pressure:   state.pressure,
					})
			}

			states = newStates
		}
	}

	return states
}

func (d *day16) Part1() int {
	states := d.getMaxPressure(30)

	bestPressure := 0
	for _, state := range states {
		if state.pressure > bestPressure {
			bestPressure = state.pressure
		}
	}

	return bestPressure
}

type state2 struct {
	person     int
	elephant   int
	openValves uint
	pressure   int
}

type cacheKey2 struct {
	low, high  int
	openValves uint
}

func orderLowToHigh(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func (d *day16) Part2() int {
	states := d.getMaxPressure(26)

	gen := combin.NewCombinationGenerator(len(states), 2)
	combinations := make([]int, 2)

	bestPressure := 0

	for gen.Next() {
		gen.Combination(combinations)

		person, elephant := states[combinations[0]], states[combinations[1]]

		// check if openValves are disjoint
		bestPressure = max(bestPressure, person.pressure+elephant.pressure)
	}

	return bestPressure
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: most pressure that can be released in 30 minutes:", day.Part1())
	fmt.Println("ANSWER2: most pressure that can be released in 30 minutes with elephant:", day.Part2())
}
