package core

import (
	"go_rt/maths"
	"math"
)

type Pixel struct {
    colour Colour
    sampleCount uint64
}

func (this *Pixel) AddSamplePixel(pixel *Pixel) {
    this.AddSamples(&pixel.colour, pixel.sampleCount)
}

func (this *Pixel) AddSamples(colour *Colour, sampleCount uint64) {
    this.colour.Add(colour)
    this.sampleCount += sampleCount
}

func (this Pixel) ComputeColour() Colour {
    if this.sampleCount == 0 {
        return this.colour
    }

    scale := 1.0 / float64(this.sampleCount)
    
    return Colour{
        R: maths.Clamp(math.Sqrt(scale * this.colour.R), 0.0, 0.999),
        G: maths.Clamp(math.Sqrt(scale * this.colour.G), 0.0, 0.999),
        B: maths.Clamp(math.Sqrt(scale * this.colour.B), 0.0, 0.999)}
}
