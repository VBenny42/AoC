package day11

import (
	"fmt"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

func (g *grid) inBounds(c utils.Coord) bool {
	return (c.X >= 0 && c.X < len((*g)[0])) && (c.Y >= 0 && c.Y < len(*g))
}

func (g *grid) expand() grid {
	newGrid := make(grid, len(*g))
	for i := 0; i < len(*g); i++ {
		newGrid[i] = make([]cell, len((*g)[0]))
		copy(newGrid[i], (*g)[i])
	}

	for y := 0; y < len(*g); y++ {
		shouldExpand := true
		for _, cell := range (*g)[y] {
			if cell == galaxy {
				shouldExpand = false
				break
			}
		}
		if shouldExpand {
			for x := 0; x < len((*g)[0]); x++ {
				newGrid[y][x] = expansion
			}
		}
	}

	for x := 0; x < len((*g)[0]); x++ {
		shouldExpand := true
		for y := 0; y < len(*g); y++ {
			if (*g)[y][x] == galaxy {
				shouldExpand = false
				break
			}
		}
		if shouldExpand {
			for y := 0; y < len(*g); y++ {
				newGrid[y][x] = expansion
			}
		}
	}

	return newGrid
}

type stepWithExpansion struct {
	actual     int
	expansions int
}

func (g *grid) bfs(start, end utils.Coord, memo map[utils.Coord]map[utils.Coord]stepWithExpansion) (stepWithExpansion, error) {
	if _, ok := memo[start][end]; ok {
		return memo[start][end], nil
	}
	if _, ok := memo[end][start]; ok {
		return memo[end][start], nil
	}

	type node struct {
		coord utils.Coord
		steps stepWithExpansion
	}

	var (
		queue = []node{{start, stepWithExpansion{0, 0}}}
		seen  = map[utils.Coord]struct{}{}
	)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if _, ok := seen[curr.coord]; ok {
			continue
		}
		seen[curr.coord] = struct{}{}

		if (*g)[curr.coord.Y][curr.coord.X] == galaxy {
			memo[start][curr.coord] = curr.steps
			memo[curr.coord][start] = curr.steps
		}

		for _, dir := range utils.Directions {
			next := curr.coord.Add(dir)
			if g.inBounds(next) {
				stepsVal := curr.steps
				if (*g)[next.Y][next.X] > galaxy {
					stepsVal = stepWithExpansion{stepsVal.actual, stepsVal.expansions + 1}
				} else {
					stepsVal = stepWithExpansion{stepsVal.actual + 1, stepsVal.expansions}
				}
				queue = append(queue, node{next, stepsVal})
			}
		}
	}

	if _, ok := memo[start][end]; !ok {
		return stepWithExpansion{0, 0}, fmt.Errorf("no path found between %v and %v", start, end)
	}

	return memo[start][end], nil
}
