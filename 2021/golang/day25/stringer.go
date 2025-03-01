package day25

import (
	"fmt"
	"image"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

func printGrid(grid utils.Grid[rune]) {
	for y := 0; y < gridHeight; y++ {
		for x := 0; x < gridWidth; x++ {
			fmt.Printf("%c", grid.Get(image.Pt(x, y)))
		}
		fmt.Println()
	}
}
