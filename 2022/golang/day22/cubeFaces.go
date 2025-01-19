package day22

import (
	"github.com/VBenny42/AoC/2022/golang/utils"
)

var sampleFaces = map[face]squareBounds{
	one: {
		row: bounds{start: 8, end: 11},
		col: bounds{start: 0, end: 3},
	},
	two: {
		row: bounds{start: 0, end: 3},
		col: bounds{start: 4, end: 7},
	},
	three: {
		row: bounds{start: 4, end: 7},
		col: bounds{start: 4, end: 7},
	},
	four: {
		row: bounds{start: 8, end: 11},
		col: bounds{start: 4, end: 7},
	},
	five: {
		row: bounds{start: 8, end: 11},
		col: bounds{start: 8, end: 11},
	},
	six: {
		row: bounds{start: 12, end: 15},
		col: bounds{start: 8, end: 11},
	},
}

type (
	newPosition struct {
		face      face
		direction direction
		newPos    func(p utils.Coord) (newP utils.Coord)
	}
	transition map[direction]newPosition
)

var sampleFaceTransitions = map[face]transition{
	one: {
		right: {
			face:      six,
			direction: left,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: sampleFaces[six][row].end,
					Y: utils.Abs(sampleFaces[one][col].end-p.Y) + sampleFaces[six][col].start,
				}
			},
		},
		left: {
			face:      three,
			direction: down,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(sampleFaces[one][col].start-p.Y) + sampleFaces[three][row].start,
					Y: sampleFaces[three][col].start,
				}
			},
		},
		up: {
			face:      two,
			direction: down,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(sampleFaces[one][row].end-p.X) + sampleFaces[two][row].start,
					Y: sampleFaces[two][col].start,
				}
			},
		},
		down: {
			face:      four,
			direction: down,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(sampleFaces[one][row].start-p.X) + sampleFaces[four][row].start,
					Y: sampleFaces[four][col].start,
				}
			},
		},
	},
	two: {
		right: {
			face:      three,
			direction: right,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: sampleFaces[three][row].start,
					Y: utils.Abs(sampleFaces[two][col].start-p.Y) + sampleFaces[three][col].start,
				}
			},
		},
		left: {
			face:      six,
			direction: up,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(sampleFaces[two][col].end-p.Y) + sampleFaces[six][row].start,
					Y: sampleFaces[six][col].end,
				}
			},
		},
		up: {
			face:      one,
			direction: down,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(sampleFaces[two][row].end-p.X) + sampleFaces[one][row].start,
					Y: sampleFaces[one][col].start,
				}
			},
		},
		down: {
			face:      five,
			direction: up,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(sampleFaces[two][row].end-p.X) + sampleFaces[five][row].start,
					Y: sampleFaces[five][col].end,
				}
			},
		},
	},
	three: {
		right: {
			face:      four,
			direction: right,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: sampleFaces[four][row].start,
					Y: utils.Abs(sampleFaces[three][col].start-p.Y) + sampleFaces[four][col].start,
				}
			},
		},
		left: {
			face:      two,
			direction: left,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: sampleFaces[two][row].end,
					Y: utils.Abs(sampleFaces[three][col].start-p.Y) + sampleFaces[two][col].start,
				}
			},
		},
		up: {
			face:      one,
			direction: right,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: sampleFaces[one][row].start,
					Y: utils.Abs(sampleFaces[three][row].start-p.X) + sampleFaces[one][col].start,
				}
			},
		},
		down: {
			face:      five,
			direction: right,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: sampleFaces[five][row].start,
					Y: utils.Abs(sampleFaces[three][row].end-p.X) + sampleFaces[five][col].start,
				}
			},
		},
	},
	four: {
		right: {
			face:      six,
			direction: down,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(sampleFaces[four][col].end-p.Y) + sampleFaces[six][row].start,
					Y: sampleFaces[six][col].start,
				}
			},
		},
		left: {
			face:      three,
			direction: left,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: sampleFaces[three][row].end,
					Y: utils.Abs(sampleFaces[four][col].start-p.Y) + sampleFaces[three][col].start,
				}
			},
		},
		up: {
			face:      one,
			direction: up,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(sampleFaces[four][row].start-p.X) + sampleFaces[one][row].start,
					Y: sampleFaces[one][col].end,
				}
			},
		},
		down: {
			face:      five,
			direction: down,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(sampleFaces[four][row].start-p.X) + sampleFaces[five][row].start,
					Y: sampleFaces[five][col].start,
				}
			},
		},
	},
	five: {
		right: {
			face:      six,
			direction: right,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: sampleFaces[six][row].start,
					Y: utils.Abs(sampleFaces[five][col].start-p.Y) + sampleFaces[six][col].start,
				}
			},
		},
		left: {
			face:      three,
			direction: up,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(sampleFaces[five][col].end-p.Y) + sampleFaces[three][row].start,
					Y: sampleFaces[three][col].end,
				}
			},
		},
		up: {
			face:      four,
			direction: up,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(sampleFaces[five][row].start-p.X) + sampleFaces[four][row].start,
					Y: sampleFaces[four][col].end,
				}
			},
		},
		down: {
			face:      two,
			direction: up,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(sampleFaces[five][row].end-p.X) + sampleFaces[two][row].start,
					Y: sampleFaces[two][col].end,
				}
			},
		},
	},
	six: {
		right: {
			face:      one,
			direction: left,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: sampleFaces[one][row].end,
					Y: utils.Abs(sampleFaces[six][col].end-p.Y) + sampleFaces[one][col].start,
				}
			},
		},
		left: {
			face:      five,
			direction: left,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: sampleFaces[five][row].end,
					Y: utils.Abs(sampleFaces[six][col].start-p.Y) + sampleFaces[five][col].start,
				}
			},
		},
		up: {
			face:      four,
			direction: left,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: sampleFaces[four][row].end,
					Y: utils.Abs(sampleFaces[six][row].start-p.X) + sampleFaces[four][col].start,
				}
			},
		},
		down: {
			face:      two,
			direction: right,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: sampleFaces[two][row].start,
					Y: utils.Abs(sampleFaces[six][row].end-p.X) + sampleFaces[two][col].start,
				}
			},
		},
	},
}

