package day05

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func (g grid) ColorModel() color.Model {
	return color.Gray16Model
}

func (g grid) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(g[0]), len(g))
}

func (g grid) At(x, y int) color.Color {
	count := g[y][x]
	if count > 1 {
		return color.Black
	}
	alpha := uint8(count * 255 / 6)
	return color.Gray{255 - alpha}
}

func (g grid) writeImage(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := png.Encode(f, g); err != nil {
		panic(err)
	}
}
