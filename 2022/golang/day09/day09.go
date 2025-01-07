package day09

import (
	"fmt"

	"github.com/VBenny42/AoC/2022/golang/utils"
	"github.com/deckarep/golang-set/v2"
)

type Motion struct {
	Direction rune
	Distance  int
}

type day09 struct {
	Motions []Motion
}

func updateHead(head *utils.Coord, direction rune) {
	switch direction {
	case 'R':
		head.Y++
	case 'L':
		head.Y--
	case 'U':
		head.X++
	case 'D':
		head.X--
	}
}

var headTailDiff = map[utils.Coord]utils.Coord{
	{X: 2, Y: 1}:   {X: 1, Y: 1},
	{X: 1, Y: 2}:   {X: 1, Y: 1},
	{X: 2, Y: 0}:   {X: 1, Y: 0},
	{X: 2, Y: -1}:  {X: 1, Y: -1},
	{X: 1, Y: -2}:  {X: 1, Y: -1},
	{X: 0, Y: -2}:  {X: 0, Y: -1},
	{X: -1, Y: -2}: {X: -1, Y: -1},
	{X: -2, Y: -1}: {X: -1, Y: -1},
	{X: -2, Y: 0}:  {X: -1, Y: 0},
	{X: -2, Y: 1}:  {X: -1, Y: 1},
	{X: -1, Y: 2}:  {X: -1, Y: 1},
	{X: 0, Y: 2}:   {X: 0, Y: 1},
	{X: 2, Y: 2}:   {X: 1, Y: 1},
	{X: -2, Y: -2}: {X: -1, Y: -1},
	{X: -2, Y: 2}:  {X: -1, Y: 1},
	{X: 2, Y: -2}:  {X: 1, Y: -1},
}

func updateTail(tail, head utils.Coord) utils.Coord {
	return tail.Add(headTailDiff[head.Sub(tail)])
}

func (d *day09) findTailPositions(knotNum int) int {
	ropeStack := make([]utils.Coord, knotNum)
	tailSet := mapset.NewSet(utils.Coord{X: 0, Y: 0})

	for _, motion := range d.Motions {
		for range motion.Distance {
			updateHead(&ropeStack[0], motion.Direction)
			for i := 0; i < knotNum-1; i++ {
				ropeStack[i+1] = updateTail(ropeStack[i+1], ropeStack[i])
			}
			tailSet.Add(ropeStack[knotNum-1])
		}
	}

	return tailSet.Cardinality()
}

func (d *day09) Part1() int {
	return d.findTailPositions(2)
}

func (d *day09) Part2() int {
	return d.findTailPositions(10)
}

func Parse(filename string) *day09 {
	data := utils.ReadLines(filename)

	motions := make([]Motion, len(data))

	for i, line := range data {
		fmt.Sscanf(line, "%c %d", &motions[i].Direction, &motions[i].Distance)
	}

	return &day09{motions}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: number of distinct tail positions:", day.Part1())
	fmt.Println("ANSWER2: number of distinct tail positions:", day.Part2())
}
