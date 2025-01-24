package day15

import (
	"fmt"
	"sort"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type pair struct {
	sensor utils.Coord
	beacon utils.Coord
	bounds int
}

type day15 struct {
	pairs      []pair
	rowToCheck int
	bounds     int
}

type kleePoint struct {
	x     int
	isEnd bool
}

type span struct {
	start int
	end   int
}

func manhattanDistance(a, b utils.Coord) int {
	return utils.Abs(a.X-b.X) + utils.Abs(a.Y-b.Y)
}

func (p *pair) manhattanSpan(row int) (span, bool) {
	dy := utils.Abs(p.sensor.Y - row)

	if dx := p.bounds - dy; dx >= 0 {
		return span{p.sensor.X - dx, p.sensor.X + dx}, true
	} else {
		return span{}, false
	}
}

func buildKleePoints(row int, pairs []pair) []kleePoint {
	points := make([]kleePoint, 0)

	for _, p := range pairs {
		if span, ok := p.manhattanSpan(row); ok {
			points = append(points, kleePoint{span.start, false})
			points = append(points, kleePoint{span.end, true})
		}
	}

	sort.Slice(points, func(i, j int) bool {
		left, right := points[i], points[j]
		if left.x < right.x {
			return true
		} else if left.x > right.x {
			return false
		} else {
			return left.isEnd
		}
	})

	return points
}

// https://iq.opengenus.org/klee-algorithm/
func calculateTotalLength(points []kleePoint) int {
	totalLength := 0
	depth := 1

	for i := range len(points) - 1 {
		curr, next := points[i], points[i+1]
		if diff := next.x - curr.x; depth > 0 && diff > 0 {
			totalLength += diff
		}
		if next.isEnd {
			depth--
		} else {
			depth++
		}
	}
	return totalLength
}

func (p *pair) inExclusionRange(point utils.Coord) bool {
	return p.bounds >= manhattanDistance(p.sensor, point)
}

func (d *day15) notInAnyRange(point utils.Coord) bool {
	for _, p := range d.pairs {
		if p.inExclusionRange(point) {
			return false
		}
	}
	return true
}

func tuningFrequency(c utils.Coord) int {
	return c.X*4000000 + c.Y
}

func (d *day15) Part1() int {
	points := buildKleePoints(d.rowToCheck, d.pairs)
	return calculateTotalLength(points)
}

// Borrowed from https://github.com/BuonHobo/advent-of-code/blob/master/2022/15/Alex/second.py
func (d *day15) Part2() int {
	type lineEqn struct {
		// y = mx + c
		isRising bool
		c        int
	}

	lines := make(map[lineEqn]int)

	for _, p := range d.pairs {
		topRising := lineEqn{true, p.sensor.Y - p.bounds - 1 - p.sensor.X}
		topFalling := lineEqn{false, p.sensor.Y - p.bounds - 1 + p.sensor.X}
		bottomRising := lineEqn{true, p.sensor.Y + p.bounds + 1 - p.sensor.X}
		bottomFalling := lineEqn{false, p.sensor.Y + p.bounds + 1 + p.sensor.X}

		for _, line := range []lineEqn{topRising, topFalling, bottomRising, bottomFalling} {
			lines[line]++
		}
	}

	risingLines := make([]int, 0)
	fallingLines := make([]int, 0)

	for line, count := range lines {
		if count > 1 {
			if !line.isRising {
				risingLines = append(risingLines, line.c)
			} else {
				fallingLines = append(fallingLines, line.c)
			}
		}
	}

	points := make([]utils.Coord, 0)

	for _, rising := range risingLines {
		for _, falling := range fallingLines {
			x := (rising - falling) / 2
			y := x + falling
			points = append(points, utils.Crd(x, y))
		}
	}

	for _, point := range points {
		if (0 <= point.X && point.X < d.bounds) &&
			(0 <= point.Y && point.Y < d.bounds) &&
			d.notInAnyRange(point) {
			return tuningFrequency(point)
		}
	}

	return -1
}

func Parse(filename string, rowToCheck, bounds int) *day15 {
	data := utils.ReadLines(filename)

	pairs := make([]pair, len(data))

	for i, line := range data {
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&pairs[i].sensor.X, &pairs[i].sensor.Y, &pairs[i].beacon.X, &pairs[i].beacon.Y)
		pairs[i].bounds = manhattanDistance(pairs[i].sensor, pairs[i].beacon)
	}

	return &day15{pairs, rowToCheck, bounds}
}

func Solve(filename string) {
	day := Parse(filename, 2000000, 4000000)

	fmt.Println("ANSWER1: number of positions with no beacon in row 2000000:", day.Part1())
	fmt.Println("ANSWER2: tuning frequency of the only vacant point:", day.Part2())
}
