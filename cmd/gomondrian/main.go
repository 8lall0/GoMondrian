package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"

	"github.com/8lall0/GoMondrian"
)

func main() {
	width := flag.Int("x", 640, "set width")
	height := flag.Int("y", 480, "set height")
	padding := flag.Int("p", 10, "set padding")
	divisions := flag.Int("d", 50, "set number of divisions")
	colors := flag.Int("c", 30, "set number of colored squares")
	filename := flag.String("s", "out.png", "set file output name")

	flag.Parse()

	img, err := gomondrian.Generate(*width, *height, *padding, *divisions, *colors)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f, _ := os.OpenFile(*filename, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
