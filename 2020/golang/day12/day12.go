package day12

import (
	"fmt"
	"image"
	"slices"

	"github.com/VBenny42/AoC/2020/golang/utils"
)

type instruction struct {
	opcode   string
	argument int
}

const (
	north   = "N"
	south   = "S"
	east    = "E"
	west    = "W"
	left    = "L"
	right   = "R"
	forward = "F"
)

type day12 struct {
	instructions []instruction

	point     image.Point
	waypoint  image.Point
	direction string
}

func (d *day12) moveForward(steps int) {
	switch d.direction {
	case north:
		d.point.Y += steps
	case south:
		d.point.Y -= steps
	case east:
		d.point.X += steps
	case west:
		d.point.X -= steps
	default:
		fmt.Printf("ERROR: Invalid direction %s\n", d.direction)
	}
	return
}

func (d *day12) processInstruction(i instruction) {
	directions := []string{north, east, south, west}

	turnLeft := func(current string, degrees int) string {
		steps := (degrees / 90) % 4

		currentIndex := slices.Index(directions, current)
		if currentIndex == -1 {
			fmt.Printf("ERROR: Invalid direction %s\n", current)
			return current
		}

		newIndex := (currentIndex - steps + len(directions)) % len(directions)

		return directions[newIndex]
	}

	turnRight := func(current string, degrees int) string {
		steps := (degrees / 90) % 4
		currentIndex := slices.Index(directions, current)
		if currentIndex == -1 {
			fmt.Printf("ERROR: Invalid direction %s\n", current)
			return current
		}

		newIndex := (currentIndex + steps) % len(directions)
		return directions[newIndex]
	}

	switch i.opcode {
	case north:
		d.point.Y += i.argument
	case south:
		d.point.Y -= i.argument
	case east:
		d.point.X += i.argument
	case west:
		d.point.X -= i.argument
	case left:
		d.direction = turnLeft(d.direction, i.argument)
	case right:
		d.direction = turnRight(d.direction, i.argument)
	case forward:
		d.moveForward(i.argument)
	default:
		fmt.Printf("ERROR: Unknown opcode %s\n", i.opcode)
	}
}

func (d *day12) moveToWaypoint(steps int) {
	d.point.X += d.waypoint.X * steps
	d.point.Y += d.waypoint.Y * steps
}

func (d *day12) rotateWaypointLeft(degrees int) {
	steps := (degrees / 90) % 4
	for i := 0; i < steps; i++ {
		newX := -d.waypoint.Y
		newY := d.waypoint.X
		d.waypoint.X = newX
		d.waypoint.Y = newY
	}
}

func (d *day12) rotateWaypointRight(degrees int) {
	steps := (degrees / 90) % 4
	for i := 0; i < steps; i++ {
		newX := d.waypoint.Y
		newY := -d.waypoint.X
		d.waypoint.X = newX
		d.waypoint.Y = newY
	}
}

func (d *day12) processInstruction2(i instruction) {
	switch i.opcode {
	case north:
		d.waypoint.Y += i.argument
	case south:
		d.waypoint.Y -= i.argument
	case east:
		d.waypoint.X += i.argument
	case west:
		d.waypoint.X -= i.argument
	case left:
		d.rotateWaypointLeft(i.argument)
	case right:
		d.rotateWaypointRight(i.argument)
	case forward:
		d.moveToWaypoint(i.argument)
	default:
		fmt.Printf("ERROR: Unknown opcode %s\n", i.opcode)
	}
}

func (d *day12) Part1() int {
	for _, instruction := range d.instructions {
		d.processInstruction(instruction)
	}

	return utils.Abs(d.point.X) + utils.Abs(d.point.Y)
}

func (d *day12) Part2() int {
	d.point = image.Point{X: 0, Y: 0}
	d.waypoint = image.Point{X: 10, Y: 1}
	d.direction = east

	for _, instruction := range d.instructions {
		d.processInstruction2(instruction)
	}

	return utils.Abs(d.point.X) + utils.Abs(d.point.Y)
}

func Parse(filename string) *day12 {
	var (
		lines        = utils.ReadLines(filename)
		instructions = make([]instruction, 0, len(lines))

		opcode   string
		argument int
	)

	for _, line := range lines {
		opcode = line[:1]
		argument = utils.Atoi(line[1:])
		instructions = append(instructions, instruction{opcode: opcode, argument: argument})
	}

	return &day12{
		instructions: instructions,
		point:        image.Point{X: 0, Y: 0},
		direction:    east,
	}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: manhattan difference between ending and starting position:", day.Part1())
	fmt.Println("ANSWER2: manhattan difference between ending and starting position with waypoint:", day.Part2())
}
