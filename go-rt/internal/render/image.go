package render

import (
	"fmt"
	"log"
	"os"

	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/utils"
)

// TODO: store width and height as Size2i struct
type Image struct {
	Width  int
	Height int
	Pixels []Pixel
}

func NewImage(size maths.Size2[int]) Image {
	return Image{
		Width:  size.Width,
		Height: size.Height,
		Pixels: make([]Pixel, (size.Width * size.Height))}
}

func (this *Image) AddSample(x, y int, colour Colour) {
	index := x + (y * this.Width)
	this.Pixels[index].AddSample(colour)
}

func (this *Image) AddSamplePixel(x, y int, pixel Pixel) {
	index := x + (y * this.Width)
	this.Pixels[index].AddSamplePixel(pixel)
}

func (this *Image) AddSampleImage(rhs Image) {
	for j := this.Height - 1; j >= 0; j-- {
		for i := 0; i < this.Width; i++ {
			index := i + (j * this.Width)
			this.AddSamplePixel(i, j, rhs.Pixels[index])
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

	log.Printf("[main] Saving image \"%s\" to disk\n", filename)

	file.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", this.Width, this.Height))

	for j := this.Height - 1; j >= 0; j-- {
		for i := 0; i < this.Width; i++ {
			index := i + (j * this.Width)
			colour := this.Pixels[index].GetSampledColour().ToTrueColour()

			file.WriteString(fmt.Sprintf("%s\n", colour.String()))
		}
	}
}

func TryDeletingExistingImage(filename string) {
	if utils.DoesFileExist(filename) {
		log.Printf("[main] Deleting existing \"%s\" file\n", filename)

		os.Remove(filename)
	}
}
