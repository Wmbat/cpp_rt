package maths

import "fmt"

type Point3 struct {
	X float64
	Y float64
	Z float64
}

func (this Point3) Clone() Point3 {
	return Point3{this.X, this.Y, this.Z}
}

func (this Point3) String() string {
	return fmt.Sprintf("%f %f %f", this.X, this.Y, this.Z)
}

func (this Point3) ToVec3() Vec3 {
	return Vec3{this.X, this.Y, this.Z}
}

func (this Point3) Add(val Point3) Point3 {
	return Point3{
		X: this.X + val.X,
		Y: this.Y + val.Y,
		Z: this.Z + val.Z}
}

func (this Point3) Sub(val Point3) Point3 {
	return Point3{
		X: this.X - val.X,
		Y: this.Y - val.Y,
		Z: this.Z - val.Z}
}

func (this Point3) Mult(val Point3) Point3 {
	return Point3{
		X: this.X * val.X,
		Y: this.Y * val.Y,
		Z: this.Z * val.Z}
}

func (this Point3) Scale(factor float64) Point3 {
	return Point3{
		X: this.X * factor,
		Y: this.Y * factor,
		Z: this.Z * factor}
}

func Point3FromVec3(value Vec3) Point3 {
	return Point3{X: value.X, Y: value.Y, Z: value.Z}
}
