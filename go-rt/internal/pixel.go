package internal

import (
	"fmt"
	"github.com/wmbat/ray_tracer/internal/maths"
)

type Pixel struct {
	Red   float64
	Green float64
	Blue  float64
}

func (this *Pixel) Init(val *maths.Vec3) {
	this.Red = val.X
	this.Green = val.Y
	this.Blue = val.Z
}

func (this Pixel) String() string {
	const factor float64 = 255.999

	red := this.Red * factor
	green := this.Green * factor
	blue := this.Blue * factor

	return fmt.Sprint(red, green, blue)
}
