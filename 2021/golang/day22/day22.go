package day22

import (
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

type (
	point struct {
		x, y, z int
	}
	cube struct {
		min, max point
		isOn     bool
	}
)

type day22 struct {
	cubes []cube
}

func (c cube) intersection(other cube) (intersection cube, intersects bool) {
	intersection.min.x = max(c.min.x, other.min.x)
	intersection.min.y = max(c.min.y, other.min.y)
	intersection.min.z = max(c.min.z, other.min.z)

	intersection.max.x = min(c.max.x, other.max.x)
	intersection.max.y = min(c.max.y, other.max.y)
	intersection.max.z = min(c.max.z, other.max.z)

	if intersection.min.x > intersection.max.x ||
		intersection.min.y > intersection.max.y ||
		intersection.min.z > intersection.max.z {
		return
	}

	var isOn bool
	if c.isOn && other.isOn {
		isOn = false
	} else if !c.isOn && !other.isOn {
		isOn = true
	} else {
		isOn = other.isOn
	}

	intersection.isOn = isOn
	intersects = true
	return
}

func (c cube) volume() int {
	volume := (c.max.x - c.min.x + 1) * (c.max.y - c.min.y + 1) * (c.max.z - c.min.z + 1)
	if c.isOn {
		return volume
	}
	return -volume
}

func (d *day22) Part1() (sum int) {
	var finalCubes []cube

	for _, c := range d.cubes {
		if c.min.x < -50 || c.min.x > 50 ||
			c.min.y < -50 || c.min.y > 50 ||
			c.min.z < -50 || c.min.z > 50 ||
			c.max.x < -50 || c.max.x > 50 ||
			c.max.y < -50 || c.max.y > 50 ||
			c.max.z < -50 || c.max.z > 50 {
			continue
		}

		var cubesToAdd []cube

		for _, fc := range finalCubes {
			intersection, intersects := fc.intersection(c)
			if intersects {
				cubesToAdd = append(cubesToAdd, intersection)
			}
		}

		if c.isOn {
			cubesToAdd = append(cubesToAdd, c)
		}

		finalCubes = append(finalCubes, cubesToAdd...)
	}

	for _, c := range finalCubes {
		sum += c.volume()
	}

	return
}

func (d *day22) Part2() (sum int) {
	var finalCubes []cube

	for _, c := range d.cubes {
		var cubesToAdd []cube

		for _, fc := range finalCubes {
			intersection, ok := fc.intersection(c)
			if ok {
				cubesToAdd = append(cubesToAdd, intersection)
			}
		}

		if c.isOn {
			cubesToAdd = append(cubesToAdd, c)
		}

		finalCubes = append(finalCubes, cubesToAdd...)
	}

	for _, c := range finalCubes {
		sum += c.volume()
	}

	return
}

func Parse(filename string) *day22 {
	var (
		data  = utils.ReadLines(filename)
		cubes = make([]cube, len(data))
	)

	for i, line := range data {
		split := strings.SplitN(line, " ", 2)
		if len(split) != 2 {
			panic(fmt.Sprintf("Space expected in %s, got: %v", line, split))
		}

		cubes[i].isOn = split[0] == "on"

		n, err := fmt.Sscanf(
			split[1],
			"x=%d..%d,y=%d..%d,z=%d..%d",
			&cubes[i].min.x,
			&cubes[i].max.x,
			&cubes[i].min.y,
			&cubes[i].max.y,
			&cubes[i].min.z,
			&cubes[i].max.z,
		)
		if n != 6 || err != nil {
			panic(fmt.Sprintf("Invalid format: %v, %v", n, err))
		}
	}

	return &day22{cubes: cubes}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: number of cubes on:", day.Part1())
	fmt.Println("ANSWER2: number of cubes on using all steps:", day.Part2())
}
