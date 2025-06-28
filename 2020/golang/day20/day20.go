package day20

import (
	"fmt"
	"math"

	"github.com/VBenny42/AoC/2020/golang/utils"
)

type tile struct {
	id    int
	data  []string
	edges [4]string // 0: top, 1: right, 2: bottom, 3: left
}

type day20 struct {
	tiles           map[int]*tile
	corners         []*tile
	gridSize        int
	image           []string
	monster         []string
	monsterW        int
	monsterH        int
	hashesInMonster int
}

func (d *day20) Part1() int {
	cornerProduct := 1
	edgeMatches := make(map[string]int)

	for _, t := range d.tiles {
		for _, edge := range t.edges {
			edgeMatches[edge]++
			edgeMatches[utils.Reverse(edge)]++
		}
	}

	for _, t := range d.tiles {
		unmatchedEdges := 0
		for _, edge := range t.edges {
			if edgeMatches[edge] == 1 {
				unmatchedEdges++
			}
		}
		if unmatchedEdges == 2 {
			d.corners = append(d.corners, t)
			cornerProduct *= t.id
		}
	}

	return cornerProduct
}

func (d *day20) Part2() int {
	grid := d.assembleGrid()
	if grid == nil {
		fmt.Println("Failed to assemble grid.")
		return 0
	}

	d.createImage(grid)

	d.defineMonster()

	for flip := 0; flip < 2; flip++ {
		for rot := 0; rot < 4; rot++ {
			monsterCount := d.findMonsters()
			if monsterCount > 0 {

				d.markMonsters()
				waterRoughness := 0
				for _, row := range d.image {
					for _, char := range row {
						if char == '#' {
							waterRoughness++
						}
					}
				}
				return waterRoughness
			}
			d.rotateImage()
		}
		d.flipImage()
	}

	return 0
}

func (d *day20) assembleGrid() [][]*tile {
	d.gridSize = int(math.Sqrt(float64(len(d.tiles))))
	grid := make([][]*tile, d.gridSize)
	for i := range grid {
		grid[i] = make([]*tile, d.gridSize)
	}

	startTile := d.corners[0]

	edgeMatches := make(map[string]int)
	for _, t := range d.tiles {
		if t.id == startTile.id {
			continue
		}
		for _, edge := range t.edges {
			edgeMatches[edge]++
			edgeMatches[utils.Reverse(edge)]++
		}
	}

	// Rotate the corner until its non-matching edges are top (0) and left (3).
	for rot := 0; rot < 4; rot++ {
		topEdge := startTile.edges[0]
		leftEdge := startTile.edges[3]
		if edgeMatches[topEdge] == 0 && edgeMatches[leftEdge] == 0 {
			break
		}
		startTile.rotate()
	}

	grid[0][0] = startTile
	used := map[int]bool{startTile.id: true}

	for r := 0; r < d.gridSize; r++ {
		for c := 0; c < d.gridSize; c++ {
			if r == 0 && c == 0 {
				continue // Skip the starting tile
			}

			var prevTile *tile
			var prevEdgeIndex, currentEdgeIndex int

			if c > 0 {
				prevTile = grid[r][c-1]
				prevEdgeIndex = 1
				currentEdgeIndex = 3
			} else {
				prevTile = grid[r-1][c]
				prevEdgeIndex = 2
				currentEdgeIndex = 0
			}

			targetEdge := prevTile.edges[prevEdgeIndex]

			found := false
			for _, t := range d.tiles {
				if used[t.id] {
					continue
				}

				for flip := 0; flip < 2; flip++ {
					for rot := 0; rot < 4; rot++ {
						if t.edges[currentEdgeIndex] == targetEdge {
							grid[r][c] = t
							used[t.id] = true
							found = true
							goto nextTile
						}
						t.rotate()
					}
					t.flipHorizontally()
				}
			}
		nextTile:
			if !found {
				return nil
			}
		}
	}
	return grid
}

