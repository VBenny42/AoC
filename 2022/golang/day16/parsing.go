package day16

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type valve struct {
	index    int
	tunnels  []int
	flowRate int
}

func (v valve) String() string {
	tunnels := make([]string, len(v.tunnels))
	for i, t := range v.tunnels {
		tunnels[i] = fmt.Sprintf("%d", t)
	}
	return fmt.Sprintf("Valve %d: %s", v.index, strings.Join(tunnels, ", "))
}

type day16 struct {
	names     []string
	indices   map[string]int
	valves    map[int]valve
	distances map[int]map[int]int
}

func (d *day16) findValve(name string) int {
	if index, ok := d.indices[name]; ok {
		return index
	}
	n := len(d.names)
	d.names = append(d.names, name)
	d.indices[name] = n
	return n
}

func (d *day16) addValve(v valve) {
	d.valves[v.index] = v
}

func Parse(filename string) *day16 {
	data := utils.ReadLines(filename)

	day := day16{
		names:   make([]string, 0, len(data)),
		indices: make(map[string]int, len(data)),
		valves:  make(map[int]valve, len(data)),
	}

	pattern := regexp.MustCompile(`^Valve ([A-Z]+) has flow rate=(\d+); tunnels? leads? to valves? (.+)+$`)

	for _, line := range data {
		matches := pattern.FindStringSubmatch(line)
		if len(matches) == 0 {
			panic("Invalid input")
		}
		var v valve
		v.index = day.findValve(matches[1])
		v.flowRate = utils.Must(strconv.Atoi(matches[2]))
		for _, tunnel := range strings.Split(matches[3], ", ") {
			v.tunnels = append(v.tunnels, day.findValve(tunnel))
		}
		day.addValve(v)
	}

	return &day
}
