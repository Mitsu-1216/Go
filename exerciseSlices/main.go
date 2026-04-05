package main

import "golang.org/x/tour/pic"

func Pic(l, c int) [][]uint8 {

	img := make([][]uint8, c)

	for y := 0; y < c; y++ {
		img[y] = make([]uint8, l)

		for x := 0; x < l; l++ {
			img[y][x] = uint8(x * y)
		}

	}

	return img
}

func main() {
	pic.Show(Pic)
}
