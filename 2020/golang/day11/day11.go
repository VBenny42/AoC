package day11

import (
	"fmt"
	"image"
	"strings"

	"github.com/VBenny42/AoC/2020/golang/utils"
)

const (
	floor    = '.'
	empty    = 'L'
	occupied = '#'
)

type day11 struct {
	layout string
	width  int
	height int
}

func (d *day11) String() string {
	var result strings.Builder
	for i, char := range d.layout {
		if i > 0 && i%d.width == 0 {
			result.WriteByte('\n')
		}
		result.WriteRune(char)
	}
	return result.String()
}

func (d *day11) indexToCoords(index int) image.Point {
	return image.Pt(index%d.width, index/d.width)
}

func (d *day11) coordsToIndex(point image.Point) int {
	return point.Y*d.width + point.X
}

func (d *day11) isValidCoord(point image.Point) bool {
	return point.X >= 0 && point.X < d.width && point.Y >= 0 && point.Y < d.height
}

func (d *day11) getNeighborIndices(index int) []int {
	point := d.indexToCoords(index)
	var neighbors []int

	for _, dir := range utils.AllDirections {
		neighbor := point.Add(dir)
		if d.isValidCoord(neighbor) {
			neighbors = append(neighbors, d.coordsToIndex(neighbor))
		}
	}

	return neighbors
}

func (d *day11) countOccupiedNeighbors(index int) int {
	count := 0
	neighbors := d.getNeighborIndices(index)

	for _, neighborIndex := range neighbors {
		if d.layout[neighborIndex] == occupied {
			count++
		}
	}

	return count
}

func (d *day11) changeSeats() bool {
	newLayout := make([]rune, len(d.layout))
	changed := false

	for i, seat := range d.layout {
		if seat == floor {
			newLayout[i] = floor
			continue
		}

		occupiedCount := d.countOccupiedNeighbors(i)

		if seat == empty && occupiedCount == 0 {
			newLayout[i] = occupied
			changed = true
		} else if seat == occupied && occupiedCount >= 4 {
			newLayout[i] = empty
			changed = true
		} else {
			newLayout[i] = seat
		}
	}

	d.layout = string(newLayout)
	return changed
}

func (d *day11) getVisibleNeighborIndices(index int) []int {
	point := d.indexToCoords(index)
	var neighbors []int

	for _, dir := range utils.AllDirections {
		for step := 1; ; step++ {
			neighbor := point.Add(dir.Mul(step))
			if !d.isValidCoord(neighbor) {
				break
			}

			seatIndex := d.coordsToIndex(neighbor)
			if d.layout[seatIndex] != floor {
				neighbors = append(neighbors, seatIndex)
				break
			}
		}
	}

	return neighbors
}

func (d *day11) countVisibleOccupiedSeats(index int) int {
	count := 0
	neighbors := d.getVisibleNeighborIndices(index)

	for _, neighborIndex := range neighbors {
		if d.layout[neighborIndex] == occupied {
			count++
		}
	}

	return count
}

func (d *day11) changeSeatsPart2() bool {
	newLayout := make([]rune, len(d.layout))
	changed := false

	for i, seat := range d.layout {
		if seat == floor {
			newLayout[i] = floor
			continue
		}

		occupiedCount := d.countVisibleOccupiedSeats(i)

		if seat == empty && occupiedCount == 0 {
			newLayout[i] = occupied
			changed = true
		} else if seat == occupied && occupiedCount >= 5 {
			newLayout[i] = empty
			changed = true
		} else {
			newLayout[i] = seat
		}
	}

	d.layout = string(newLayout)
	return changed
}

func (d *day11) Part1() (seats int) {
	for d.changeSeats() {
	}

	for _, seat := range d.layout {
		if seat == occupied {
			seats++
		}
	}

	return
}

func (d *day11) Part2() (seats int) {
	for d.changeSeatsPart2() {
	}

	for _, seat := range d.layout {
		if seat == occupied {
			seats++
		}
	}

	return
}

func Parse(filename string) *day11 {
	lines := utils.ReadLines(filename)

	width := len(lines[0])
	height := len(lines)

	var layout strings.Builder
	layout.Grow(width * height)

	for _, line := range lines {
		if len(line) != width {
			panic("inconsistent line lengths")
		}

		for _, char := range line {
			switch char {
			case '.', 'L', '#':
				layout.WriteRune(char)
			default:
				panic(fmt.Sprintf("unexpected character '%c'", char))
			}
		}
	}

	return &day11{
		layout: layout.String(),
		width:  width,
		height: height,
	}
}

func Solve(filename string) {
	day := Parse(filename)
	day1 := &day11{layout: day.layout, width: day.width, height: day.height}
	day2 := &day11{layout: day.layout, width: day.width, height: day.height}

	fmt.Println(
		"ANSWER1: number of occupied seats after no changes:",
		day1.Part1(),
	)
	fmt.Println(
		"ANSWER2: number of occupied seats with visibility rules:",
		day2.Part2(),
	)

	// Do I care enough about the speed of this code to use goroutines?
	// SolveGoRoutinesBuffered(filename)
}

func SolveGoRoutinesBuffered(filename string) {
	results := make(chan struct {
		part  int
		value int
	}, 2)

	day := Parse(filename)
	day1 := &day11{layout: day.layout, width: day.width, height: day.height}
	day2 := &day11{layout: day.layout, width: day.width, height: day.height}

	go func() {
		results <- struct {
			part  int
			value int
		}{1, day1.Part1()}
	}()

	go func() {
		results <- struct {
			part  int
			value int
		}{2, day2.Part2()}
	}()

	for i := 0; i < 2; i++ {
		result := <-results
		switch result.part {
		case 1:
			fmt.Println("ANSWER1: number of occupied seats after no changes:", result.value)
		case 2:
			fmt.Println("ANSWER2: number of occupied seats with visibility rules:", result.value)
		}
	}
}
