package render

import (
	"fmt"
	"os"

	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/utils"
)

// TODO: store width and height as Size2i struct
type Image struct {
	Width  int64
	Height int64
	Pixels []Pixel
}

func NewImage(size maths.Size2i) Image {
	return Image{
		Width:  size.Width,
		Height: size.Height,
		Pixels: make([]Pixel, (size.Width * size.Height))}
}

func (this *Image) AddSample(x int64, y int64, colour Colour) {
	index := x + (y * this.Width)
	this.Pixels[index].AddSample(colour)
}

func (this *Image) AddSamplePixel(x, y int64, pixel Pixel) {
	index := x + (y * this.Width)
	this.Pixels[index].AddSamplePixel(pixel)
}

func (this *Image) AddSampleImage(rhs Image) {
	for y := int64(0); y < this.Height; y++ {
		for x := this.Width - 1; x >= 0; x-- {
			index := x + (y * this.Width)
			this.AddSamplePixel(x, y, rhs.Pixels[index])
		}
	}
}

func (this Image) SaveAsPPM(filename string) {
	filename += ".ppm"

	TryDeletingExistingImage(filename)

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Failed to open file: ", err)
		return
	}
	defer file.Close()

	fmt.Printf("Saving image \"%s\" to disk\n", filename)

	file.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", this.Width, this.Height))
	for y := int64(0); y < this.Height; y++ {
		for x := this.Width - 1; x >= 0; x-- {
			index := x + (y * this.Width)
			colour := this.Pixels[index].GetSampledColour().ToTrueColour()

			file.WriteString(fmt.Sprintf("%s\n", colour.String()))
		}
	}
}

func TryDeletingExistingImage(filename string) {
	if utils.DoesFileExist(filename) {
		fmt.Printf("Deleting existing \"%s\" file\n", filename)

		os.Remove(filename)
	}
}
