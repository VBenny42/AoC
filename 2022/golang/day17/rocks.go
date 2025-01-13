package day17

import (
	"github.com/VBenny42/AoC/2022/golang/utils"
)

var minusShape rock = []utils.Coord{
	{X: 0, Y: 0},
	{X: 1, Y: 0},
	{X: 2, Y: 0},
	{X: 3, Y: 0},
}

var plusShape rock = []utils.Coord{
	{X: 0, Y: 1},
	{X: 1, Y: 0},
	{X: 1, Y: 1},
	{X: 1, Y: 2},
	{X: 2, Y: 1},
}

var lShape rock = []utils.Coord{
	{X: 0, Y: 0},
	{X: 1, Y: 0},
	{X: 2, Y: 0},
	{X: 2, Y: 1},
	{X: 2, Y: 2},
}

var lineShape rock = []utils.Coord{
	{X: 0, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: 2},
	{X: 0, Y: 3},
}

var boxShape rock = []utils.Coord{
	{X: 0, Y: 0},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: 1, Y: 1},
}
