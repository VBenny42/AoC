package day10

import (
	"github.com/VBenny42/AoC/2023/golang/utils"
)

func (g *grid) inBounds(c utils.Coord) bool {
	return (c.X >= 0 && c.X < g.cols) &&
		(c.Y >= 0 && c.Y < g.rows)
}

// Determine the direction of the pipe at the given coordinate
func (g *grid) markPipe(c utils.Coord) {
	var (
		left      = c.Add(utils.Left)
		right     = c.Add(utils.Right)
		up        = c.Add(utils.Up)
		down      = c.Add(utils.Down)
		direction string
	)

	if g.inBounds(left) &&
		(g.grid[left.Y][left.X] == '-' ||
			g.grid[left.Y][left.X] == 'L' ||
			g.grid[left.Y][left.X] == 'F') {
		direction += "left"
	}
	if g.inBounds(right) &&
		(g.grid[right.Y][right.X] == '-' ||
			g.grid[right.Y][right.X] == 'J' ||
			g.grid[right.Y][right.X] == '7') {
		direction += "right"
	}

	if g.inBounds(up) &&
		(g.grid[up.Y][up.X] == '|' ||
			g.grid[up.Y][up.X] == '7' ||
			g.grid[up.Y][up.X] == 'F') {
		direction += "up"
	}
	if g.inBounds(down) &&
		(g.grid[down.Y][down.X] == '|' ||
			g.grid[down.Y][down.X] == 'L' ||
			g.grid[down.Y][down.X] == 'J') {
		direction += "down"
	}

	switch direction {
	case "leftright":
		g.grid[c.Y][c.X] = '-'
	case "updown":
		g.grid[c.Y][c.X] = '|'
	case "leftup":
		g.grid[c.Y][c.X] = 'J'
	case "leftdown":
		g.grid[c.Y][c.X] = '7'
	case "rightup":
		g.grid[c.Y][c.X] = 'L'
	case "rightdown":
		g.grid[c.Y][c.X] = 'F'
	}
}
