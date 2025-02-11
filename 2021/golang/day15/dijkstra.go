package day15

import (
	"container/heap"
	"image"
	"math"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

func (d *day15) dijkstra() int {
	var (
		maxX = len(d.grid[0]) - 1
		maxY = len(d.grid) - 1

		end = image.Pt(maxX, maxY)
	)

	distance := utils.NewGrid[int](maxX+1, maxY+1)
	for y := range maxY + 1 {
		for x := range maxX + 1 {
			distance[y][x] = math.MaxInt
		}
	}

	var pq priorityQueue
	heap.Init(&pq)

	heap.Push(&pq, &state{
		pos:  utils.Down,
		risk: d.grid.Get(utils.Down),
	})
	heap.Push(&pq, &state{
		pos:  utils.Right,
		risk: d.grid.Get(utils.Right),
	})

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*state)

		if current.pos == end {
			return current.risk
		}

		if current.risk > distance[current.pos.Y][current.pos.X] {
			continue
		}

		for _, dir := range utils.Directions {
			next := current.pos.Add(dir)
			if !d.grid.InBounds(next) {
				continue
			}

			risk := d.grid[next.Y][next.X] + current.risk
			if risk < distance[next.Y][next.X] {
				distance[next.Y][next.X] = risk
				heap.Push(&pq, &state{
					pos:  next,
					risk: risk,
				})
			}
		}
	}

	return -1
}
