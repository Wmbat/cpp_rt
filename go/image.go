package main

import (
	"fmt"
	"go_pt/maths"
)

type Image struct {
    Width int
    Height int

    Pixels []maths.Vec3
}

func NewImage(width int, height int) Image {
    return Image{Width: width, Height: height, Pixels: make([]maths.Vec3, width * height)}
}

func (image *Image) SetImagePixelColour(x int, y int, colour *maths.Vec3) {
    image.Pixels[x + y * image.Width] = *colour;
}

func (image *Image) String() string {
    data := fmt.Sprintf("P3\n%d %d\n255\n", image.Width, image.Height)
    for y := image.Height - 1; y >= 0; y-- {
        for x := 0; x < image.Width; x++ {
            colour := image.Pixels[x + y * image.Width]

            data += fmt.Sprintf(
                "%d %d %d\n", 
                int32(colour.X * 255.999), 
                int32(colour.Y * 255.999), 
                int32(colour.Z * 255.999))
        }
    }

    return data
}
