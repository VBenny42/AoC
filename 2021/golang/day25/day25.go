package day25

import (
	"fmt"
	"image"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

const (
	east  = '>'
	south = 'v'
	empty = '.'
)

var gridWidth, gridHeight int

type day25 struct {
	grid      utils.Grid[rune]
	eastList  []image.Point
	southList []image.Point
}

func (d *day25) step() (changed int) {
	var eastMoves []int
	for i, pos := range d.eastList {
		next := pos.Add(utils.Right)
		next.X %= gridWidth
		if d.grid.Get(next) == empty {
			eastMoves = append(eastMoves, i)
		}
	}

	for _, i := range eastMoves {
		oldPos := d.eastList[i]
		newPos := oldPos.Add(utils.Right)
		newPos.X %= gridWidth

		d.grid.Set(oldPos, empty)
		d.grid.Set(newPos, east)
		changed++

		d.eastList[i] = newPos
	}

	var southMoves []int
	for i, pos := range d.southList {
		next := pos.Add(utils.Down)
		next.Y %= gridHeight
		if d.grid.Get(next) == empty {
			southMoves = append(southMoves, i)
		}
	}

	for _, i := range southMoves {
		oldPos := d.southList[i]
		newPos := oldPos.Add(utils.Down)
		newPos.Y %= gridHeight

		d.grid.Set(oldPos, empty)
		d.grid.Set(newPos, south)
		changed++

		d.southList[i] = newPos
	}

	return
}

func (d *day25) Part1() (answer int) {
	for {
		answer++
		if changed := d.step(); changed == 0 {
			break
		}
	}

	return
}

func Parse(filename string) *day25 {
	var (
		data      = utils.ReadLines(filename)
		grid      = utils.NewGrid[rune](len(data[0]), len(data))
		eastList  []image.Point
		southList []image.Point
	)

	gridWidth, gridHeight = len(data[0]), len(data)

	for y, line := range data {
		for x, r := range line {
			pt := image.Pt(x, y)
			grid.Set(pt, r)

			switch r {
			case east:
				eastList = append(eastList, pt)
			case south:
				southList = append(southList, pt)
			default:
				// empty
			}
		}
	}

	return &day25{grid, eastList, southList}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER: first step on which no cucumbers move:", day.Part1())
}
