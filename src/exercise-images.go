package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (Image)ColorModel() color.Model {
	return color.RGBAModel
}

func (Image)Bounds() image.Rectangle {
	return image.Rect(0, 0, 511, 511)
}

func (img Image)At(x, y int) color.Color {
	x %= img.Bounds().Max.X
	y %= img.Bounds().Max.Y
	v := uint8((x+y))
	return color.RGBA{0, v, 0, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}

