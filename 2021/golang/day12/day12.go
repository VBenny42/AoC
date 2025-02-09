package day12

import (
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

type day12 struct {
	graph map[string]map[string]struct{}
}

func (d *day12) Part1() int {
	var traverse func(string, map[string]bool) int
	traverse = func(curr string, visited map[string]bool) (endPaths int) {
		if curr == "end" {
			return 1
		}

		for cave := range d.graph[curr] {
			if visited[cave] && strings.ToUpper(cave) != cave {
				continue
			}

			visited[cave] = true

			endPaths += traverse(cave, visited)

			visited[cave] = false
		}

		return
	}

	return traverse("start", map[string]bool{"start": true})
}

func (d *day12) Part2() int {
	var traverse func(string, map[string]int, bool) int
	traverse = func(curr string, visited map[string]int, visitedTwice bool) (endPaths int) {
		if curr == "end" {
			return 1
		}

		visited[curr]++

		for cave := range d.graph[curr] {
			if cave == "start" {
				continue
			}

			if strings.ToUpper(cave) != cave && visited[cave] > 0 {
				if visitedTwice {
					continue
				}

				visitedTwice = true
			}

			endPaths += traverse(cave, visited, visitedTwice)

			visited[cave]--
			if strings.ToUpper(cave) != cave && visited[cave] == 1 {
				visitedTwice = false
			}
		}

		return
	}

	return traverse("start", map[string]int{"start": 0}, false)
}

func Parse(filename string) *day12 {
	var (
		data  = utils.ReadLines(filename)
		graph = make(map[string]map[string]struct{})
	)

	for _, line := range data {
		left, right, found := strings.Cut(line, "-")
		if !found {
			panic("Invalid input, no '-' found")
		}

		if _, ok := graph[left]; !ok {
			graph[left] = make(map[string]struct{})
		}
		if _, ok := graph[right]; !ok {
			graph[right] = make(map[string]struct{})
		}

		graph[left][right] = struct{}{}
		graph[right][left] = struct{}{}
	}

	return &day12{graph: graph}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println(
		"ANSWER1: paths through cave system that visit small caves at least once:",
		day.Part1(),
	)
	fmt.Println(
		"ANSWER2: paths through cave system that visit one small cave twice:",
		day.Part2(),
	)
}
