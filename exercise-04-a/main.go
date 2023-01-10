package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"net/http"
	"strconv"
)

func createFrame(width int, height int, offset int, palette color.Palette) *image.Paletted {
	var scale = (height - 1) / 2
	var img = image.NewPaletted(image.Rect(0, 0, width, height), palette)

	for i := 0; i < width; i += 1 {
		var sin = math.Sin(float64(i+offset)/100.0) * float64(scale)
		var y = int(sin) + scale
		img.SetColorIndex(i, y-1, 1)
		img.SetColorIndex(i, y, 1)
		img.SetColorIndex(i, y+1, 1)
	}

	return img
}

func buildGif(w io.Writer, width int, height int, delay int, frames int, speed int) {
	var palette = []color.Color{
		color.RGBA{255, 255, 255, 255},
		color.RGBA{255, 0, 0, 255},
	}

	animation := gif.GIF{Delay: []int{}, Image: []*image.Paletted{}}

	for i := 0; i < frames; i += 1 {
		animation.Image = append(animation.Image, createFrame(width, height, i*speed, palette))
		animation.Delay = append(animation.Delay, delay)
	}

	gif.EncodeAll(w, &animation)
}

func getInt(val string, defaultValue int) int {
	intval, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}

	return intval
}

func main() {
	http.HandleFunc("/sin", func(w http.ResponseWriter, r *http.Request) {
		val := r.URL.Query().Get("val")
		intval, _ := strconv.Atoi(val)
		result := math.Sin(float64(intval))
		fmt.Fprintf(w, "Sin of %v is %v", val, result)
	})

	http.HandleFunc("/gif", func(w http.ResponseWriter, r *http.Request) {
		width := getInt(r.URL.Query().Get("width"), 1000)
		height := getInt(r.URL.Query().Get("height"), 1000)
		delay := getInt(r.URL.Query().Get("delay"), 5)
		frames := getInt(r.URL.Query().Get("frames"), 1000)
		speed := getInt(r.URL.Query().Get("speed"), 1)

		w.Header().Add("Content-Type", "image/gif")
		buildGif(w, width, height, delay, frames, speed)
	})

	port := ":3000"
	fmt.Printf("[INFO] Server listening on port %v", port)

	http.ListenAndServe(port, nil)
}
