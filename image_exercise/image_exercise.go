package image_exercise

import (
	"image"
	"image/color"
)

type Image struct {
	width, height int
}

func (i Image) ShowImage() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) At(x, y int) color.Color {
	v := uint8((x ^ y) % 256)
	return color.RGBA{v, v, 255, 255}
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}
