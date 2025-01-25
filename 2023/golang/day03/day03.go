package day03

import (
	"fmt"
	"unicode"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type grid [][]rune

type number struct {
	value int
	start utils.Coord
	end   utils.Coord
}

type day03 struct {
	grid    grid
	gears   []utils.Coord
	numbers []number
}

func (d *day03) Part1() (sum int) {
	length := len(d.grid[0])

	for _, num := range d.numbers {
		symbolAdjacent := false

		for x := num.start.X; x <= num.end.X; x++ {
			coord := utils.Crd(x, num.start.Y)
			for _, direction := range utils.AllDirections {
				adj := coord.Add(direction)

				isValid := adj.X >= 0 && adj.X < length &&
					adj.Y >= 0 && adj.Y < len(d.grid)

				if isValid &&
					d.grid[adj.Y][adj.X] != '.' &&
					!unicode.IsDigit(d.grid[adj.Y][adj.X]) {
					symbolAdjacent = true
					break
				}
			}

			if symbolAdjacent {
				break
			}
		}

		if symbolAdjacent {
			sum += num.value
		}
	}

	return
}

func (d *day03) Part2() (sum int) {
	yMap := make(map[int][]number)

	for _, num := range d.numbers {
		yMap[num.start.Y] = append(yMap[num.start.Y], num)
	}

	length := len(d.grid[0])

	// Check up, same row, down
	directions := []utils.Coord{
		utils.Up,
		utils.Down,
		utils.Crd(0, 0),
	}

	for _, gear := range d.gears {
		adjacentNumbers := make([]number, 0, 2)

		for _, direction := range directions {
			coord := gear.Add(direction)

			isValid := coord.X >= 0 && coord.X < length &&
				coord.Y >= 0 && coord.Y < len(d.grid)

			if isValid {
				partsToCheck, ok := yMap[coord.Y]
				if !ok {
					continue
				}

				leftC, rightC := coord.X-1, coord.X+1
				for _, part := range partsToCheck {
					if (leftC >= part.start.X && leftC <= part.end.X) ||
						(coord.X >= part.start.X && coord.X <= part.end.X) ||
						(rightC >= part.start.X && rightC <= part.end.X) {
						adjacentNumbers = append(adjacentNumbers, part)
					}
				}
			}
		}

		if len(adjacentNumbers) == 2 {
			sum += adjacentNumbers[0].value * adjacentNumbers[1].value
		}
	}

	return
}

func Parse(filename string) *day03 {
	var (
		data    = utils.ReadLines(filename)
		grid    = make(grid, len(data))
		length  = len(data[0])
		gears   = make([]utils.Coord, 0)
		numbers = make([]number, 0)
	)

	for y, line := range data {
		grid[y] = make([]rune, length)
		grid[y] = []rune(line)
		for x := 0; x < length; x++ {
			char := grid[y][x]
			if unicode.IsDigit(char) {
				num := number{start: utils.Coord{X: x, Y: y}}
				for x < length &&
					unicode.IsDigit(grid[y][x]) {
					num.value = num.value*10 + int(grid[y][x]-'0')
					num.end = utils.Coord{X: x, Y: y}
					x++
				}
				numbers = append(numbers, num)
			}
		}
	}

	for y, line := range grid {
		for x, char := range line {
			if char == '*' {
				gears = append(gears, utils.Crd(x, y))
			}
		}
	}

	return &day03{grid, gears, numbers}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: sum of all part numbers:", day.Part1())
	fmt.Println("ANSWER2: sum of all gear ratios:", day.Part2())
}
