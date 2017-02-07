package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	X, Y, Wide, Heigth int
	Model              color.Model
	StripeSize         int
}

func (i *Image) ColorModel() color.Model {
	return i.Model
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rectangle{
		image.Point{i.X, i.Y},
		image.Point{i.X + i.Wide, i.Y + i.Heigth},
	}
}

func (i *Image) At(x, y int) color.Color {
	var c color.Color
	switch i.ColorModel() {
	case color.RGBAModel:
		c = color.RGBA{
			uint8((x - i.X) / i.StripeSize % 2 * 0xff),
			uint8((x - i.X) / i.StripeSize % 2 * 0xff),
			uint8((x - i.X) / i.StripeSize % 2 * 0xff),
			0xff,
		}
	case color.RGBA64Model:
		c = color.RGBA64{
			uint16((x - i.X) / i.StripeSize % 2 * 0xffff),
			uint16((x - i.X) / i.StripeSize % 2 * 0xffff),
			uint16((x - i.X) / i.StripeSize % 2 * 0xffff),
			0xffff,
		}
	case color.NRGBAModel:
		c = color.NRGBA{
			uint8((x - i.X) / i.StripeSize % 2 * 0xff),
			uint8((x - i.X) / i.StripeSize % 2 * 0xff),
			uint8((x - i.X) / i.StripeSize % 2 * 0xff),
			0xff,
		}
	case color.NRGBA64Model:
		c = color.NRGBA64{
			uint16((x - i.X) / i.StripeSize % 2 * 0xffff),
			uint16((x - i.X) / i.StripeSize % 2 * 0xffff),
			uint16((x - i.X) / i.StripeSize % 2 * 0xffff),
			0xffff,
		}
	case color.AlphaModel:
		c = color.Alpha{
			uint8((x - i.X) / i.StripeSize % 2 * 0xff),
		}
	case color.Alpha16Model:
		c = color.Alpha16{
			uint16((x - i.X) / i.StripeSize % 2 * 0xffff),
		}
	case color.GrayModel:
		c = color.Gray{
			uint8((x - i.X) / i.StripeSize % 2 * 0xff),
		}
	case color.Gray16Model:
		c = color.Gray16{
			uint16((x - i.X) / i.StripeSize % 2 * 0xffff),
		}
	}

	return c
}

func main() {
	p := Image{
		X:          30,
		Y:          10,
		Wide:       100,
		Heigth:     60,
		Model:      color.RGBA64Model,
		StripeSize: 20,
	}
	pic.ShowImage(&p)
}
