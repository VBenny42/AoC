package day17

import (
	"fmt"

	"github.com/VBenny42/AoC/2020/golang/utils"
)

type point3d struct {
	x, y, z int
}

type point4d struct {
	x, y, z, w int
}

type day17 struct {
	space3d map[point3d]bool
	space4d map[point4d]bool
}

func (p *point3d) neighbors() (neighbors []point3d) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				if dx == 0 && dy == 0 && dz == 0 {
					continue
				}
				neighbors = append(neighbors, point3d{p.x + dx, p.y + dy, p.z + dz})
			}
		}
	}
	return
}

func (p *point4d) neighbors() (neighbors []point4d) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dw := -1; dw <= 1; dw++ {
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
						continue
					}
					neighbors = append(neighbors, point4d{p.x + dx, p.y + dy, p.z + dz, p.w + dw})
				}
			}
		}
	}
	return
}

func (d *day17) cycle3d() {
	newSpace := make(map[point3d]bool)
	neighborsCount := make(map[point3d]int)

	for p := range d.space3d {
		for _, neighbor := range p.neighbors() {
			neighborsCount[neighbor]++
		}
	}

	for p, count := range neighborsCount {
		if d.space3d[p] {
			if count == 2 || count == 3 {
				newSpace[p] = true
			}
		} else {
			if count == 3 {
				newSpace[p] = true
			}
		}
	}

	d.space3d = newSpace
}

func (d *day17) cycle4d() {
	newSpace := make(map[point4d]bool)
	neighborsCount := make(map[point4d]int)

	for p := range d.space4d {
		for _, neighbor := range p.neighbors() {
			neighborsCount[neighbor]++
		}
	}

	for p, count := range neighborsCount {
		if d.space4d[p] {
			if count == 2 || count == 3 {
				newSpace[p] = true
			}
		} else {
			if count == 3 {
				newSpace[p] = true
			}
		}
	}

	d.space4d = newSpace
}

func (d *day17) Part1() (sum int) {
	for range 6 {
		d.cycle3d()
	}

	for _, v := range d.space3d {
		if v == true {
			sum++
		}
	}

	return
}

func (d *day17) Part2() (sum int) {
	for range 6 {
		d.cycle4d()
	}

	for _, v := range d.space4d {
		if v == true {
			sum++
		}
	}

	return
}

func Parse(filename string) *day17 {
	data := utils.ReadLines(filename)
	space3d := make(map[point3d]bool)
	space4d := make(map[point4d]bool)

	for y, line := range data {
		for x, char := range line {
			if char == '#' {
				space3d[point3d{x, y, 0}] = true    // z is 0 for the initial state
				space4d[point4d{x, y, 0, 0}] = true // w is also 0 for the initial state
			}
		}
	}

	return &day17{space3d: space3d, space4d: space4d}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: active cubes after 6 cycles:", day.Part1())
	fmt.Println("ANSWER2: active cubes in 4D after 6 cycles:", day.Part2())
}
