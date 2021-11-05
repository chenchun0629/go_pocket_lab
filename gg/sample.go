package main

import (
	"github.com/fogleman/gg"
)

func main() {
	// 加载图片，这里路径换成自己的
	im, err := gg.LoadImage("./asset/input.png")
	if err != nil {
		panic(err)
	}

	// 获取图片的宽度和高度
	w := 200
	h := 200

	dc := gg.NewContext(h, w)


	// 取宽度和高度的最小值作为直径
	var radius float64 = 100
	// 画圆形
	dc.DrawCircle(float64(100), float64(100), radius)
	// 对画布进行裁剪
	dc.Clip()
	// 加载图片
	dc.DrawImage(im, -150, -50)
	dc.SavePNG("output.png")
}