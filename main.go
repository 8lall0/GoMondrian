package main

import (
	"fmt"
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

func CheckCol(m [][]int, padding, x, y int) bool {
	var flag bool
	r := New()

	for k, flag := 1, true; k <= padding && flag; k++ {
		if m[x][y+k] == 1 || m[x][y-k] == 1 {
			flag = false
		}
	}
	if flag {
		return r.Bool()
	}

	return flag
}

func CheckRow(m [][]int, padding, x, y int) bool {
	var flag bool
	r := New()

	for k := 1; k <= padding && flag; k++ {
		if m[x+k][y] == 1 || m[x-k][y] == 1 {
			flag = false
		}
	}
	if flag {
		return r.Bool()
	}

	return flag
}

func main() {
	var colors = []color.Color{
		color.White,
		color.Black,
		color.RGBA{255, 0, 0, 255},
		color.RGBA{0, 0, 255, 255},
		color.RGBA{255, 255, 0, 255},
	}

	r := New()

	padding := 7
	width := 300
	height := 200
	nDiv := 350
	nColor := 30

	// Bogus placeholder to define if possible to proceed
	if (width-2)/(padding+1)*(height-2)/(padding+1) < nDiv {
		fmt.Println("Not enough")
		os.Exit(1)
	}

	m := [][]int{}
	for i := 0; i < width; i++ {
		r := make([]int, height)
		m = append(m, r)
	}

	for i := 0; i < width; i++ {
		m[i][0] = 1
		m[i][height-1] = 1
	}
	for i := 0; i < height; i++ {
		m[0][i] = 1
		m[width-1][i] = 1
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
				if m[x+j][y] == 1 || m[x-j][y] == 1 {
					rowFlag = false
				}
			}
		}
		if colFlag {
			for j := 1; j <= padding; j++ {
				if m[x][y+j] == 1 || m[x][y-j] == 1 {
					colFlag = false
				}
			}
		}

		if rowFlag && colFlag {
			rowFlag = r.Bool()
			colFlag = !rowFlag
		}

		if colFlag {
			m[x][y] = 1

			for j, intFlag := x-1, true; j > 0 && intFlag; j-- {
				if m[j][y] == 1 && j > 0 {
					intFlag = CheckCol(m, padding, j-1, y)
				}
				m[j][y] = 1
			}
			for j, intFlag := x+1, true; j < width && intFlag; j++ {
				if m[j][y] == 1 && j < width-1 {
					intFlag = CheckCol(m, padding, j+1, y)
				}
				m[j][y] = 1
			}
			i++
		}

		if rowFlag {
			m[x][y] = 1

			for j, intFlag := y-1, true; j > 0 && intFlag; j-- {
				if m[x][j] == 1 && j > 0 {
					intFlag = CheckRow(m, padding, x, j-1)
				}
				m[x][j] = 1
			}
			for j, intFlag := y+1, true; j < height && intFlag; j++ {
				if m[x][j] == 1 && j < height-1 {
					intFlag = CheckRow(m, padding, x, j+1)
				}
				m[x][j] = 1
			}
			i++
		}
	}

	for i := 0; i < nColor; {
		var j int
		x := RandInt(width, 0)
		y := RandInt(height, 0)

		if m[x][y] == 0 {
			color := RandInt(5, 2)

			for j = x; m[j][y] == 0; j++ {
			}
			xEnd := j - 1
			for j = x; m[j][y] == 0; j-- {
			}
			xStart := j + 1
			for j = y; m[x][j] == 0; j++ {
			}
			yEnd := j - 1
			for j = y; m[x][j] == 0; j-- {
			}
			yStart := j + 1
			for j := xStart; j <= xEnd; j++ {
				for k := yStart; k <= yEnd; k++ {
					m[j][k] = color
				}
			}
			i++
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width-2, height-2))
	for i := 1; i < width-1; i++ {
		for j := 1; j < height-1; j++ {
			img.Set(i-1, j-1, colors[m[i][j]])
		}
	}

	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
