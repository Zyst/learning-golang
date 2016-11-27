package main

import (
	"image"
	"image/color"
)

type Image struct {
	width  int
	height int
}

// We return an established model
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Return a rectangle with our own width and height
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

// We just create a new RGBA and return that
func (img Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := Image{}

	// I think this is a playground-only method
	pic.ShowImage(m)
}
