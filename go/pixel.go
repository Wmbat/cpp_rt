package go_pt

import "go_pt/math"

type Pixel struct {
   colour math.Vec3
   samples_count uint64
}

func (p* Pixel) AddSample(colour *math.Vec3, count uint64)  {
    p.colour.Add(colour);
    p.samples_count += count; 
}

func (p *Pixel) AddSamplePixel(sample *Pixel) {
    p.AddSample(&sample.colour, sample.samples_count)
}

func ComputePixelColour(pixel *Pixel) math.Vec3 {
    if pixel.samples_count == 0 {
        return pixel.colour
    } else {
        return math.MultScalar(&pixel.colour, 1.0 / float32(pixel.samples_count))
    }
}
