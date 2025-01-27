package day11

import (
	// "slices"
	"fmt"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

func (g *grid) inBounds(c utils.Coord) bool {
	return (c.X >= 0 && c.X < len((*g)[0])) && (c.Y >= 0 && c.Y < len(*g))
}

func (g *grid) expand() grid {
	// Expand rows first
	var expandedRowGrid grid
	for i := 0; i < len(*g); i++ {
		shouldExpand := true
		for _, cell := range (*g)[i] {
			if cell == galaxy {
				shouldExpand = false
				break
			}
		}
		expandedRowGrid = append(expandedRowGrid, (*g)[i])
		if shouldExpand {
			// Insert duplicate row
			newRow := make([]int, len((*g)[i]))
			copy(newRow, (*g)[i])
			expandedRowGrid = append(expandedRowGrid, newRow)
		}
	}

	// Column expansion
	var expandedColGrid grid
	for i := 0; i < len(expandedRowGrid); i++ {
		expandedColGrid = append(expandedColGrid, make([]int, len(expandedRowGrid[0])))
		copy(expandedColGrid[i], expandedRowGrid[i])
	}

	// Find columns that need expansion
	colsToExpand := []int{}
	for col := 0; col < len(expandedRowGrid[0]); col++ {
		shouldExpand := true
		for row := 0; row < len(expandedRowGrid); row++ {
			if expandedRowGrid[row][col] == galaxy {
				shouldExpand = false
				break
			}
		}
		if shouldExpand {
			colsToExpand = append(colsToExpand, col)
		}
	}

	// Expand columns
	offset := 0
	for _, col := range colsToExpand {
		col += offset // Adjust for previously inserted columns
		// For each row in the grid
		for i := 0; i < len(expandedColGrid); i++ {
			// Insert duplicate column
			oldRow := expandedColGrid[i]
			newRow := make([]int, len(oldRow)+1)
			copy(newRow, oldRow[:col+1])
			newRow[col+1] = oldRow[col]
			copy(newRow[col+2:], oldRow[col+1:])
			expandedColGrid[i] = newRow
		}
		offset++
	}

	return expandedColGrid
}

func (g *grid) bfs(start, end utils.Coord, memo map[utils.Coord]map[utils.Coord]int) (int, error) {
	if memo[start][end] != 0 {
		return memo[start][end], nil
	}
	if memo[end][start] != 0 {
		return memo[end][start], nil
	}

	type node struct {
		coord utils.Coord
		steps int
	}

	var (
		queue = []node{{start, 0}}
		seen  = map[utils.Coord]struct{}{}
	)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// fmt.Printf("Visiting (%d,%d) at step %d\n", curr.coord.X, curr.coord.Y, curr.steps)
		// (*g)[curr.coord.Y][curr.coord.X] = 'X'

		if _, ok := seen[curr.coord]; ok {
			continue
		}
		seen[curr.coord] = struct{}{}

		if (*g)[curr.coord.Y][curr.coord.X] == galaxy {
			memo[start][curr.coord] = curr.steps
		}

		if curr.coord == end {
			memo[start][end] = curr.steps
		}

		for _, dir := range utils.Directions {
			next := curr.coord.Add(dir)
			if g.inBounds(next) {
				queue = append(queue, node{next, curr.steps + 1})
			}
		}
	}

	if _, ok := memo[start][end]; !ok {
		return 0, fmt.Errorf("no path found between %v and %v", start, end)
	}

	return memo[start][end], nil
}
