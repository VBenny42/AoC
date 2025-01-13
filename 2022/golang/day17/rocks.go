package day17

import (
	"github.com/VBenny42/AoC/2022/golang/utils"
)

var minusShape = rock{
	shapes: []utils.Coord{
		{X: 2, Y: 0},
		{X: 3, Y: 0},
		{X: 4, Y: 0},
		{X: 5, Y: 0},
	},
	leftMost:  0,
	rightMost: 3,
	bottoms:   []int{0, 1, 2, 3},
}

var plusShape = rock{
	shapes: []utils.Coord{
		{X: 3, Y: 0}, // middle-down
		{X: 3, Y: 1}, // middle-middle
		{X: 4, Y: 1}, // right-middle
		{X: 2, Y: 1}, // left-middle
		{X: 3, Y: 2}, // top-middle
	},
	leftMost:  3,
	rightMost: 2,
	bottoms:   []int{0, 2, 3},
}

var lShape = rock{
	shapes: []utils.Coord{
		{X: 2, Y: 0},
		{X: 3, Y: 0},
		{X: 4, Y: 0},
		{X: 4, Y: 1},
		{X: 4, Y: 2},
	},
	leftMost:  0,
	rightMost: 2,
	bottoms:   []int{0, 1, 2},
}

var lineShape = rock{
	shapes: []utils.Coord{
		{X: 2, Y: 0},
		{X: 2, Y: 1},
		{X: 2, Y: 2},
		{X: 2, Y: 3},
	},
	leftMost:  0,
	rightMost: 0,
	bottoms:   []int{0},
}

var boxShape = rock{
	shapes: []utils.Coord{
		{X: 2, Y: 0},
		{X: 3, Y: 0},
		{X: 2, Y: 1},
		{X: 3, Y: 1},
	},
	leftMost:  0,
	rightMost: 1,
	bottoms:   []int{0, 1},
}
