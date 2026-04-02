package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	// 画像サイズ
	width := 100
	height := 100

	// 画像作成
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// 全ピクセルを赤にする
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.RGBA{255, 0, 0, 255})
		}
	}

	// ファイル作成
	f, err := os.Create("simple.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// PNGとして書き出し
	png.Encode(f, img)
}
