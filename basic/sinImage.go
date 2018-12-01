package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	const size = 300
	// 创建制定大小的灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))
	// 遍历每个像素并着色为白色
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{255})
		}
	}
	// 画正弦图
	for x := 0; x < size; x++ {
		s := float64(x) * 2 * math.Pi / size
		y := size/2 - math.Sin(s)*size/2
		pic.SetGray(x, int(y), color.Gray{0})
	}
	// 保存文件
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}
	_ = png.Encode(file, pic)
	_ = file.Close()
}
