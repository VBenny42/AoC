package utils

var (
	Up        = Coord{0, -1}
	Down      = Coord{0, 1}
	Left      = Coord{-1, 0}
	Right     = Coord{1, 0}
	UpLeft    = Coord{-1, -1}
	UpRight   = Coord{1, -1}
	DownLeft  = Coord{-1, 1}
	DownRight = Coord{1, 1}
)

var Directions = []Coord{Up, Down, Left, Right}

var AllDirections = []Coord{Up, Down, Left, Right, UpLeft, UpRight, DownLeft, DownRight}

var RotateRight = map[Coord]Coord{
	Up:    Right,
	Right: Down,
	Down:  Left,
	Left:  Up,
}

var RotateLeft = map[Coord]Coord{
	Up:    Left,
	Left:  Down,
	Down:  Right,
	Right: Up,
}

var Reverse = map[Coord]Coord{
	Up:    Down,
	Down:  Up,
	Left:  Right,
	Right: Left,
}
