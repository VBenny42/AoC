package day06

import (
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2019/golang/utils"
)

type day06 struct {
	directed   map[string][]string // parent -> children (for Part 1)
	undirected map[string][]string // bidirectional (for Part 2)
}

func (d *day06) Part1() (count int) {
	var countOrbits func(node string, depth int)
	countOrbits = func(node string, depth int) {
		count += depth
		for _, child := range d.directed[node] {
			countOrbits(child, depth+1)
		}
	}

	countOrbits("COM", 0)

	return
}

func (d *day06) bfs(start, goal string) int {
	type nodeLevel struct {
		node  string
		level int
	}

	var (
		visited = make(map[string]bool)
		queue   = []nodeLevel{{node: start, level: 0}}
	)

	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.node == goal {
			return current.level
		}

		// Use undirected graph for BFS
		for _, neighbor := range d.undirected[current.node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, nodeLevel{node: neighbor, level: current.level + 1})
			}
		}
	}
	return -1 // Goal not reachable
}

func (d *day06) Part2() int {
	var (
		youOrbiting = ""
		sanOrbiting = ""
	)

	for center, orbiters := range d.directed {
		for _, orbiter := range orbiters {
			if orbiter == "YOU" {
				youOrbiting = center
			}
			if orbiter == "SAN" {
				sanOrbiting = center
			}
		}
	}

	if youOrbiting == "" {
		panic("Cannot find what YOU is orbiting")
	}
	if sanOrbiting == "" {
		panic("Cannot find what SAN is orbiting")
	}

	// BFS from what YOU is orbiting to what SAN is orbiting
	distance := d.bfs(youOrbiting, sanOrbiting)

	return distance
}

func Parse(filename string) *day06 {
	directed := make(map[string][]string)
	undirected := make(map[string][]string)

	for _, line := range utils.ReadLines(filename) {
		objects := strings.Split(line, ")")
		if len(objects) != 2 {
			panic("Invalid input line: " + line)
		}
		center := objects[0]
		orbiter := objects[1]

		// Directed graph: parent -> child (for Part 1)
		if _, exists := directed[center]; !exists {
			directed[center] = []string{}
		}
		directed[center] = append(directed[center], orbiter)

		// Initialize orbiter in directed graph even if it has no children
		if _, exists := directed[orbiter]; !exists {
			directed[orbiter] = []string{}
		}

		// Undirected graph: bidirectional edges (for Part 2)
		if _, exists := undirected[center]; !exists {
			undirected[center] = []string{}
		}
		undirected[center] = append(undirected[center], orbiter)

		if _, exists := undirected[orbiter]; !exists {
			undirected[orbiter] = []string{}
		}
		undirected[orbiter] = append(undirected[orbiter], center)
	}

	return &day06{directed: directed, undirected: undirected}
}

func Solve(filename string) {
	day := Parse(filename)
	fmt.Println("ANSWER1: total number of direct and indirect orbits:", day.Part1())
	fmt.Println("ANSWER2: minimum orbital transfers required:", day.Part2())
}
