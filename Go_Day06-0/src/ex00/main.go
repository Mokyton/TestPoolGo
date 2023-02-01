package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	width := 300
	height := 300

	upLeft := image.Point{}
	lowRight := image.Point{X: width, Y: height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	//cyan := color.RGBA{100, 200, 200, 0xff}

	for x := 0; x < width; x++ {
		img.Set(x, 0, color.Black)
	}

	for y := 0; y < height; y++ {
		img.Set(width-1, y, color.Black)
	}

	for x := 0; x < width; x++ {
		img.Set(x, height-1, color.Black)
	}

	for y := 0; y < height; y++ {
		img.Set(0, y, color.Black)
	}

	for x := 1; x < width-1; x++ {
		for y := 1; y < height/2; y++ {
			img.Set(x, y, color.White)
		}
	}

	for x := 0; x < width-1; x++ {
		for y := height / 2; y < height-1; y++ {
			img.Set(x, y, color.RGBA{A: 0xff, R: 255, G: 0, B: 0})
		}
	}

	for x := 90; x < 120; x++ {
		for y := 90; y < 120; y++ {
			img.Set(x, y, color.Black)
		}
	}

	for x := 180; x < 210; x++ {
		for y := 90; y < 120; y++ {
			img.Set(x, y, color.Black)
		}
	}

	for x := 180; x < 210; x++ {
		for y := 180; y < 210; y++ {
			img.Set(x, y, color.Black)
		}
	}

	for x, y := 120, 120; x < 210; x, y = x+1, y+1 {
		img.Set(x, y, color.Black)
	}

	for x := 120; x < 180; x++ {
		img.Set(x, 90, color.Black)
	}

	for y := 120; y < 210; y++ {
		img.Set(209, y, color.Black)
	}

	for x, y := 150, 150; x < 182; x++ {

		img.Set(x, y, color.Black)
		y--
	}

	f, _ := os.Create("image_logo.png")
	png.Encode(f, img)
}
