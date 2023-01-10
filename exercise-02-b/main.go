package main

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"os"
)

func createFrame(size int, offset int, palette color.Palette) *image.Paletted {
	var scale = (size - 1) / 2
	var img = image.NewPaletted(image.Rect(0, 0, size, size), palette)

	for i := 0; i < size; i += 1 {
		var sin = math.Sin(float64(i+offset)/100.0) * float64(scale)
		var y = int(sin) + scale
		img.SetColorIndex(i, y-1, 1)
		img.SetColorIndex(i, y, 1)
		img.SetColorIndex(i, y+1, 1)
	}

	return img
}

func main() {
	var size = 1000

	var palette = []color.Color{
		color.RGBA{255, 255, 255, 255},
		color.RGBA{255, 0, 0, 255},
	}

	animation := gif.GIF{Delay: []int{}, Image: []*image.Paletted{}}

	for i := 0; i < 1000; i += 1 {
		animation.Image = append(animation.Image, createFrame(size, i, palette))
		animation.Delay = append(animation.Delay, 5)
	}

	f, err := os.Create("my-image.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	gif.EncodeAll(f, &animation)
}
