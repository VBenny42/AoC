package day15

import (
	"fmt"
	"sort"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type pair struct {
	sensor utils.Coord
	beacon utils.Coord
}

type day15 struct {
	pairs      []pair
	rowToCheck int
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
	radius := manhattanDistance(p.sensor, p.beacon)

	dy := utils.Abs(p.sensor.Y - row)

	if dx := radius - dy; dx >= 0 {
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

func calculateTotalLength(points []kleePoint) int {
	totalLength := 0
	depth := 1

	for i := 1; i < len(points); i++ {
		prev, curr := points[i-1], points[i]
		if diff := curr.x - prev.x; depth > 0 && diff > 0 {
			totalLength += diff
		}
		if curr.isEnd {
			depth--
		} else {
			depth++
		}
	}
	return totalLength
}

func (d *day15) Part1() int {
	points := buildKleePoints(d.rowToCheck, d.pairs)
	return calculateTotalLength(points)
}

func Parse(filename string, rowToCheck int) *day15 {
	data := utils.ReadLines(filename)

	pairs := make([]pair, len(data))

	for i, line := range data {
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&pairs[i].sensor.X, &pairs[i].sensor.Y, &pairs[i].beacon.X, &pairs[i].beacon.Y)
	}

	return &day15{pairs, rowToCheck}
}

func Solve(filename string) {
	day := Parse(filename, 2000000)

	fmt.Println("ANSWER1: number of positions with no beacon in row 2000000:", day.Part1())
}
