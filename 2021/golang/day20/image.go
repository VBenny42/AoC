package day20

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

type imageGrid utils.Grid[pixel]

func (g imageGrid) ColorModel() color.Model {
	return color.GrayModel
}

func (g imageGrid) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(g[0]), len(g))
}

func (g imageGrid) At(x, y int) color.Color {
	if g[y][x] == dark {
		return color.Black
	}
	return color.White
}

func (g imageGrid) writeImage(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := png.Encode(f, g); err != nil {
		return err
	}

	return nil
}
