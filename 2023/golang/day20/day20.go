package day20

import (
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type modType int

const (
	flipFlop modType = iota
	conjunction
	broadcast
)

type module struct {
	modType      modType
	name         string
	destinations []string
	state        bool
	inputs       map[string]bool
}

type pulse struct {
	from, to string
	isHigh   bool
}

type day20 struct {
	modules map[string]*module
}

func (d *day20) resetStates() {
	for _, mod := range d.modules {
		mod.state = false
		for k := range mod.inputs {
			mod.inputs[k] = false
		}
	}
}

func (d *day20) Part1() int {
	d.resetStates()
	var lowCount, highCount int
	var queue []pulse

	for range 1000 {
		queue = queue[:0]
		queue = append(queue, pulse{from: "button", to: "broadcaster", isHigh: false})
		lowCount++

		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			mod, ok := d.modules[p.to]
			if !ok {
				continue
			}

			switch mod.modType {
			case broadcast:
				for _, dest := range mod.destinations {
					if p.isHigh {
						highCount++
					} else {
						lowCount++
					}
					queue = append(queue, pulse{from: mod.name, to: dest, isHigh: p.isHigh})
				}
			case flipFlop:
				if !p.isHigh {
					mod.state = !mod.state
					for _, dest := range mod.destinations {
						if mod.state {
							highCount++
						} else {
							lowCount++
						}
						queue = append(queue, pulse{from: mod.name, to: dest, isHigh: mod.state})
					}
				}
			case conjunction:
				mod.inputs[p.from] = p.isHigh
				allHigh := true
				for _, input := range mod.inputs {
					if !input {
						allHigh = false
						break
					}
				}
				sendHigh := !allHigh
				for _, dest := range mod.destinations {
					if sendHigh {
						highCount++
					} else {
						lowCount++
					}
					queue = append(queue, pulse{from: mod.name, to: dest, isHigh: sendHigh})
				}
			}
		}
	}

	return lowCount * highCount
}

func (d *day20) Part2() int {
	d.resetStates()

	var xModule *module
	for _, mod := range d.modules {
		for _, dest := range mod.destinations {
			if dest == "rx" {
				xModule = mod
				break
			}
		}
		if xModule != nil {
			break
		}
	}
	if xModule == nil || xModule.modType != conjunction {
		fmt.Println("rx is not connected to a conjunction module")
		return 0
	}

	sources := make([]string, 0, len(xModule.inputs))
	for source := range xModule.inputs {
		sources = append(sources, source)
	}

	highPulseTimes := make(map[string][]int)
	for _, source := range sources {
		highPulseTimes[source] = make([]int, 0, 2)
	}

	var (
		pressCount int
		queue      []pulse
	)

	for {
		pressCount++
		queue = queue[:0]
		queue = append(queue, pulse{from: "button", to: "broadcaster", isHigh: false})

		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			if p.to == xModule.name && p.isHigh {
				for _, source := range sources {
					if p.from == source {
						highPulseTimes[source] = append(highPulseTimes[source], pressCount)
					}
				}
			}

			mod, ok := d.modules[p.to]
			if !ok {
				continue
			}

			switch mod.modType {
			case broadcast:
				for _, dest := range mod.destinations {
					queue = append(queue, pulse{from: mod.name, to: dest, isHigh: p.isHigh})
				}
			case flipFlop:
				if !p.isHigh {
					mod.state = !mod.state
					for _, dest := range mod.destinations {
						queue = append(queue, pulse{from: mod.name, to: dest, isHigh: mod.state})
					}
				}
			case conjunction:
				mod.inputs[p.from] = p.isHigh
				allHigh := true
				for _, input := range mod.inputs {
					if !input {
						allHigh = false
						break
					}
				}
				sendHigh := !allHigh
				for _, dest := range mod.destinations {
					queue = append(queue, pulse{from: mod.name, to: dest, isHigh: sendHigh})
				}
			}
		}

		allEnough := true
		for _, source := range sources {
			if len(highPulseTimes[source]) < 2 {
				allEnough = false
				break
			}
		}
		if allEnough {
			break
		}
	}

	cycles := make([]int, 0, len(sources))
	for _, source := range sources {
		times := highPulseTimes[source]
		cycles = append(cycles, times[1]-times[0])
	}

	// Cycle numbers seem to be all prime numbers,
	// no need to find LCM
	product := 1
	for _, cycle := range cycles {
		product *= cycle
	}

	return product
}

func Parse(filename string) *day20 {
	data := utils.ReadLines(filename)
	modules := make(map[string]*module, len(data))

	for _, line := range data {
		parts := strings.Split(line, " -> ")
		var name string
		var mt modType

		if parts[0] == "broadcaster" {
			name = parts[0]
			mt = broadcast
		} else {
			switch parts[0][0] {
			case '%':
				mt = flipFlop
			case '&':
				mt = conjunction
			default:
				panic("invalid module type")
			}
			name = parts[0][1:]
		}

		destinations := strings.Split(parts[1], ", ")
		mod := &module{
			modType:      mt,
			name:         name,
			destinations: destinations,
		}

		if mt == conjunction {
			mod.inputs = make(map[string]bool)
		}

		modules[name] = mod
	}

	for _, mod := range modules {
		for _, dest := range mod.destinations {
			destMod, exists := modules[dest]
			if exists && destMod.modType == conjunction {
				destMod.inputs[mod.name] = false
			}
		}
	}

	return &day20{modules: modules}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: product of low pulses and high pulses:", day.Part1())
	fmt.Println("ANSWER2: fewest button presses to deliver low to rx:", day.Part2())
}