var inputFaces = map[face]squareBounds{
	one: {
		row: bounds{start: 50, end: 99},
		col: bounds{start: 0, end: 49},
	},
	two: {
		row: bounds{start: 100, end: 149},
		col: bounds{start: 0, end: 49},
	},
	three: {
		row: bounds{start: 50, end: 99},
		col: bounds{start: 50, end: 99},
	},
	four: {
		row: bounds{start: 0, end: 49},
		col: bounds{start: 100, end: 149},
	},
	five: {
		row: bounds{start: 50, end: 99},
		col: bounds{start: 100, end: 149},
	},
	six: {
		row: bounds{start: 0, end: 49},
		col: bounds{start: 150, end: 199},
	},
}

var inputFaceTransitions = map[face]transition{
	one: {
		right: {
			face:      two,
			direction: right,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: inputFaces[two][row].start,
					Y: utils.Abs(inputFaces[one][col].start-p.Y) + inputFaces[two][col].start,
				}
			},
		},
		left: {
			face:      four,
			direction: right,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: inputFaces[four][row].start,
					Y: utils.Abs(inputFaces[one][col].end-p.Y) + inputFaces[four][col].start,
				}
			},
		},
		up: {
			face:      six,
			direction: right,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: inputFaces[six][row].start,
					Y: utils.Abs(inputFaces[one][row].start-p.X) + inputFaces[six][col].start,
				}
			},
		},
		down: {
			face:      three,
			direction: down,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(inputFaces[one][row].start-p.X) + inputFaces[three][row].start,
					Y: inputFaces[three][col].start,
				}
			},
		},
	},
	two: {
		right: {
			face:      five,
			direction: left,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: inputFaces[five][row].end,
					Y: utils.Abs(inputFaces[two][col].end-p.Y) + inputFaces[five][col].start,
				}
			},
		},
		left: {
			face:      one,
			direction: left,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: inputFaces[one][row].end,
					Y: utils.Abs(inputFaces[two][col].start-p.Y) + inputFaces[one][col].start,
				}
			},
		},
		up: {
			face:      six,
			direction: up,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(inputFaces[two][row].start-p.X) + inputFaces[six][row].start,
					Y: inputFaces[six][col].end,
				}
			},
		},
		down: {
			face:      three,
			direction: left,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: inputFaces[three][row].end,
					Y: utils.Abs(inputFaces[two][row].start-p.X) + inputFaces[three][col].start,
				}
			},
		},
	},
	three: {
		right: {
			face:      two,
			direction: up,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(inputFaces[three][col].start-p.Y) + inputFaces[two][row].start,
					Y: inputFaces[two][col].end,
				}
			},
		},
		left: {
			face:      four,
			direction: down,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(inputFaces[three][col].start-p.Y) + inputFaces[four][row].start,
					Y: inputFaces[four][col].start,
				}
			},
		},
		up: {
			face:      one,
			direction: up,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(inputFaces[three][row].start-p.X) + inputFaces[one][row].start,
					Y: inputFaces[one][col].end,
				}
			},
		},
		down: {
			face:      five,
			direction: down,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(inputFaces[three][row].start-p.X) + inputFaces[five][row].start,
					Y: inputFaces[five][col].start,
				}
			},
		},
	},
	four: {
		right: {
			face:      five,
			direction: right,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: inputFaces[five][row].start,
					Y: utils.Abs(inputFaces[four][col].start-p.Y) + inputFaces[five][col].start,
				}
			},
		},
		left: {
			face:      one,
			direction: right,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: inputFaces[one][row].start,
					Y: utils.Abs(inputFaces[four][col].end-p.Y) + inputFaces[one][col].start,
				}
			},
		},
		up: {
			face:      three,
			direction: right,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: inputFaces[three][row].start,
					Y: utils.Abs(inputFaces[four][row].start-p.X) + inputFaces[three][col].start,
				}
			},
		},
		down: {
			face:      six,
			direction: down,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(inputFaces[four][row].start-p.X) + inputFaces[six][row].start,
					Y: inputFaces[six][col].start,
				}
			},
		},
	},
	five: {
		right: {
			face:      two,
			direction: left,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: inputFaces[two][row].end,
					Y: utils.Abs(inputFaces[five][col].end-p.Y) + inputFaces[two][col].start,
				}
			},
		},
		left: {
			face:      four,
			direction: left,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: inputFaces[four][row].end,
					Y: utils.Abs(inputFaces[five][col].start-p.Y) + inputFaces[four][col].start,
				}
			},
		},
		up: {
			face:      three,
			direction: up,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(inputFaces[five][row].start-p.X) + inputFaces[three][row].start,
					Y: inputFaces[three][col].end,
				}
			},
		},
		down: {
			face:      six,
			direction: left,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: inputFaces[six][row].end,
					Y: utils.Abs(inputFaces[five][row].start-p.X) + inputFaces[six][col].start,
				}
			},
		},
	},
	six: {
		right: {
			face:      five,
			direction: up,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(inputFaces[six][col].start-p.Y) + inputFaces[five][row].start,
					Y: inputFaces[five][col].end,
				}
			},
		},
		left: {
			face:      one,
			direction: down,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(inputFaces[six][col].start-p.Y) + inputFaces[one][row].start,
					Y: inputFaces[one][col].start,
				}
			},
		},
		up: {
			face:      four,
			direction: up,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(inputFaces[six][row].start-p.X) + inputFaces[four][row].start,
					Y: inputFaces[four][col].end,
				}
			},
		},
		down: {
			face:      two,
			direction: down,
			newPos: func(p utils.Coord) utils.Coord {
				return utils.Coord{
					X: utils.Abs(inputFaces[six][row].start-p.X) + inputFaces[two][row].start,
					Y: inputFaces[two][col].start,
				}
			},
		},
	},
}
