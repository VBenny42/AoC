package day25

import (
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type (
	node string
	edge struct {
		n1, n2 node
	}
)

type graph struct {
	nodes map[node]struct{}
	edges map[edge]struct{}
}

type day25 struct {
	g graph
}

func (d *day25) findAndRemoveEdge(n1, n2 node) error {
	a := edge{n1, n2}
	if _, ok := d.g.edges[a]; ok {
		delete(d.g.edges, a)
		return nil
	}
	b := edge{n2, n1}
	if _, ok := d.g.edges[b]; ok {
		delete(d.g.edges, b)
		return nil
	}
	return fmt.Errorf("edge %v not found", a)
}

// First print out the graph in graphviz format
// edges that should be removed using graphviz tool for my input:
// sqh -- jbz
// nvg -- vfj
// fch -- fvh
// Once these edges are removed, the graph should become bipartite
// Count the number of nodes in one partite set and
// multiply it by the number of nodes in the other partite set
func (d *day25) Part1() (product int, err error) {
	if err = d.findAndRemoveEdge("sqh", "jbz"); err != nil {
		return
	}
	if err = d.findAndRemoveEdge("nvg", "vfj"); err != nil {
		return
	}
	if err = d.findAndRemoveEdge("fch", "fvh"); err != nil {
		return
	}

	visited := make(map[node]bool)

	var dfs func(n node)
	dfs = func(n node) {
		visited[n] = true

		for e := range d.g.edges {
			var neighbor node
			if e.n1 == n {
				neighbor = e.n2
			} else if e.n2 == n {
				neighbor = e.n1
			} else {
				continue
			}

			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	dfs("ssd")

	product = len(visited) * (len(d.g.nodes) - len(visited))
	return
}

func Parse(filename string) *day25 {
	var (
		data  = utils.ReadLines(filename)
		nodes = make(map[node]struct{})
		edges = make(map[edge]struct{})
	)

	for _, line := range data {
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return r == ' ' || r == ':'
		})
		if len(fields) < 2 {
			panic("Node has no edges")
		}
		component := node(fields[0])
		others := fields[1:]
		nodes[component] = struct{}{}
		for _, other := range others {
			otherNode := node(other)
			nodes[otherNode] = struct{}{}
			edges[edge{component, otherNode}] = struct{}{}
		}
	}

	return &day25{
		g: graph{
			nodes: nodes,
			edges: edges,
		},
	}
}

func Solve(filename string) {
	day := Parse(filename)

	part1, err := day.Part1()
	if err != nil {
		fmt.Println("No solution found:", err)
		return
	}

	fmt.Println("ANSWER1: product of the size of the two partite sets:", part1)
}
