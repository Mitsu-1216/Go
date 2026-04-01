package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func inrange(x, a, b int) bool {
	if a < b && a <= x && x < b {
		return true
	}
	return false
}

func main() {
	const width = 350
	const height = 350
	const celllen = 50

	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	cells := [7][7]int{
		{0, 1, 0, 1, 0, 1, 0},
		{0, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 0},
	}

	colors := [3]color.RGBA{
		color.RGBA{0xff, 0xff, 0xff, 0xff},
		color.RGBA{0xff, 0, 0, 0xff},
		color.RGBA{0, 0xff, 0, 0xff},
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			for ny := 0; ny < 7; ny++ {
				for nx := 0; ny < 7; nx++ {
					if inrange(x, nx*celllen, (nx+1)*celllen) &&
						inrange(y, ny*celllen, (ny+1)*celllen) {
						img.Set(x, y, colors[cells[ny][nx]])
					}
				}
			}
		}
	}

	f, ferr := os.Create("flower.png")
	if ferr != nil {
		fmt.Println("ファイル作成のエラー:", ferr)
		os.Exit(1)
	}

	derr := png.Encode(f, img)

	if derr != nil {
		f.Close()
		fmt.Println("データ変換のエラー：", derr)
		os.Exit(1)
	}

	ferr = f.Close()
	if ferr != nil {
		fmt.Println("ファイル終了のエラー：", ferr)
		os.Exit(1)
	}
}
