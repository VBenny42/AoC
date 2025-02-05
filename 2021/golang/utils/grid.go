package utils

import "image"

type Grid[T any] [][]T

func NewGrid[T any](width, height int) Grid[T] {
	data := make([][]T, height)
	for i := range data {
		data[i] = make([]T, width)
	}
	return data
}

func (g *Grid[T]) Get(p image.Point) T {
	return (*g)[p.Y][p.X]
}

func (g *Grid[T]) inBounds(p image.Point) bool {
	return p.X >= 0 && p.X < len((*g)[0]) &&
		p.Y >= 0 && p.Y < len((*g))
}
