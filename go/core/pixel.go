package core

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

    result := this.colour
    result.MultScalar(1.0 / float64(this.sampleCount))

    return result
}
