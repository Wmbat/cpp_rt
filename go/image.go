package main

import (
	"fmt"
	"go_pt/maths"
)

type Image struct {
    width int
    height int

    pixels []Pixel
}

func NewImage(width int, height int) Image {
    return Image{width: width, height: height, pixels: make([]Pixel, width * height)}
}

func (image *Image) AddSample(x int, y int, colour *maths.Vec3, count uint64) {
    image.pixels[x + y * image.width].AddSample(colour, count)
}

func (image *Image) AddSamplePixel(x int, y int, pixel *Pixel) {
    image.pixels[x + y * image.width].AddSamplePixel(pixel)
}

func (image *Image) String() string {
    output := fmt.Sprintf("P3\n%v %v\n255\n", image.width, image.height)

    for x := 0; x < image.width; x++ {
        for y := 0; y < image.width; y++ {
            output += image.pixels[x + y * image.width].String()
            output += "\n"
        }
    }

    return output
}
