package day15

import (
	"fmt"
	"runtime"
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

	if dx := manhattanDistance(p.sensor, p.beacon) - dy; dx >= 0 {
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

func (d *day15) findVacantPoint() utils.Coord {
	beacons := d.beaconsByRow()

	results := make(chan utils.Coord)
	failed := make(chan bool)

	numWorkers := runtime.NumCPU()

	for row := range numWorkers {
		go checkRows(row, numWorkers, d.bounds, d.pairs, beacons, results, failed)
	}

	for remaining := numWorkers; remaining > 0; {
		select {
		case <-failed:
			remaining--
		case result := <-results:
			// Guaranteed to have one vacant point, so return immediately
			return result
		}
	}

	return utils.Coord{}
}

func tuningFrequency(c utils.Coord) int {
	return c.X*4000000 + c.Y
}

func checkRows(
	startRow int,
	step int,
	bounds int,
	pairs []pair,
	beacons map[int][]span,
	results chan utils.Coord,
	failed chan bool,
) {
	spans := make([]span, 0, len(pairs))

	for row := startRow; row < bounds; row += step {
		spans := spans[:0]

		if beaconsInRow, ok := beacons[row]; ok {
			spans = append(spans, beaconsInRow...)
		}

		for _, p := range pairs {
			if span, ok := p.manhattanSpan(row); ok {
				spans = append(spans, span)
			}
		}

		sort.Slice(spans, func(i, j int) bool {
			return spans[i].start < spans[j].start
		})

		currentSpan := spans[0]

		for _, s := range spans[1:] {
			if s.start <= currentSpan.end+1 {
				if s.end > currentSpan.end {
					currentSpan.end = s.end
				}
			} else {
				// Found a gap
				results <- utils.Coord{X: currentSpan.end + 1, Y: row}
				return
			}
		}
	}

	failed <- true
}

func (d *day15) beaconsByRow() map[int][]span {
	beacons := make(map[int][]span)

	for _, p := range d.pairs {
		row := p.beacon.Y
		beaconSpan := span{p.beacon.X, p.beacon.X}

		if spans, ok := beacons[row]; ok {
			beacons[row] = append(spans, beaconSpan)
		} else {
			beacons[row] = []span{beaconSpan}
		}
	}

	return beacons
}

func (d *day15) Part1() int {
	points := buildKleePoints(d.rowToCheck, d.pairs)
	return calculateTotalLength(points)
}

func (d *day15) Part2() int {
	vacantPoint := d.findVacantPoint()
	return tuningFrequency(vacantPoint)
}

func Parse(filename string, rowToCheck, bounds int) *day15 {
	data := utils.ReadLines(filename)

	pairs := make([]pair, len(data))

	for i, line := range data {
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&pairs[i].sensor.X, &pairs[i].sensor.Y, &pairs[i].beacon.X, &pairs[i].beacon.Y)
	}

	return &day15{pairs, rowToCheck, bounds}
}

func Solve(filename string) {
	day := Parse(filename, 2000000, 4000000)

	fmt.Println("ANSWER1: number of positions with no beacon in row 2000000:", day.Part1())
	fmt.Println("ANSWER2: tuning frequency of the only vacant point:", day.Part2())
}
