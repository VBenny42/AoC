package day22

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type (
	cell byte
	grid [][]cell
)

type (
	movement struct {
		amount   *int
		rotation *rune
	}
	state struct {
		position  utils.Coord
		direction direction
	}
)

type (
	bounds struct {
		start int
		end   int
	}
	gridBounds   [2]map[int]bounds
	squareBounds [2]bounds
)

type day22 struct {
	grid        grid
	start       state
	movements   []movement
	gridBounds  gridBounds
	faces       map[face]squareBounds
	transitions map[face]transition
}

const (
	row = iota
	col
)

const (
	wall = iota + 4
	movable
	last
)

type direction int

const (
	right direction = iota
	down
	left
	up
)

type face int

const (
	one face = iota
	two
	three
	four
	five
	six
)

var rotateRight = map[direction]direction{
	up:    right,
	right: down,
	down:  left,
	left:  up,
}

var rotateLeft = map[direction]direction{
	up:    left,
	left:  down,
	down:  right,
	right: up,
}

var moveDirection = map[direction]utils.Coord{
	up:    utils.Up,
	down:  utils.Down,
	left:  utils.Left,
	right: utils.Right,
}

func (d *day22) move(curr state, amount int) (newPos utils.Coord, hitWall bool) {
	newPos = curr.position
	for i := 0; i < amount; i++ {
		nextPos := newPos.Add(moveDirection[curr.direction])

		if nextPos.Y < 0 || nextPos.Y >= len(d.grid) ||
			nextPos.X < 0 || nextPos.X >= len(d.grid[nextPos.Y]) ||
			d.grid[nextPos.Y][nextPos.X]&(1<<movable) == 0 {
			switch curr.direction {
			case right:
				nextPos.X = d.gridBounds[row][nextPos.Y].start
			case left:
				nextPos.X = d.gridBounds[row][nextPos.Y].end
			case down:
				nextPos.Y = d.gridBounds[col][nextPos.X].start
			case up:
				nextPos.Y = d.gridBounds[col][nextPos.X].end
			}
		}

		if d.grid[nextPos.Y][nextPos.X]&(1<<wall) != 0 {
			return newPos, true
		}

		d.grid[newPos.Y][newPos.X] |= 1 << curr.direction
		newPos = nextPos
	}
	return
}

func (d *day22) move3d(curr state, currFace face, amount int) (
	newState state,
	newFace face,
	hitWall bool,
) {
	newState = curr
	newFace = currFace

	var (
		nextState = newState
		nextFace  = newFace
	)

	for i := 0; i < amount; i++ {
		nextState.position = newState.position.Add(moveDirection[newState.direction])

		if nextState.position.Y > d.faces[newFace][col].end ||
			nextState.position.Y < d.faces[newFace][col].start ||
			nextState.position.X > d.faces[newFace][row].end ||
			nextState.position.X < d.faces[newFace][row].start {
			// Wrapping around to next face
			nextFace = d.transitions[newFace][newState.direction].face
			nextState.direction = d.transitions[newFace][newState.direction].direction
			nextState.position = d.transitions[newFace][newState.direction].
				newPos(newState.position)
		}

		if d.grid[nextState.position.Y][nextState.position.X]&(1<<wall) != 0 {
			return newState, newFace, true
		}

		d.grid[newState.position.Y][newState.position.X] |= 1 << newState.direction
		newState = nextState
		newFace = nextFace
	}

	return
}

func (s *state) password() (sum int) {
	sum += 1000 * (s.position.Y + 1)
	sum += 4 * (s.position.X + 1)
	sum += int(s.direction)

	return
}

func (d *day22) Part1() int {
	var (
		curr    = d.start
		hitWall bool
	)

	for _, move := range d.movements {
		if move.amount != nil {
			curr.position, hitWall = d.move(curr, *move.amount)
			if hitWall {
				continue
			}
		}
		if move.rotation != nil {
			switch *move.rotation {
			case 'R':
				curr.direction = rotateRight[curr.direction]
			case 'L':
				curr.direction = rotateLeft[curr.direction]
			}
		}
	}

	d.grid[curr.position.Y][curr.position.X] |= 1 << curr.direction
	d.grid[curr.position.Y][curr.position.X] |= 1 << last

	return curr.password()
}

func (d *day22) Part2() int {
	var (
		curr     = d.start
		currFace = one
		hitWall  bool
	)

	for _, move := range d.movements {
		if move.amount != nil {
			curr, currFace, hitWall = d.move3d(curr, currFace, *move.amount)
			if hitWall {
				continue
			}
		}
		if move.rotation != nil {
			switch *move.rotation {
			case 'R':
				curr.direction = rotateRight[curr.direction]
			case 'L':
				curr.direction = rotateLeft[curr.direction]
			}
		}
	}

	d.grid[curr.position.Y][curr.position.X] |= 1 << curr.direction
	d.grid[curr.position.Y][curr.position.X] |= 1 << last

	return curr.password()
}

func Parse(filename string, sideLength int) *day22 {
	var (
		data      = utils.ReadTrimmed(filename)
		split     = strings.Split(data, "\n\n")
		gridLines = strings.Split(split[0], "\n")
	)

	var (
		start    *state
		maxWidth int
	)

	for _, line := range gridLines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	grid := make(grid, len(gridLines))
	for y, line := range gridLines {
		grid[y] = make([]cell, maxWidth)
		for x, char := range line {
			if char == '.' {
				if start == nil {
					start = &state{utils.Coord{X: x, Y: y}, right}
					grid[y][x] |= 1 << right
				}
				grid[y][x] |= 1 << movable
			} else if char == '#' {
				grid[y][x] |= 1 << wall
				grid[y][x] |= 1 << movable
			}
		}
	}

	var (
		rowBounds = make(map[int]bounds, len(grid))
		colBounds = make(map[int]bounds, maxWidth)
	)

	for y, row := range grid {
		for x, cell := range row {
			if cell&(1<<movable) != 0 {
				if yBound, ok := rowBounds[y]; !ok {
					rowBounds[y] = bounds{start: x, end: x}
				} else {
					yBound.end = x
					rowBounds[y] = yBound
				}

				// Handle column bounds
				if xBound, ok := colBounds[x]; !ok {
					colBounds[x] = bounds{start: y, end: y}
				} else {
					xBound.end = y
					colBounds[x] = xBound
				}
			}
		}
	}

	var (
		faces       map[face]squareBounds
		transitions map[face]transition
	)

	if sideLength == 4 {
		faces = sampleFaces
		transitions = sampleFaceTransitions
	} else {
		faces = inputFaces
		transitions = inputFaceTransitions
	}

	var (
		pattern = regexp.MustCompile(`(\d+|R|L)`)
		matches = pattern.FindAllStringSubmatch(split[1], -1)
	)

	movements := make([]movement, len(matches))
	for i, match := range matches {
		if match[1] == "R" || match[1] == "L" {
			direction := rune(match[1][0])
			movements[i] = movement{rotation: &direction}
		} else {
			amount := utils.Must(strconv.Atoi(match[1]))
			movements[i] = movement{amount: &amount}
		}
	}

	return &day22{
		grid:        grid,
		start:       *start,
		movements:   movements,
		gridBounds:  gridBounds{row: rowBounds, col: colBounds},
		faces:       faces,
		transitions: transitions,
	}
}

func Solve(filename string) {
	day := Parse(filename, 50)

	fmt.Println("ANSWER1: final password:", day.Part1())
	fmt.Println("ANSWER2: final password of cube map:", day.Part2())
}
