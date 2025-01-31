package day21

import (
	"image"
	"image/color"
	"os"

	"golang.org/x/image/bmp"
)

func (g grid) ColorModel() color.Model {
	return color.GrayModel
}

func (g grid) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(g[0]), len(g))
}

func (g grid) At(x, y int) color.Color {
	if y == 65 && x == 65 {
		return color.RGBA{255, 0, 0, 255}
	}
	if g[y][x] == plot {
		return color.Black
	}
	return color.White
}

func (g grid) writeBitmap(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return bmp.Encode(f, g)
}
