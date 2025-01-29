package day17

import (
	"container/heap"
	"math"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

func (d *day17) dijkstra1() int {
	var (
		maxY = len(d.grid) - 1
		maxX = len(d.grid[0]) - 1
		end  = utils.Crd(maxX, maxY)
	)

	distance := make([][][][]int, maxX+1)
	for x := 0; x <= maxX; x++ {
		distance[x] = make([][][]int, maxY+1)
		for y := 0; y <= maxY; y++ {
			distance[x][y] = make([][]int, 4)
			for dir := 0; dir < 4; dir++ {
				distance[x][y][dir] = make([]int, 4)
				for steps := 0; steps < 4; steps++ {
					distance[x][y][dir][steps] = math.MaxInt
				}
			}
		}
	}

	dirs := []utils.Coord{utils.Up, utils.Right, utils.Down, utils.Left}

	pq := make(priorityQueue, 0)
	heap.Init(&pq)

	for i, dir := range dirs {
		if d.grid.inBounds(dir) {
			heatLoss := d.grid[dir.Y][dir.X]
			if heatLoss < distance[dir.X][dir.Y][i][1] {
				distance[dir.X][dir.Y][i][1] = heatLoss
				heap.Push(&pq, &state{
					pos:      dir,
					dir:      i,
					steps:    1,
					heatLoss: heatLoss,
				})
			}
		}
	}

	for pq.Len() > 0 {
		currentPtr := heap.Pop(&pq).(*state)
		current := *currentPtr

		if current.pos == end {
			return current.heatLoss
		}

		if current.heatLoss >
			distance[current.pos.X][current.pos.Y][current.dir][current.steps] {
			continue
		}

		for deltaDir := -1; deltaDir <= 1; deltaDir++ {
			nextDir := (current.dir + deltaDir + 4) % 4
			var newSteps int

			if nextDir == current.dir {
				newSteps = current.steps + 1
				if newSteps > 3 {
					continue
				}
			} else {
				newSteps = 1
			}

			next := current.pos.Add(dirs[nextDir])

			if !d.grid.inBounds(next) {
				continue
			}

			newHeat := current.heatLoss + d.grid[next.Y][next.X]

			if newHeat < distance[next.X][next.Y][nextDir][newSteps] {
				distance[next.X][next.Y][nextDir][newSteps] = newHeat
				heap.Push(&pq, &state{
					pos:      next,
					dir:      nextDir,
					steps:    newSteps,
					heatLoss: newHeat,
				})
			}
		}
	}

	return -1
}

func (d *day17) dijkstra2() int {
	var (
		maxY = len(d.grid) - 1
		maxX = len(d.grid[0]) - 1
		end  = utils.Crd(maxX, maxY)
	)

	distance := make([][][][]int, maxX+1)
	for x := 0; x <= maxX; x++ {
		distance[x] = make([][][]int, maxY+1)
		for y := 0; y <= maxY; y++ {
			distance[x][y] = make([][]int, 4)
			for dir := 0; dir < 4; dir++ {
				distance[x][y][dir] = make([]int, 11)
				for steps := 0; steps < 11; steps++ {
					distance[x][y][dir][steps] = math.MaxInt
				}
			}
		}
	}

	dirs := []utils.Coord{utils.Up, utils.Right, utils.Down, utils.Left}

	pq := make(priorityQueue, 0)
	heap.Init(&pq)

	// Push initial moves from start (0,0) to possible directions with step 1
	for dir := 0; dir < 4; dir++ {
		nextPos := utils.Crd(0, 0).Add(dirs[dir])
		if d.grid.inBounds(nextPos) {
			heatLoss := d.grid[nextPos.Y][nextPos.X]
			if heatLoss < distance[nextPos.X][nextPos.Y][dir][1] {
				distance[nextPos.X][nextPos.Y][dir][1] = heatLoss
				heap.Push(&pq, &state{
					pos:      nextPos,
					dir:      dir,
					steps:    1,
					heatLoss: heatLoss,
				})
			}
		}
	}

	for pq.Len() > 0 {
		currentPtr := heap.Pop(&pq).(*state)
		current := *currentPtr

		// Check if we reached the end with at least 4 steps in the current direction
		if current.pos == end && current.steps >= 4 {
			return current.heatLoss
		}

		// Skip if a better path has already been found
		if current.heatLoss >
			distance[current.pos.X][current.pos.Y][current.dir][current.steps] {
			continue
		}

		// Can go straight, left or right
		for deltaDir := -1; deltaDir <= 1; deltaDir++ {
			nextDir := (current.dir + deltaDir + 4) % 4
			var newSteps int

			if deltaDir == 0 {
				// Moving straight
				newSteps = current.steps + 1
				if newSteps > 10 {
					continue
				}
			} else {
				// Turning left or right, only allowed if steps >=4
				if current.steps < 4 {
					continue
				}
				newSteps = 1
			}

			nextPos := current.pos.Add(dirs[nextDir])
			if !d.grid.inBounds(nextPos) {
				continue
			}

			newHeat := current.heatLoss + d.grid[nextPos.Y][nextPos.X]

			if newHeat < distance[nextPos.X][nextPos.Y][nextDir][newSteps] {
				distance[nextPos.X][nextPos.Y][nextDir][newSteps] = newHeat
				heap.Push(&pq, &state{
					pos:      nextPos,
					dir:      nextDir,
					steps:    newSteps,
					heatLoss: newHeat,
				})
			}
		}
	}

	return -1
}
