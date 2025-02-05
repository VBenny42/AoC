package day24

import (
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type point struct {
	x, y, z float64
}

type hailstone struct {
	pos          point
	vel          point
	verticalPath bool
	slope        float64
}

type day24 struct {
	hailstones []hailstone
	minBound   float64
	maxBound   float64
}

func hailstonesParallel(h1, h2 hailstone) bool {
	if h1.verticalPath && h2.verticalPath {
		return true
	}
	return h1.slope == h2.slope
}

func intersectingCoords(h1, h2 hailstone) (x, y float64, ok bool) {
	if hailstonesParallel(h1, h2) {
		return
	}

	ok = true
	x = (h1.slope*h1.pos.x - h2.slope*h2.pos.x + h2.pos.y - h1.pos.y) /
		(h1.slope - h2.slope)
	y = h1.slope*(x-h1.pos.x) + h1.pos.y

	return
}

func timeToIntersection(h hailstone, x float64) float64 {
	return (x - h.pos.x) / h.vel.x
}

func possibleVelocities(pos1, pos2 int, vel int) (velocities []int) {
	for possible := -1000; possible < 1000; possible++ {
		if possible != vel && (pos1-pos2)%(possible-vel) == 0 {
			velocities = append(velocities, possible)
		}
	}

	return
}

func sliceIntersection(s1, s2 []int) (intersection []int) {
	// for _, v1 := range s1 {
	// 	for _, v2 := range s2 {
	// 		if v1 == v2 {
	// 			intersection = append(intersection, v1)
	// 		}
	// 	}
	// }

	m := make(map[int]struct{}, len(s1))
	for _, v := range s2 {
		m[v] = struct{}{}
	}

	for _, v := range s1 {
		if _, ok := m[v]; ok {
			intersection = append(intersection, v)
		}
	}

	return
}

func (d *day24) Part1() (intersections int) {
	for i, h1 := range d.hailstones {
		for _, h2 := range d.hailstones[i+1:] {
			x, y, ok := intersectingCoords(h1, h2)
			if !ok {
				continue
			}

			if timeToIntersection(h1, x) < 0 || timeToIntersection(h2, x) < 0 {
				continue
			}

			if (x >= d.minBound && x <= d.maxBound) && (y >= d.minBound && y <= d.maxBound) {
				intersections++
			}
		}
	}

	return
}

func (d *day24) Part2() (sum int) {
	var possibleXVel, possibleYVel, possibleZVel []int

	for i, h1 := range d.hailstones {
		for _, h2 := range d.hailstones[i+1:] {
			if h1.vel.x == h2.vel.x {
				possibilities := possibleVelocities(int(h2.pos.x), int(h1.pos.x), int(h1.vel.x))
				if len(possibleXVel) == 0 {
					possibleXVel = possibilities
				} else {
					possibleXVel = sliceIntersection(possibleXVel, possibilities)
				}
			}
			if h1.vel.y == h2.vel.y {
				possibilities := possibleVelocities(int(h2.pos.y), int(h1.pos.y), int(h1.vel.y))
				if len(possibleYVel) == 0 {
					possibleYVel = possibilities
				} else {
					possibleYVel = sliceIntersection(possibleYVel, possibilities)
				}
			}
			if h1.vel.z == h2.vel.z {
				possibilities := possibleVelocities(int(h2.pos.z), int(h1.pos.z), int(h1.vel.z))
				if len(possibleZVel) == 0 {
					possibleZVel = possibilities
				} else {
					possibleZVel = sliceIntersection(possibleZVel, possibilities)
				}
			}
		}
	}

	if len(possibleXVel) != 1 || len(possibleYVel) != 1 || len(possibleZVel) != 1 {
		panic("could not narrow down to a single possible velocity")
	}

	var (
		rockVelX = float64(possibleXVel[0])
		rockVelY = float64(possibleYVel[0])
		rockVelZ = float64(possibleZVel[0])
	)

	// Find two hailstones with the same x velocity to ensure valid calculation
	var h1, h2 hailstone
	found := false
	for i := range d.hailstones {
		for j := i + 1; j < len(d.hailstones); j++ {
			if d.hailstones[i].vel.x == d.hailstones[j].vel.x {
				h1 = d.hailstones[i]
				h2 = d.hailstones[j]
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	if !found {
		panic("no pair found with same x velocity")
	}

	var (
		m1 = (h1.vel.y - rockVelY) / (h1.vel.x - rockVelX)
		m2 = (h2.vel.y - rockVelY) / (h2.vel.x - rockVelX)
		c1 = h1.pos.y - m1*h1.pos.x
		c2 = h2.pos.y - m2*h2.pos.x
	)

	var (
		rockX = (c2 - c1) / (m1 - m2)
		rockY = m1*rockX + c1
		time  = (rockX - h1.pos.x) / (h1.vel.x - rockVelX)
		rockZ = h1.pos.z + (h1.vel.z-rockVelZ)*time
	)

	sum = int(rockX + rockY + rockZ)
	return
}

func Parse(filename string, minBound, maxBound int) *day24 {
	data := utils.ReadLines(filename)
	hailstones := make([]hailstone, len(data))

	for i, line := range data {
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return r == ' ' || r == '@' || r == ','
		})

		hailstones[i].pos.x = float64(utils.Atoi(fields[0]))
		hailstones[i].pos.y = float64(utils.Atoi(fields[1]))
		hailstones[i].pos.z = float64(utils.Atoi(fields[2]))
		hailstones[i].vel.x = float64(utils.Atoi(fields[3]))
		hailstones[i].vel.y = float64(utils.Atoi(fields[4]))
		hailstones[i].vel.z = float64(utils.Atoi(fields[5]))

		hailstones[i].verticalPath = hailstones[i].vel.x == 0

		if !hailstones[i].verticalPath {
			hailstones[i].slope = hailstones[i].vel.y / hailstones[i].vel.x
		}
	}

	return &day24{
		hailstones: hailstones,
		minBound:   float64(minBound),
		maxBound:   float64(maxBound),
	}
}

func Solve(filename string) {
	day := Parse(filename, 200000000000000, 400000000000000)

	fmt.Println("ANSWER1: number of intersections:", day.Part1())
	fmt.Println("ANSWER2: sum of rock's coordinates to hit all hailstones:", day.Part2())
}
