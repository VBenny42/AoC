package day12

import (
	"fmt"
	"sync"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type grid [][]int

type day12 struct {
	grid        grid
	start, end  utils.Coord
	startPoints []utils.Coord
}

func (g grid) String() string {
	s := ""
	for _, row := range g {
		for _, cell := range row {
			s += string(rune(cell))
		}
		s += "\n"
	}
	return s
}

func (g *grid) isValid(c utils.Coord) bool {
	return c.Y >= 0 && c.Y < len(*g) && c.X >= 0 && c.X < len((*g)[0])
}

type node struct {
	c       utils.Coord
	pathLen int
}

func (d *day12) bfs(start utils.Coord) (int, error) {
	queue := make([]node, 1, len(d.grid)*len(d.grid[0]))
	queue[0] = node{start, 0}

	visited := make([][]bool, len(d.grid))
	for i := range visited {
		visited[i] = make([]bool, len(d.grid[0]))
	}
	visited[d.start.Y][d.start.X] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		currentValue := d.grid[current.c.Y][current.c.X]

		if current.c == d.end {
			return current.pathLen, nil
		}

		for _, dir := range utils.Directions {
			next := current.c.Add(dir)
			if !d.grid.isValid(next) {
				continue
			}

			nextValue := d.grid[next.Y][next.X]
			if !visited[next.Y][next.X] && (nextValue-currentValue <= 1) {
				visited[next.Y][next.X] = true
				queue = append(queue, node{next, current.pathLen + 1})
			}
		}
	}
	return -1, fmt.Errorf("no path found")
}

func (d *day12) Part1And2() (int, int) {
	// Guaranteed to have a path from marked start to end
	actualStart, _ := d.bfs(d.start)
	minStart := actualStart

	var wg sync.WaitGroup
	wg.Add(len(d.startPoints))
	ch := make(chan int, len(d.startPoints))

	for _, start := range d.startPoints {
		go func(start utils.Coord) {
			defer wg.Done()

			pathLen, err := d.bfs(start)
			if err != nil {
				return
			}

			ch <- pathLen
		}(start)
	}

	wg.Wait()
	close(ch)

	for pathLen := range ch {
		minStart = min(minStart, pathLen)
	}

	return actualStart, minStart
}

func Parse(filename string) *day12 {
	data := utils.ReadLines(filename)

	grid := make(grid, len(data))
	var start, end utils.Coord
	startPoints := make([]utils.Coord, 0)

	for y, line := range data {
		grid[y] = make([]int, len(line))
		for x, c := range line {
			grid[y][x] = int(c)
			if c == 'a' {
				startPoints = append(startPoints, utils.Coord{X: x, Y: y})
			} else if c == 'S' {
				start = utils.Coord{X: x, Y: y}
				grid[y][x] = int('a')
			} else if c == 'E' {
				end = utils.Coord{X: x, Y: y}
				grid[y][x] = int('z')
			}
		}
	}

	return &day12{grid, start, end, startPoints}
}

func Solve(filename string) {
	day := Parse(filename)

	actualStart, minStart := day.Part1And2()

	fmt.Println("ANSWER1: shortest path length from S to E is", actualStart)
	fmt.Println("ANSWER2: shortest path lengths from lowest point to E are:", minStart)
}
