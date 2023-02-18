package rt

import (
	"fmt"

	"github.com/wmbat/ray_tracer/internal/maths"
)

type Colour struct {
	Red   float64
	Green float64
	Blue  float64
}

func ColourFromVec3(vec maths.Vec3) Colour {
	return Colour{vec.X, vec.Y, vec.Z}
}

func (this Colour) Clone() Colour {
	return Colour{this.Red, this.Green, this.Blue}
}

func (this Colour) String() string {
	return fmt.Sprintf("%f %f %f", this.Red, this.Green, this.Blue)
}

func (lhs Colour) Add(rhs Colour) Colour {
	return Colour{
		Red:   lhs.Red + rhs.Red,
		Green: lhs.Green + rhs.Green,
		Blue:  lhs.Blue + rhs.Blue}
}

func (lhs Colour) Sub(rhs Colour) Colour {
	return Colour{
		Red:   lhs.Red - rhs.Red,
		Green: lhs.Green - rhs.Green,
		Blue:  lhs.Blue - rhs.Blue}
}

func (lhs Colour) Mult(rhs Colour) Colour {
	return Colour{
		Red:   lhs.Red * rhs.Red,
		Green: lhs.Green * rhs.Green,
		Blue:  lhs.Blue * rhs.Blue}
}

func (lhs Colour) Scale(scalar float64) Colour {
	return Colour{
		Red:   lhs.Red * scalar,
		Green: lhs.Green * scalar,
		Blue:  lhs.Blue * scalar}
}

func (this Colour) ToTrueColour() TrueColour {
	rawColour := this.Scale(256.0)

	return TrueColour{
		Red: uint8(rawColour.Red),
		Green: uint8(rawColour.Green),
		Blue: uint8(rawColour.Blue)}
}
