package utils

import "fmt"

type Coord struct {
	X, Y int
}

func (c Coord) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

func (c Coord) Add(other Coord) Coord {
	return Coord{c.X + other.X, c.Y + other.Y}
}

func (c Coord) Sub(other Coord) Coord {
	return Coord{c.X - other.X, c.Y - other.Y}
}
