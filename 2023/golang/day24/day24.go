package day24

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/VBenny42/AoC/2023/golang/utils"
	"gonum.org/v1/gonum/stat/combin"
)

type point struct {
	x, y, z int64
}

type hailstone struct {
	pos point
	vel point
}

type day24 struct {
	hailstones []hailstone
	minBound   int64
	maxBound   int64
}

func (d *day24) Part1() (intersections int) {
	combinations := combin.Combinations(len(d.hailstones), 2)

	for _, combo := range combinations {
		h1 := d.hailstones[combo[0]]
		h2 := d.hailstones[combo[1]]

		// Coefficients for the system of equations:
		// vel.x1 * t1 - vel.x2 * t2 = pos.x2 - pos.x1
		// vel.y1 * t1 - vel.y2 * t2 = pos.y2 - pos.y1
		// Matrix form: [vel.x1, -vel.x2; vel.y1, -vel.y2] * [t1; t2] = [dx; dy]
		var (
			dx  = h2.pos.x - h1.pos.x
			dy  = h2.pos.y - h1.pos.y
			vx1 = h1.vel.x
			vy1 = h1.vel.y
			vx2 = h2.vel.x
			vy2 = h2.vel.y
		)

		// Compute determinant
		det := vx1*(-vy2) - (-vx2)*vy1
		if det == 0 {
			// Paths are parallel, check if they are the same line
			// Check cross product: dx*vel.y1 - dy*vel.x1 == 0
			if dx*vy1 != dy*vx1 {
				continue
			}
			// Paths are the same line, check if there's any t1, t2 >=0 that brings both into the test area
			if d.coincidingPathsIntersect(h1, h2) {
				intersections++
			}
			continue
		}

		// Solve for t1 and t2 using Cramer's rule with exact fractions
		// t1 = (dx*(-vel.y2) - (-vel.x2)*dy) / det
		// t2 = (vel.x1*dy - vel.y1*dx) / det
		numeratorT1 := dx*(-vy2) - (-vx2)*dy
		numeratorT2 := vx1*dy - vy1*dx

		// Check if t1 and t2 are non-negative
		if (numeratorT1 > 0 && det < 0) || (numeratorT1 < 0 && det > 0) {
			// t1 is negative
			continue
		}
		if (numeratorT2 > 0 && det < 0) || (numeratorT2 < 0 && det > 0) {
			// t2 is negative
			continue
		}

		// Compute absolute values to check if numerator and denominator have the same sign
		absNumT1 := numeratorT1
		if absNumT1 < 0 {
			absNumT1 = -absNumT1
		}
		absDet := det
		if absDet < 0 {
			absDet = -absDet
		}

		// Compute t1 and t2 as fractions to avoid precision loss
		t1Num := big.NewInt(numeratorT1)
		t1Den := big.NewInt(det)
		t1Rat := new(big.Rat).SetFrac(t1Num, t1Den)

		t2Num := big.NewInt(numeratorT2)
		t2Den := big.NewInt(det)
		t2Rat := new(big.Rat).SetFrac(t2Num, t2Den)

		if t1Rat.Sign() < 0 || t2Rat.Sign() < 0 {
			continue
		}

		// Calculate intersection point using h1's path at t1
		xNum := big.NewInt((h1.pos.x))
		xNum.Mul(xNum, t1Den)
		xNum.Add(xNum, big.NewInt(h1.vel.x).Mul(big.NewInt(h1.vel.x), t1Num))
		xDen := t1Den

		yNum := big.NewInt(int64(h1.pos.y))
		yNum.Mul(yNum, t1Den)
		yNum.Add(yNum, big.NewInt(h1.vel.y).Mul(big.NewInt(h1.vel.y), t1Num))
		yDen := t1Den

		// Check if x and y are within the test area
		if d.isWithinBounds(xNum, xDen) && d.isWithinBounds(yNum, yDen) {
			intersections++
		}
	}

	return
}

func (d *day24) isWithinBounds(num, den *big.Int) bool {
	if den.Sign() == 0 {
		return false
	}
	minBound := big.NewRat(d.minBound, 1)
	maxBound := big.NewRat(d.maxBound, 1)

	val := new(big.Rat).SetFrac(num, den)

	return val.Cmp(minBound) >= 0 && val.Cmp(maxBound) <= 0
}

func (d *day24) coincidingPathsIntersect(h1, h2 hailstone) bool {
	return d.pathEntersTestArea(h1) || d.pathEntersTestArea(h2)
}

func (d *day24) pathEntersTestArea(h hailstone) bool {
	// Check if the hailstone's path enters the test area at any time t >=0
	enterX := false
	if h.vel.x > 0 {
		enterX = h.pos.x <= d.maxBound
	} else if h.vel.x < 0 {
		enterX = h.pos.x >= d.minBound
	} else {
		enterX = h.pos.x >= d.minBound && int64(h.pos.x) <= d.maxBound
	}

	enterY := false
	if h.vel.y > 0 {
		enterY = h.pos.y <= d.maxBound
	} else if h.vel.y < 0 {
		enterY = h.pos.y >= d.minBound
	} else {
		enterY = h.pos.y >= d.minBound && int64(h.pos.y) <= d.maxBound
	}

	return enterX && enterY
}

func Parse(filename string, minBound, maxBound int) *day24 {
	data := utils.ReadLines(filename)
	hailstones := make([]hailstone, len(data))

	fieldsFunc := func(r rune) bool {
		return r == ' ' || r == '@' || r == ','
	}

	for i, line := range data {
		fields := strings.FieldsFunc(line, fieldsFunc)

		hailstones[i].pos.x = int64(utils.Atoi(fields[0]))
		hailstones[i].pos.y = int64(utils.Atoi(fields[1]))
		hailstones[i].pos.z = int64(utils.Atoi(fields[2]))
		hailstones[i].vel.x = int64(utils.Atoi(fields[3]))
		hailstones[i].vel.y = int64(utils.Atoi(fields[4]))
		hailstones[i].vel.z = int64(utils.Atoi(fields[5]))

	}

	return &day24{
		hailstones: hailstones,
		minBound:   int64(minBound),
		maxBound:   int64(maxBound),
	}
}

func Solve(filename string) {
	day := Parse(filename, 200000000000000, 400000000000000)
	fmt.Println("ANSWER1: number of intersections:", day.Part1())
}
