package day13

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

type grid utils.Grid[rune]

func (g grid) ColorModel() color.Model {
	return color.GrayModel
}

func (g grid) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(g[0]), len(g))
}

func (g grid) At(x, y int) color.Color {
	if g[y][x] == '#' {
		return color.Black
	}
	return color.White
}

func (g grid) writeImage(filename string) error {
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
