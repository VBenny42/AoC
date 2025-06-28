package day20

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
)

func (d day20) ColorModel() color.Model {
	return color.RGBAModel
}

func (d day20) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(d.image[0]), len(d.image))
}

func (d day20) At(x, y int) color.Color {
	if d.image[y][x] == '#' {
		return color.Black
	}
	if d.image[y][x] == 'O' {
		return color.RGBA{255, 0, 0, 255}
	}
	return color.White
}

func (d day20) writePNG(filename string) error {
	base := filepath.Base(filename)
	base = base[:len(base)-len(filepath.Ext(base))]

	f, err := os.Create(base + ".png")
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, d)
}
