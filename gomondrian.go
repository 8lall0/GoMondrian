package gomondrian

import (
	"errors"
	"image"
	"image/color"
)

// The standard colors for a Mondrian image are red, yellow and black.
var colors = []color.Color{
	color.White,
	color.Black,
	color.RGBA{255, 0, 0, 255},
	color.RGBA{0, 0, 255, 255},
	color.RGBA{255, 255, 0, 255},
}

func Generate(width, height, padding, nDiv, nColor int) (Image, error) {
	// Bogus placeholder to define if possible to proceed
	if (width-2)/(padding+1)*(height-2)/(padding+1) < nDiv {
		return nil, errors.New("Too much divisions for your width|height.")
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

	r := randGen()

	for i := 0; i < nDiv; {
		rowFlag, colFlag := true, true

		x := randInt(width, 0)
		y := randInt(height, 0)

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
			rowFlag = r.randBool()
			colFlag = !rowFlag
		}

		if colFlag {
			m[x][y] = 1

			for j, intFlag := x-1, true; j > 0 && intFlag; j-- {
				if m[j][y] == 1 && j > 0 {
					intFlag = checkCol(m, padding, j-1, y)
					if intFlag {
						return r.randBool()
					}
				}
				m[j][y] = 1
			}
			for j, intFlag := x+1, true; j < width && intFlag; j++ {
				if m[j][y] == 1 && j < width-1 {
					intFlag = checkCol(m, padding, j+1, y)
					if intFlag {
						return r.randBool()
					}
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
					if intFlag {
						return r.randBool()
					}
				}
				m[x][j] = 1
			}
			for j, intFlag := y+1, true; j < height && intFlag; j++ {
				if m[x][j] == 1 && j < height-1 {
					intFlag = CheckRow(m, padding, x, j+1)
					if intFlag {
						return r.randBool()
					}
				}
				m[x][j] = 1
			}
			i++
		}
	}

	for i, j := 0, 0; i < nColor; i++ {
		x := randInt(width, 0)
		y := randInt(height, 0)

		if m[x][y] == 0 {
			for j = x; m[j][y] == 0; j-- {
			}
			xStart := j + 1
			for j = x; m[j][y] == 0; j++ {
			}
			xEnd := j - 1
			for j = y; m[x][j] == 0; j-- {
			}
			yStart := j + 1
			for j = y; m[x][j] == 0; j++ {
			}
			yEnd := j - 1

			color := randInt(5, 2)
			for j := xStart; j <= xEnd; j++ {
				for k := yStart; k <= yEnd; k++ {
					m[j][k] = color;
				}
			}
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width-2, height-2))
	for i := 1; i < width-1; i++ {
		for j := 1; j < height-1; j++ {
			img.Set(i-1, j-1, colors[m[i][j]])
		}
	}

	return img
}

func checkCol(m [][]int, padding, x, y int) flag bool {
	for k, flag := 1, true; k <= padding && flag; k++ {
		if m[x][y+k] == 1 || m[x][y-k] == 1 {
			flag = false
		}
	}

	return
}

func checkRow(m [][]int, padding, x, y int) flag bool {
	for k, flag := 1, true; k <= padding && flag; k++ {
		if m[x+k][y] == 1 || m[x-k][y] == 1 {
			flag = false
		}
	}

	return
}
