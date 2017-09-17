package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

type boolgen struct {
	src       rand.Source
	cache     int64
	remaining int
}

func (b *boolgen) Bool() bool {
	if b.remaining == 0 {
		b.cache, b.remaining = b.src.Int63(), 63
	}

	result := b.cache&0x01 == 1
	b.cache >>= 1
	b.remaining--

	return result
}

func New() *boolgen {
	return &boolgen{src: rand.NewSource(time.Now().UnixNano())}
}

func RandInt(max, min int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func main() {
	r := New()

	padding := 7
	width := 300
	height := 200
	nDiv := 50
	nColor := 30

	red := color.RGBA{255, 0, 0, 255}
	ble := color.RGBA{0, 0, 255, 255}
	ylw := color.RGBA{255, 255, 0, 255}
	blk := color.RGBA{0, 0, 0, 255}
	wht := color.RGBA{255, 255, 255, 255}

	var matrix [300][200]int

	for i := 0; i < width; i++ {
		matrix[i][0] = 1
		matrix[i][height-1] = 1
	}
	for i := 0; i < height; i++ {
		matrix[0][i] = 1
		matrix[width-1][i] = 1
	}

	for i := 0; i < nDiv; {

		rowFlag, colFlag := true, true

		x := RandInt(width, 0)
		y := RandInt(height, 0)

		if x+padding >= width || x-padding <= 0 {
			rowFlag = false
		}
		if y+padding >= height || y-padding <= 0 {
			colFlag = false
		}

		if rowFlag {
			for j := 1; j <= padding; j++ {
				if matrix[x+j][y] == 1 || matrix[x-j][y] == 1 {
					rowFlag = false
				}
			}
		}
		if colFlag {
			for j := 1; j <= padding; j++ {
				if matrix[x][y+j] == 1 || matrix[x][y-j] == 1 {
					colFlag = false
				}
			}
		}

		if rowFlag && colFlag {
			rowFlag = r.Bool()
			colFlag = !rowFlag
		}

		if colFlag {
			matrix[x][y] = 1

			for j, intFlag := x-1, true; j > 0 && intFlag; j-- {
				if matrix[j][y] == 1 && j > 0 {
					for k := 1; k <= padding && intFlag; k++ {
						if matrix[j-1][y+k] == 1 || matrix[j-1][y-k] == 1 {
							intFlag = false
						}
					}
					if intFlag {
						intFlag = r.Bool()
					}
				}
				matrix[j][y] = 1
			}
			for j, intFlag := x+1, true; j < width && intFlag; j++ {
				if matrix[j][y] == 1 && j < width-1 {
					for k := 1; k <= padding && intFlag; k++ {
						if matrix[j+1][y+k] == 1 || matrix[j+1][y-k] == 1 {
							intFlag = false
						}
					}
					if intFlag {
						intFlag = r.Bool()
					}
				}
				matrix[j][y] = 1
			}
			i++
		}

		if rowFlag {
			matrix[x][y] = 1

			for j, intFlag := y-1, true; j > 0 && intFlag; j-- {
				if matrix[x][j] == 1 && j > 0 {
					for k := 1; k <= padding && intFlag; k++ {
						if matrix[x+k][j+1] == 1 || matrix[x-k][j+1] == 1 {
							intFlag = false
						}
					}
					if intFlag {
						intFlag = r.Bool()
					}
				}
				matrix[x][j] = 1
			}
			for j, intFlag := y+1, true; j < height && intFlag; j++ {
				if matrix[x][j] == 1 && j < height-1 {
					for k := 1; k <= padding && intFlag; k++ {
						if matrix[x+k][j+1] == 1 || matrix[x-k][j+1] == 1 {
							intFlag = false
						}
					}
					if intFlag {
						intFlag = r.Bool()
					}
				}
				matrix[x][j] = 1
			}
			i++
		}
	}

	for i := 0; i < nColor; {
		var j int
		x := RandInt(width, 0)
		y := RandInt(height, 0)

		if matrix[x][y] == 0 {
			color := RandInt(5, 2)

			for j = x; matrix[j][y] == 0; j++ {
			}
			xEnd := j - 1
			for j = x; matrix[j][y] == 0; j-- {
			}
			xStart := j + 1
			for j = y; matrix[x][j] == 0; j++ {
			}
			yEnd := j - 1
			for j = y; matrix[x][j] == 0; j-- {
			}
			yStart := j + 1
			for j := xStart; j <= xEnd; j++ {
				for k := yStart; k <= yEnd; k++ {
					matrix[j][k] = color
				}
			}
			i++
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if matrix[i][j] == 1 {
				img.Set(i, j, blk)
			} else if matrix[i][j] == 2 {
				img.Set(i, j, red)
			} else if matrix[i][j] == 3 {
				img.Set(i, j, ylw)
			} else if matrix[i][j] == 4 {
				img.Set(i, j, ble)
			} else {
				img.Set(i, j, wht)
			}
		}
	}

	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
