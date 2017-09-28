package main

import (
	"golang.org/x/tour/pic"
	"image/color"
	"image"
	"fmt"
)


type Image struct{
	Rect image.Rectangle
}

func (image Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (i Image) Bounds() image.Rectangle {
	return i.Rect
}

func (image Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x*y), uint8(y+x), 255, 255}
}

func main() {
	m := image.NewRGBA(image.Rect(0,0,100,100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0,0).RGBA())

	ex_m := Image{image.Rect(0,0,300,300)}
	pic.ShowImage(ex_m)
}
