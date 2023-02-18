package render

import (
	"github.com/wmbat/ray_tracer/internal/maths"
)

type Pixel struct {
	Colour      Colour
	SampleCount uint64
}

func (this *Pixel) AddSample(colour Colour) {
    this.Colour = this.Colour.Add(colour)
    this.SampleCount += 1
}

func (this *Pixel) AddSamples(colour Colour, sampleCount uint64) {
    this.Colour = this.Colour.Add(colour)
    this.SampleCount += sampleCount
}

func (this *Pixel) AddSamplePixel(pixel Pixel) {
	this.Colour = this.Colour.Add(pixel.Colour)
    this.SampleCount += pixel.SampleCount
}


func (this Pixel) GetSampledColour() Colour {
	if this.SampleCount == 0 {
		return this.Colour
	}

	sampleColour := this.Colour.Scale(1.0 / float64(this.SampleCount))

	return Colour{
		Red: maths.Clamp(sampleColour.Red, 0, 0.999),
		Green: maths.Clamp(sampleColour.Green, 0, 0.999),
		Blue: maths.Clamp(sampleColour.Blue, 0, 0.999)}
}