func (d *day20) createImage(grid [][]*tile) {
	tileSizeWithoutBorder := len(grid[0][0].data) - 2
	imageSize := d.gridSize * tileSizeWithoutBorder
	d.image = make([]string, imageSize)

	for r := 0; r < d.gridSize; r++ {
		for c := 0; c < d.gridSize; c++ {
			tileData := grid[r][c].data
			for i := 1; i < len(tileData)-1; i++ {
				imageRow := r*tileSizeWithoutBorder + (i - 1)
				d.image[imageRow] += tileData[i][1 : len(tileData[i])-1]
			}
		}
	}
}

func (d *day20) defineMonster() {
	d.monster = []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}
	d.monsterH = len(d.monster)
	d.monsterW = len(d.monster[0])
	d.hashesInMonster = 0
	for _, row := range d.monster {
		for _, char := range row {
			if char == '#' {
				d.hashesInMonster++
			}
		}
	}
}

func (d *day20) findMonsters() int {
	count := 0
	imageH := len(d.image)
	imageW := len(d.image[0])

	for r := 0; r <= imageH-d.monsterH; r++ {
		for c := 0; c <= imageW-d.monsterW; c++ {
			isMonster := true
			for mr := 0; mr < d.monsterH; mr++ {
				for mc := 0; mc < d.monsterW; mc++ {
					if d.monster[mr][mc] == '#' && d.image[r+mr][c+mc] != '#' {
						isMonster = false
						break
					}
				}
				if !isMonster {
					break
				}
			}
			if isMonster {
				count++
			}
		}
	}
	return count
}

func (d *day20) markMonsters() {
	imageH := len(d.image)
	if imageH == 0 {
		return
	}
	imageW := len(d.image[0])

	imageRunes := make([][]rune, imageH)
	for i, row := range d.image {
		imageRunes[i] = []rune(row)
	}

	for r := 0; r <= imageH-d.monsterH; r++ {
		for c := 0; c <= imageW-d.monsterW; c++ {
			isMonster := true
			for mr := 0; mr < d.monsterH; mr++ {
				for mc := 0; mc < d.monsterW; mc++ {
					if d.monster[mr][mc] == '#' && d.image[r+mr][c+mc] != '#' {
						isMonster = false
						break
					}
				}
				if !isMonster {
					break
				}
			}

			// If a monster is found, mark its parts with 'O' on the rune grid.
			if isMonster {
				for mr := 0; mr < d.monsterH; mr++ {
					for mc := 0; mc < d.monsterW; mc++ {
						if d.monster[mr][mc] == '#' {
							imageRunes[r+mr][c+mc] = 'O'
						}
					}
				}
			}
		}
	}

	for i, rowRunes := range imageRunes {
		d.image[i] = string(rowRunes)
	}
}

func (d *day20) rotateImage() {
	size := len(d.image)
	newImage := make([]string, size)
	for i := range newImage {
		newImage[i] = ""
	}
	for i := size - 1; i >= 0; i-- {
		for j := 0; j < size; j++ {
			newImage[j] += string(d.image[i][j])
		}
	}
	d.image = newImage
}

func (d *day20) flipImage() {
	for i, row := range d.image {
		d.image[i] = utils.Reverse(row)
	}
}

func parseTile(lines []string) *tile {
	id := utils.Atoi(lines[0][5:9])
	data := lines[1:]
	t := &tile{id: id, data: data}
	t.updateEdges()
	return t
}

func Parse(filename string) *day20 {
	lines := utils.ReadLines(filename)
	tiles := make(map[int]*tile)

	for i := 0; i < len(lines); {
		if lines[i] == "" {
			i++
			continue
		}

		tileLines := []string{lines[i]}
		i++
		for i < len(lines) && lines[i] != "" {
			tileLines = append(tileLines, lines[i])
			i++
		}

		t := parseTile(tileLines)
		tiles[t.id] = t
	}
	return &day20{tiles: tiles}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: product of the IDs of the 4 corner tiles:", day.Part1())
	fmt.Println("ANSWER2: water roughness after removing sea monsters:", day.Part2())

	// if err := day.writePNG(filename); err != nil {
	// 	fmt.Println("Error writing image:", err)
	// }
}
