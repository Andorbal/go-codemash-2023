package main

import (
	"image"
	"image/color"
	"image/gif"
	"os"
)

func main() {
	var size = 1000
	var checker = 100

	var rect = image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: size,
			Y: size,
		},
	}

	var red = color.RGBA{
		R: 255,
		A: 255,
	}
	var blue = color.RGBA{
		B: 255,
		A: 255,
	}

	var palette = []color.Color{
		red, blue,
	}

	var img = image.NewPaletted(rect, palette)
	var currentRowColor = uint8(0)
	checkerboard := gif.GIF{Delay: []int{0}, Image: []*image.Paletted{img}}

	for y := 0; y < size; y += 1 {
		var currentColumnColor = currentRowColor

		for x := 0; x < size; x += 1 {
			img.SetColorIndex(x, y, currentColumnColor)

			if x%checker == 0 {
				if currentColumnColor == 0 {
					currentColumnColor = 1
				} else {
					currentColumnColor = 0
				}
			}
		}

		if y%checker == 0 {
			if currentRowColor == 0 {
				currentRowColor = 1
			} else {
				currentRowColor = 0
			}
		}
	}

	f, err := os.Create("my-image.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	gif.EncodeAll(f, &checkerboard)
}
