package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/8lall0/GoMondrian"
)

func main() {
	img, _ := gomondrian.Generate(300, 200, 7, 350, 30)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
