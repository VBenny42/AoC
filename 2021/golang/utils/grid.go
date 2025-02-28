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

func (g *Grid[T]) Set(p image.Point, val T) {
	(*g)[p.Y][p.X] = val
}

func (g *Grid[T]) InBounds(p image.Point) bool {
	return p.X >= 0 && p.X < len((*g)[0]) &&
		p.Y >= 0 && p.Y < len((*g))
}

func (g *Grid[T]) Neighbors(p image.Point) (neighbors []image.Point) {
	for _, d := range Directions {
		neighbor := p.Add(d)
		if g.InBounds(neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}

	return
}

func (g *Grid[T]) Clone() Grid[T] {
	clone := NewGrid[T](len((*g)[0]), len((*g)))
	for y := range *g {
		copy(clone[y], (*g)[y])
	}
	return clone
}
