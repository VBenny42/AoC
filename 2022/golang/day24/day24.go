package day24

import (
	"fmt"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type blizzard struct {
	startPos  utils.Coord
	direction utils.Coord
	char      rune
}

type grid [][]rune

type day24 struct {
	startPos  utils.Coord
	endPos    utils.Coord
	blizzards []blizzard
}

func (b blizzard) calculateNextPos(steps int) (next utils.Coord) {
	next.X = (b.startPos.X + b.direction.X*steps) % totalColsGlobal
	next.Y = (b.startPos.Y + b.direction.Y*steps) % totalRowsGlobal

	next.X += totalColsGlobal
	next.X %= totalColsGlobal
	next.Y += totalRowsGlobal
	next.Y %= totalRowsGlobal

	return
}

func (d *day24) getRoomState(steps int, memo map[int]grid) grid {
	if grid, ok := memo[steps]; ok {
		return grid
	}

	grid := make(grid, totalRowsGlobal)
	for i := range grid {
		grid[i] = make([]rune, totalColsGlobal)
	}

	for _, b := range d.blizzards {
		nextPos := b.calculateNextPos(steps)
		grid[nextPos.Y][nextPos.X] = b.char
	}

	for y := 0; y < totalRowsGlobal; y++ {
		for x := 0; x < totalColsGlobal; x++ {
			if grid[y][x] == 0 {
				grid[y][x] = '.'
			}
		}
	}

	memo[steps] = grid

	return grid
}

func (d *day24) bfs(stepsElapsed int) int {
	memo := make(map[int]grid)

	type node struct {
		steps int
		pos   utils.Coord
	}

	queue := make([]node, 0, 1400)
	queue = append(queue, node{steps: stepsElapsed, pos: d.startPos})

	seen := make(map[node]bool)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		grid := d.getRoomState(curr.steps+1, memo)

		for _, direction := range utils.Directions {
			nextPos := curr.pos.Add(direction)
			if nextPos == d.startPos {
				continue
			}

			if nextPos != d.endPos &&
				(nextPos.X < 0 || nextPos.X >= totalColsGlobal ||
					nextPos.Y < 0 || nextPos.Y >= totalRowsGlobal) {
				continue
			}

			h := node{curr.steps + 1, nextPos}
			if seen[h] {
				continue
			}
			seen[h] = true

			if nextPos == d.endPos {
				return curr.steps + 1
			}

			if grid[nextPos.Y][nextPos.X] != '.' {
				continue
			}

			queue = append(queue, node{steps: curr.steps + 1, pos: nextPos})
		}

		if curr.pos == d.startPos ||
			grid[curr.pos.Y][curr.pos.X] == '.' {
			queue = append(queue, node{steps: curr.steps + 1, pos: curr.pos})
		}
	}

	panic("No path found")
}

func (d *day24) Part1And2() (int, int) {
	firstLeg := d.bfs(0)
	actualStart, actualEnd := d.startPos, d.endPos
	d.startPos, d.endPos = actualEnd, actualStart
	secondLeg := d.bfs(firstLeg)
	d.startPos, d.endPos = actualStart, actualEnd
	return firstLeg, d.bfs(secondLeg)
}

var totalRowsGlobal, totalColsGlobal int

func Parse(filename string) *day24 {
	data := utils.ReadLines(filename)

	var start, end utils.Coord
	var blizzards []blizzard

	for x := range data {
		if data[0][x] == '.' {
			start = utils.Coord{X: x - 1, Y: -1}
			break
		}
	}

	totalRows := len(data) - 2
	totalCols := len(data[0]) - 2

	for x := totalCols + 1; x >= 0; x-- {
		if data[totalRows+1][x] == '.' {
			end = utils.Coord{X: x - 1, Y: totalRows}
			break
		}
	}

	for y := 1; y < len(data)-1; y++ {
		for x := 1; x < len(data[y])-1; x++ {
			startPos := utils.Coord{X: x - 1, Y: y - 1}
			switch data[y][x] {
			case '>':
				blizzards = append(blizzards, blizzard{
					startPos:  startPos,
					direction: utils.Right,
					char:      '>',
				})
			case '<':
				blizzards = append(blizzards, blizzard{
					startPos:  startPos,
					direction: utils.Left,
					char:      '<',
				})
			case '^':
				blizzards = append(blizzards, blizzard{
					startPos:  startPos,
					direction: utils.Up,
					char:      '^',
				})
			case 'v':
				blizzards = append(blizzards, blizzard{
					startPos:  startPos,
					direction: utils.Down,
					char:      'v',
				})
			case '.', '#':
			// Do nothing
			default:
				panic("Unhandled character")
			}
		}
	}

	totalRowsGlobal = totalRows
	totalColsGlobal = totalCols

	return &day24{
		startPos:  start,
		endPos:    end,
		blizzards: blizzards,
	}
}

func Solve(filename string) {
	day := Parse(filename)

	part1, part2 := day.Part1And2()

	fmt.Println("ANSWER1: fewest minutes required to avoid the blizzards and reach the end:", part1)
	fmt.Println("ANSWER2: fewest minutes required to go back to start then end:", part2)
}
