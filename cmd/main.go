package main

import (
	"image/color"
	"image/png"
	"os"
)

func main() {
	var colors = []color.Color{
		color.White,
		color.Black,
		color.RGBA{255, 0, 0, 255},
		color.RGBA{0, 0, 255, 255},
		color.RGBA{255, 255, 0, 255},
	}

	img := gomondrian.Generate(300, 200, 7, 350, 30)

	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
