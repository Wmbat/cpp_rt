package render

import (
	"fmt"

	"github.com/wmbat/ray_tracer/internal/maths"
)

// Represents a raw colour used by the ray tracer
type Colour struct {
	Red, Green, Blue  float64
}

// Convert a 3d vector to a colour
func ColourFromVec3(vec maths.Vec3) Colour {
	return Colour{vec.X, vec.Y, vec.Z}
}

func (this Colour) String() string {
	return fmt.Sprintf("%f %f %f", this.Red, this.Green, this.Blue)
}

// Add one colour to another and returns the resultant colour
func (lhs Colour) Add(rhs Colour) Colour {
	return Colour{
		Red:   lhs.Red + rhs.Red,
		Green: lhs.Green + rhs.Green,
		Blue:  lhs.Blue + rhs.Blue}
}

// Subtract one colour to another and returns the resultant colour
func (lhs Colour) Sub(rhs Colour) Colour {
	return Colour{
		Red:   lhs.Red - rhs.Red,
		Green: lhs.Green - rhs.Green,
		Blue:  lhs.Blue - rhs.Blue}
}

// Multiply one colour to another and returns the resultant colour
func (lhs Colour) Mult(rhs Colour) Colour {
	return Colour{
		Red:   lhs.Red * rhs.Red,
		Green: lhs.Green * rhs.Green,
		Blue:  lhs.Blue * rhs.Blue}
}

// Scale the colour by a scalar factor
func (lhs Colour) Scale(scalar float64) Colour {
	return Colour{
		Red:   lhs.Red * scalar,
		Green: lhs.Green * scalar,
		Blue:  lhs.Blue * scalar}
}

// Convert a raw colour to a true colour.
func (this Colour) ToTrueColour() TrueColour {
	rawColour := this.Scale(256.0)

	return TrueColour{
		Red: uint8(rawColour.Red),
		Green: uint8(rawColour.Green),
		Blue: uint8(rawColour.Blue)}
}
