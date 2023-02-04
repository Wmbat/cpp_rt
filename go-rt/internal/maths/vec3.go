package maths

import (
	"fmt"
	"math"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (this Vec3) Clone() Vec3 {
	return Vec3{this.X, this.Y, this.Z}
}

func (this Vec3) String() string {
	return fmt.Sprintf("%f %f %f", this.X, this.Y, this.Z)
}

func (this Vec3) ToPoint3() Point3 {
	return Point3{this.X, this.Y, this.Z}
}

func (this Vec3) Length() float64 {
	return math.Sqrt(this.LengthSquared())
}

func (this Vec3) LengthSquared() float64 {
	x := this.X * this.X
	y := this.Y * this.Y
	z := this.Z * this.Z

	return x + y + z
}

func (this Vec3) Add(val *Vec3) Vec3 {
	return Vec3{
		X: this.X + val.X,
		Y: this.Y + val.Y,
		Z: this.Z + val.Z}
}

func (this Vec3) Sub(val *Vec3) Vec3 {
	return Vec3{
		X: this.X - val.X,
		Y: this.Y - val.Y,
		Z: this.Z - val.Z}
}

func (this Vec3) Mult(val *Vec3) Vec3 {
	return Vec3{
		X: this.X * val.X,
		Y: this.Y * val.Y,
		Z: this.Z * val.Z}
}

func (this Vec3) Scale(factor float64) Vec3 {
	return Vec3{
		X: this.X * factor,
		Y: this.Y * factor,
		Z: this.Z * factor}
}

func DotProduct(lhs *Vec3, rhs *Vec3) float64 {
	x := lhs.X * rhs.X
	y := lhs.Y * rhs.Y
	z := lhs.Z * rhs.Z

	return x + y + z
}

func CrossProduct(lhs *Vec3, rhs *Vec3) Vec3 {
	x := lhs.Y*rhs.Z - lhs.Z*rhs.Y
	y := lhs.Z*rhs.X - lhs.X*rhs.Z
	z := lhs.X*rhs.Y - lhs.Y*rhs.X

	return Vec3{X: x, Y: y, Z: z}
}

func UnitVector(value *Vec3) Vec3 {
	return value.Scale(1 / value.Length())
}
