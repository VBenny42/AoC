package day04

import (
	"fmt"
	"image"
	"strings"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

type board struct {
	numbers  [5][5]int
	occupied [5][5]bool
	mapTo    map[int]image.Point
}

type day04 struct {
	numbersToDraw []int
	boards        []board
}

func (b *board) bingo() bool {
	for i := 0; i < 5; i++ {
		if (b.occupied[i][0] &&
			b.occupied[i][1] &&
			b.occupied[i][2] &&
			b.occupied[i][3] &&
			b.occupied[i][4]) ||

			(b.occupied[0][i] &&
				b.occupied[1][i] &&
				b.occupied[2][i] &&
				b.occupied[3][i] &&
				b.occupied[4][i]) {
			return true
		}
	}
	return false
}

func (b *board) calculateScore(number int) int {
	var unmarked int
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if !b.occupied[y][x] {
				unmarked += b.numbers[y][x]
			}
		}
	}
	return unmarked * number
}

func (d *day04) Part1() int {
	for _, number := range d.numbersToDraw {
		for i := range d.boards {
			point, ok := d.boards[i].mapTo[number]
			if !ok {
				continue
			}
			// number is on board, occupy it, and check rows and columns for bingo
			d.boards[i].occupied[point.Y][point.X] = true

			if d.boards[i].bingo() {
				return d.boards[i].calculateScore(number)
			}
		}
	}

	return -1
}

func (d *day04) Part2() int {
	var (
		lastWinner int
		lastNumber int
		bingoFound = make([]bool, len(d.boards))
	)

	for _, number := range d.numbersToDraw {
		for i := range d.boards {
			if bingoFound[i] {
				continue
			}

			point, ok := d.boards[i].mapTo[number]
			if !ok {
				continue
			}

			d.boards[i].occupied[point.Y][point.X] = true

			if d.boards[i].bingo() {
				bingoFound[i] = true
				lastWinner = i
				lastNumber = number
			}
		}
	}

	return d.boards[lastWinner].calculateScore(lastNumber)
}

func Parse(filename string) *day04 {
	var (
		data          = utils.ReadLines(filename)
		boards        = make([]board, 0, len(data[2:])/5)
		numbers       = strings.Split(data[0], ",")
		numbersToDraw = make([]int, len(numbers))
	)

	for i, n := range numbers {
		numbersToDraw[i] = utils.Atoi(n)
	}

	for i := 2; i < len(data)-4; i += 6 {
		var b board
		b.mapTo = make(map[int]image.Point, 25)

		for y := 0; y < 5; y++ {
			fields := strings.Fields(data[i+y])

			for x, n := range fields {
				number := utils.Atoi(n)

				if _, ok := b.mapTo[number]; ok {
					fmt.Println("duplicate number on board")
				}

				b.numbers[y][x] = number
				b.mapTo[number] = image.Pt(x, y)
			}
		}
		boards = append(boards, b)
	}

	return &day04{
		numbersToDraw: numbersToDraw,
		boards:        boards,
	}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: final score for first bingo board:", day.Part1())
	fmt.Println("ANSWER2: final score for last bingo board:", day.Part2())
}
