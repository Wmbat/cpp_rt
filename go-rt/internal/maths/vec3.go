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

func (lhs Vec3) Add(rhs Vec3) Vec3 {
	return Vec3{
		X: rhs.X + lhs.X,
		Y: rhs.Y + lhs.Y,
		Z: rhs.Z + lhs.Z}
}

func (lhs Vec3) Sub(rhs Vec3) Vec3 {
	return Vec3{
		X: rhs.X - lhs.X,
		Y: rhs.Y - lhs.Y,
		Z: rhs.Z - lhs.Z}
}

func (lhs Vec3) Mult(rhs Vec3) Vec3 {
	return Vec3{
		X: lhs.X * rhs.X,
		Y: lhs.Y * rhs.Y,
		Z: lhs.Z * rhs.Z}
}

func (lhs Vec3) Scale(factor float64) Vec3 {
	return Vec3{
		X: lhs.X * factor,
		Y: lhs.Y * factor,
		Z: lhs.Z * factor}
}

func (vec Vec3) Normalize() Vec3 {
	return vec.Scale(1 / vec.Length())
}

func (vec Vec3) Negate() Vec3 {
	return vec.Scale(-1.0)
}

func DotProduct(lhs, rhs Vec3) float64 {
	x := lhs.X * rhs.X
	y := lhs.Y * rhs.Y
	z := lhs.Z * rhs.Z

	return x + y + z
}

func CrossProduct(lhs, rhs Vec3) Vec3 {
	x := (lhs.Y * rhs.Z) - (lhs.Z * rhs.Y)
	y := (lhs.Z * rhs.X) - (lhs.X * rhs.Z)
	z := (lhs.X * rhs.Y) - (lhs.Y * rhs.X)

	return Vec3{X: x, Y: y, Z: z}
}
