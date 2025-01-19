package day23

import (
	"strings"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

func (g *grid) set(c utils.Coord) {
	if _, ok := (*g)[c.Y]; !ok {
		(*g)[c.Y] = make(map[int]struct{})
	}
	(*g)[c.Y][c.X] = struct{}{}
}

func (g *grid) get(c utils.Coord) bool {
	if _, ok := (*g)[c.Y]; !ok {
		return false
	}
	_, ok := (*g)[c.Y][c.X]
	return ok
}

func (g *grid) delete(c utils.Coord) {
	if _, ok := (*g)[c.Y]; !ok {
		return
	}
	delete((*g)[c.Y], c.X)
}

type bounds struct {
	minX, maxX, minY, maxY int
}

func (g *grid) bounds() bounds {
	var minX, maxX, minY, maxY int

	for y, row := range *g {
		for x := range row {
			if y < minY {
				minY = y
			} else if y > maxY {
				maxY = y
			}

			if x < minX {
				minX = x
			} else if x > maxX {
				maxX = x
			}
		}
	}

	return bounds{minX, maxX, minY, maxY}
}

func (g grid) String() string {
	var (
		res    []string
		bounds = g.bounds()
		width  = bounds.maxX - bounds.minX + 1
	)

	for i := bounds.minY; i <= bounds.maxY; i++ {
		res = append(res, strings.Repeat(".", width))
	}

	for y, row := range g {
		for x := range row {
			i := y - bounds.minY
			j := x - bounds.minX
			res[i] = res[i][:j] + "#" + res[i][j+1:]
		}
	}

	return strings.Join(res, "\n")
}
