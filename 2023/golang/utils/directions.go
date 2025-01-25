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
