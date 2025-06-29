package utils

import "image"

var (
	Up        = image.Pt(0, -1)
	Down      = image.Pt(0, 1)
	Left      = image.Pt(-1, 0)
	Right     = image.Pt(1, 0)
	UpLeft    = image.Pt(-1, -1)
	UpRight   = image.Pt(1, -1)
	DownLeft  = image.Pt(-1, 1)
	DownRight = image.Pt(1, 1)
)

var Directions = []image.Point{Up, Down, Left, Right}

var AllDirections = []image.Point{Up, Down, Left, Right, UpLeft, UpRight, DownLeft, DownRight}

var RotateRight = map[image.Point]image.Point{
	Up:    Right,
	Right: Down,
	Down:  Left,
	Left:  Up,
}

var RotateLeft = map[image.Point]image.Point{
	Up:    Left,
	Left:  Down,
	Down:  Right,
	Right: Up,
}
