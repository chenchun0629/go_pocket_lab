package main

import (
	"fmt"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"image"
)

func main() {
	img, err := imgio.Open("./asset/input.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	img = transform.Crop(img, image.Rect(30,30,130,130))
	img = transform.Translate(img, 50, 50)

	if err := imgio.Save("output.png", img, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}
}
