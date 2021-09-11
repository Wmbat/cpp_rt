package core

import (
	"fmt"
)

type Image struct {
    Width int
    Height int

    Pixels []Pixel
}

func NewImage(width int, height int) Image {
    return Image{Width: width, Height: height, Pixels: make([]Pixel, width * height)}
}

func (this *Image) AddSampleImage(image *Image) {
    // Should really check that the 2 images are the same size
    for x := 0; x < this.Width; x++ {
        for y := 0; y < this.Height; y++ {
            this.AddSamplePixel(x, y, &image.Pixels[x + y * this.Width])
        }
    }
}

func (image *Image) AddSamplePixel(x int, y int, pixel *Pixel) {
    image.AddSamples(x, y, &pixel.colour, pixel.sampleCount) 
}

func (image *Image) AddSamples(x int, y int, colour *Colour, sampleCount uint64) {
    image.Pixels[x + y * image.Width].AddSamples(colour, sampleCount)
}

func (image *Image) String() string {
    data := fmt.Sprintf("P3\n%d %d\n255\n", image.Width, image.Height)
    for y := image.Height - 1; y >= 0; y-- {
        for x := 0; x < image.Width; x++ {
            colour := image.Pixels[x + y * image.Width].ComputeColour()

            data += fmt.Sprintf(
                "%d %d %d\n", 
                int32(colour.R * 255.999), 
                int32(colour.G * 255.999), 
                int32(colour.B * 255.999))
        }
    }

    return data
}
