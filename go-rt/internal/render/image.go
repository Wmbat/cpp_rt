package render

import (
	"fmt"
	"log"
	"os"

	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/utils"
)

type Image struct {
	Size   maths.Size2[int]
	pixels []Pixel
}

// Create a new image
func NewImage(size maths.Size2[int]) Image {
	return Image{
		Size:   size,
		pixels: make([]Pixel, (size.Width * size.Height))}
}

func (this *Image) AddSample(x, y int, colour Colour) {
	index := x + (y * this.Size.Width)
	this.pixels[index].AddSample(colour)
}

func (this *Image) AddSamplePixel(x, y int, pixel Pixel) {
	index := x + (y * this.Size.Width)
	this.pixels[index].AddSamplePixel(pixel)
}

func (this *Image) AddSampleImage(rhs Image) {
	for j := this.Size.Height - 1; j >= 0; j-- {
		for i := 0; i < this.Size.Width; i++ {
			index := i + (j * this.Size.Width)
			this.AddSamplePixel(i, j, rhs.pixels[index])
		}
	}
}

// Convert and save the image as a PPM file
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

	file.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", this.Size.Width, this.Size.Height))

	for j := this.Size.Height - 1; j >= 0; j-- {
		for i := 0; i < this.Size.Width; i++ {
			index := i + (j * this.Size.Width)
			colour := this.pixels[index].GetSampledColour().ToTrueColour()

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
