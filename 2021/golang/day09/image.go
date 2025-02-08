package day09

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

type grid utils.Grid[int]

func (g grid) ColorModel() color.Model {
	return color.GrayModel
}

func (g grid) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(g[0]), len(g))
}

func (g grid) At(x, y int) color.Color {
	// lowest point is white, 9's are black
	alpha := uint8(g[y][x] * 255 / 9)
	return color.Gray{255 - alpha}
}

func (g grid) writeImage(filename string) {
	base := filepath.Base(filename)
	filename = "day09-" + base[:len(base)-4] + ".png"

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := png.Encode(f, g); err != nil {
		panic(err)
	}
}
