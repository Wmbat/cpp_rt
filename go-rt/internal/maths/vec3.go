package maths

import (
	"fmt"
	"math"
	"math/rand"
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

func RandVec3() Vec3 {
	return Vec3{
		X: rand.Float64(), 
		Y: rand.Float64(), 
		Z: rand.Float64()}
}

func RandVec3InRange(min, max float64) Vec3 {

	return Vec3{
		X: min + rand.Float64() * (max - min), 
		Y: min + rand.Float64() * (max - min), 
		Z: min + rand.Float64() * (max - min)}
}


func RandVec3InUnitSphere() Vec3 {
	for {
		vec := RandVec3InRange(-1, 1)
		if vec.LengthSquared() >= 1 {
			continue;
		}
		return vec;
	}
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
