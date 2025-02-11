package day17

import (
	"fmt"
	"image"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

type day17 struct {
	target image.Rectangle
}

func step(position, velocity image.Point) (newPosition, newVelocity image.Point) {
	newPosition = position.Add(velocity)

	switch {
	case velocity.X > 0:
		newVelocity.X = velocity.X - 1
	case velocity.X < 0:
		newVelocity.X = velocity.X + 1
	}

	newVelocity.Y = velocity.Y - 1

	return
}

func (d *day17) notPastTarget(position image.Point) bool {
	return position.X <= d.target.Max.X &&
		position.Y >= d.target.Min.Y
}

func (d *day17) hitsTarget(velocity image.Point) (maxY int, hit bool) {
	var position image.Point

	for d.notPastTarget(position) {
		position, velocity = step(position, velocity)
		maxY = max(maxY, position.Y)
		if position.In(d.target) {
			hit = true
			return
		}
	}

	return
}

func (d *day17) Part1And2() (int, int) {
	var (
		highestY      int
		hitVelocities []image.Point
	)

	for x := 0; x <= d.target.Max.X; x++ {
		for y := d.target.Min.Y; y <= utils.Abs(d.target.Min.Y); y++ {
			var (
				velocity    = image.Pt(x, y)
				validY, hit = d.hitsTarget(velocity)
			)

			if hit {
				if validY > highestY {
					highestY = validY
				}
				hitVelocities = append(hitVelocities, velocity)
			}
		}
	}

	return highestY, len(hitVelocities)
}

func Parse(filename string) *day17 {
	var (
		data       = utils.ReadTrimmed(filename)
		minX, minY int
		maxX, maxY int
	)

	n, err := fmt.Sscanf(data, "target area: x=%d..%d, y=%d..%d",
		&minX, &maxX,
		&minY, &maxY,
	)
	if err != nil || n != 4 {
		panic(err)
	}

	return &day17{target: image.Rect(
		minX, minY,
		maxX+1, maxY+1,
	)}
}

func Solve(filename string) {
	day := Parse(filename)

	part1, part2 := day.Part1And2()

	fmt.Println(
		"ANSWER1: highest y-position that can be reached and still reach the target:",
		part1,
	)
	fmt.Println("ANSWER2: number of distinct velocities that reach the target:", part2)
}
