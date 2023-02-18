package render

import (
	"github.com/wmbat/ray_tracer/internal/maths"
)

type Pixel struct {
	colour      Colour
	sampleCount uint64
}

func (this *Pixel) AddSample(colour Colour) {
    this.colour = this.colour.Add(colour)
    this.sampleCount += 1
}

func (this *Pixel) AddSamples(colour Colour, sampleCount uint64) {
    this.colour = this.colour.Add(colour)
    this.sampleCount += sampleCount
}

func (this *Pixel) AddSamplePixel(pixel Pixel) {
	this.colour = this.colour.Add(pixel.colour)
    this.sampleCount += pixel.sampleCount
}


func (this Pixel) GetSampledColour() Colour {
	if this.sampleCount == 0 {
		return this.colour
	}

	sampleColour := this.colour.Scale(1.0 / float64(this.sampleCount))

	return Colour{
		Red: maths.Clamp(sampleColour.Red, 0, 0.999),
		Green: maths.Clamp(sampleColour.Green, 0, 0.999),
		Blue: maths.Clamp(sampleColour.Blue, 0, 0.999)}
}
