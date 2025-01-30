package day18

import (
	"image"
	"image/color"
	"os"

	"golang.org/x/image/bmp"
)

func (g grid) ColorModel() color.Model {
	return color.RGBAModel
}

func (g grid) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(g[0]), len(g))
}

func (g grid) At(x, y int) color.Color {
	return g[y][x]
}

func (g grid) writeBitmap(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := bmp.Encode(f, g); err != nil {
		return err
	}

	return nil
}
