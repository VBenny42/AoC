package day24

import (
	"fmt"
	"image"

	"github.com/VBenny42/AoC/2020/golang/utils"
)

type day24 struct {
	instructions []string
}

func parseDirections(directions string) image.Point {
	var (
		point image.Point
		i     int
	)

	for i < len(directions) {
		switch {
		case i < len(directions)-1 && directions[i:i+2] == "se":
			point.Y++
			i += 2
		case i < len(directions)-1 && directions[i:i+2] == "sw":
			point.X--
			point.Y++
			i += 2
		case i < len(directions)-1 && directions[i:i+2] == "nw":
			point.Y--
			i += 2
		case i < len(directions)-1 && directions[i:i+2] == "ne":
			point.X++
			point.Y--
			i += 2
		case directions[i] == 'e':
			point.X++
			i++
		case directions[i] == 'w':
			point.X--
			i++
		}
	}

	return point
}

func (d *day24) Part1() int {
	var (
		tiles      = d.getInitialTiles()
		blackCount int
	)

	for _, isBlack := range tiles {
		if isBlack {
			blackCount++
		}
	}

	return blackCount
}

func (d *day24) Part2() int {
	var (
		tiles      = d.getInitialTiles()
		blackCount int
	)

	for day := 0; day < 100; day++ {
		tiles = d.simulateDay(tiles)
	}

	for _, isBlack := range tiles {
		if isBlack {
			blackCount++
		}
	}

	return blackCount
}

func (d *day24) getInitialTiles() map[image.Point]bool {
	tiles := make(map[image.Point]bool)

	for _, instruction := range d.instructions {
		coord := parseDirections(instruction)

		tiles[coord] = !tiles[coord]
	}

	return tiles
}

func getNeighbors(coord image.Point) []image.Point {
	return []image.Point{
		{coord.X + 1, coord.Y},     // e
		{coord.X - 1, coord.Y},     // w
		{coord.X, coord.Y + 1},     // se
		{coord.X - 1, coord.Y + 1}, // sw
		{coord.X + 1, coord.Y - 1}, // ne
		{coord.X, coord.Y - 1},     // nw
	}
}

func (d *day24) countBlackNeighbors(
	coord image.Point,
	tiles map[image.Point]bool,
) (count int) {
	for _, neighbor := range getNeighbors(coord) {
		if tiles[neighbor] { // true means black
			count++
		}
	}
	return
}

// simulateDay runs one day of the game of life simulation
func (d *day24) simulateDay(tiles map[image.Point]bool) map[image.Point]bool {
	newTiles := make(map[image.Point]bool, len(tiles))

	coordsToCheck := make(map[image.Point]bool)

	for coord, isBlack := range tiles {
		if isBlack {
			coordsToCheck[coord] = true
			for _, neighbor := range getNeighbors(coord) {
				coordsToCheck[neighbor] = true
			}
		}
	}

	for coord := range coordsToCheck {
		isCurrentlyBlack := tiles[coord] // false if not in map (white)
		blackNeighbors := d.countBlackNeighbors(coord, tiles)

		var willBeBlack bool
		if isCurrentlyBlack {
			willBeBlack = blackNeighbors == 1 || blackNeighbors == 2
		} else {
			willBeBlack = blackNeighbors == 2
		}

		if willBeBlack {
			newTiles[coord] = true
		}
	}

	return newTiles
}

func Parse(filename string) *day24 {
	return &day24{instructions: utils.ReadLines(filename)}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: number of tiles with black side up:", day.Part1())
	fmt.Println("ANSWER2: number of black tiles after 100 days:", day.Part2())
}
