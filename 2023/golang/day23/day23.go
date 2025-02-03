package day23

import (
	"fmt"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type (
	cell int
	grid [][]cell
)

const (
	empty cell = iota
	leftSlope
	rightSlope
	upSlope
	downSlope
	wall
)

type graphNode struct {
	coord utils.Coord
	dist  int
}

type graph map[utils.Coord][]graphNode

type day23 struct {
	grid          grid
	start         utils.Coord
	end           utils.Coord
	intersections []utils.Coord
	neighbors     map[utils.Coord][]utils.Coord
	graph         graph
}

func (g grid) inBounds(c utils.Coord) bool {
	return c.Y >= 0 && c.Y < len(g) && c.X >= 0 && c.X < len(g[0])
}

func (d *day23) populateNeighborsAndIntersections(part1 bool) {
	d.intersections = []utils.Coord{d.start, d.end}
	d.neighbors = make(map[utils.Coord][]utils.Coord)

	emptyCellCase := func(coord utils.Coord) {
		var count int

		for _, dir := range utils.Directions {
			next := coord.Add(dir)
			if d.grid.inBounds(next) && d.grid[next.Y][next.X] != wall {
				count++
				d.neighbors[coord] = append(d.neighbors[coord], next)
			}
		}

		if count > 2 {
			d.intersections = append(d.intersections, coord)
		}
	}

	for y, row := range d.grid {
		for x, cell := range row {
			var (
				coord = utils.Crd(x, y)
				next  utils.Coord
			)

			if part1 {
				switch cell {
				case downSlope:
					next = coord.Add(utils.Down)
				case upSlope:
					next = coord.Add(utils.Up)
				case leftSlope:
					next = coord.Add(utils.Left)
				case rightSlope:
					next = coord.Add(utils.Right)
				default:
					emptyCellCase(coord)
				}

				if d.grid.inBounds(next) {
					d.neighbors[coord] = append(d.neighbors[coord], next)
				}
			} else {
				emptyCellCase(coord)
			}
		}
	}
}

// change map to bitset later
func (d *day23) graphNode(c utils.Coord, dist int, seen map[utils.Coord]bool) graphNode {
	for _, node := range d.intersections {
		if node == c {
			return graphNode{coord: c, dist: dist}
		}
	}

	for _, neighbor := range d.neighbors[c] {
		if !seen[neighbor] {
			seen[c] = true
			return d.graphNode(neighbor, dist+1, seen)
		}
	}

	return graphNode{}
}

func (d *day23) populateGraph() {
	d.graph = make(graph)

	for _, intersection := range d.intersections {
		for _, neighbor := range d.neighbors[intersection] {
			d.graph[intersection] = append(
				d.graph[intersection],
				d.graphNode(neighbor, 1, map[utils.Coord]bool{intersection: true}),
			)
		}
	}
}

func (d *day23) bfs(start, end utils.Coord, score int, seen map[utils.Coord]bool) []int {
	if start == end {
		return []int{score}
	}

	var scores []int

	for _, node := range d.graph[start] {
		if !seen[node.coord] {
			seen[node.coord] = true
			scores = append(scores, d.bfs(node.coord, end, score+node.dist, seen)...)
			delete(seen, node.coord)
		}
	}

	return scores
}

func Parse(filename string) *day23 {
	data := utils.ReadLines(filename)
	grid := make(grid, len(data))
	for y, line := range data {
		grid[y] = make([]cell, len(line))
		for x, char := range line {
			switch char {
			case '.':
				grid[y][x] = empty
			case '<':
				grid[y][x] = leftSlope
			case '>':
				grid[y][x] = rightSlope
			case '^':
				grid[y][x] = upSlope
			case 'v':
				grid[y][x] = downSlope
			case '#':
				grid[y][x] = wall
			}
		}
	}

	start := utils.Crd(1, 0)
	end := utils.Crd(len(data[0])-2, len(data)-1)

	return &day23{
		grid:  grid,
		start: start,
		end:   end,
	}
}

func (d *day23) Part1() int {
	d.populateNeighborsAndIntersections(true)
	d.populateGraph()

	return utils.MaxSlice(d.bfs(d.start, d.end, 0, map[utils.Coord]bool{d.start: true}))
}

func (d *day23) Part2() int {
	d.populateNeighborsAndIntersections(false)
	d.populateGraph()

	return utils.MaxSlice(d.bfs(d.start, d.end, 0, map[utils.Coord]bool{d.start: true}))
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: longest path from start to end:", day.Part1())
	fmt.Println(
		"ANSWER2: longest path from start to end considering slopes as normal paths:",
		day.Part2(),
	)
}
