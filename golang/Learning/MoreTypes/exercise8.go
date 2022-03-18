//Define your own Image type, implement the necessary methods, and call pic.ShowImage.
//
//Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).
//
//ColorModel should return color.RGBAModel.
//
//At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.
package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w, h int
}

func (i Image) ColorModel() color.Model {
	//TODO implement m
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	//TODO implement me
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	//TODO implement me
	return color.RGBA{uint8(x * x), uint8(y * y), 255, 255}
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(m)
}
