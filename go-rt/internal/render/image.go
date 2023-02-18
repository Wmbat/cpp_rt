package render

import (
	"fmt"
	"os"
)

type Image struct {
	Width  int64
	Height int64
	Pixels []Pixel
}

func NewImage(width int64, height int64) Image {
	return Image{
		Width:  width,
		Height: height,
		Pixels: make([]Pixel, (width * height))}
}

func (this Image) WritePixel(x int64, y int64, pixel Pixel) {
	index := x + (y * this.Width)
	this.Pixels[index].AddSamplePixel(pixel)
}

func (this Image) SaveAsPPM(file *os.File) {
	file.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", this.Width, this.Height))
	for y := int64(0); y < this.Height; y++ {
		for x := this.Width - 1; x >= 0; x-- {
			index := x + (y * this.Width)
			colour := this.Pixels[index].GetSampledColour().ToTrueColour()

			file.WriteString(fmt.Sprintf("%s\n", colour.String()))
		}
	}
}
