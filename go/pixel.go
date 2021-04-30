package main

import (
	"fmt"
	"go_pt/maths"
	"math"
)

type Pixel struct {
   colour maths.Vec3
   samples_count uint64
}

func (p* Pixel) AddSample(colour *maths.Vec3, count uint64)  {
    p.colour.Add(colour);
    p.samples_count += count; 
}

func (p *Pixel) AddSamplePixel(sample *Pixel) {
    p.AddSample(&sample.colour, sample.samples_count)
}

func (pixel *Pixel) ComputePixelColour() maths.Vec3 {
    if pixel.samples_count == 0 {
        return pixel.colour
    } else {
        return maths.MultScalar(&pixel.colour, 1.0 / float64(pixel.samples_count))
    }
}

func (p *Pixel) String() string {
    colour := p.ComputePixelColour()

    clamped_x := math.Min(math.Max(colour.X, 0.0), 1.0) * 255.999 
    clamped_y := math.Min(math.Max(colour.Y, 0.0), 1.0) * 255.999
    clamped_z := math.Min(math.Max(colour.Z, 0.0), 1.0) * 255.999

    return fmt.Sprintf("%v %v %v", clamped_x, clamped_y, clamped_z)
}
