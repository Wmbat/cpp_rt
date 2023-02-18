package maths

import "fmt"

type Point2 struct {
	X float64
	Y float64
}

func (this Point2) Clone() Point2 {
	return Point2{this.X, this.Y}
}

func (this Point2) String() string {
	return fmt.Sprintf("%f %f", this.X, this.Y)
}

func (this Point2) Add(val Point2) Point2 {
	return Point2{
		X: this.X + val.X,
		Y: this.Y + val.Y}
}

func (this Point2) Sub(val Point2) Point2 {
	return Point2{
		X: this.X - val.X,
		Y: this.Y - val.Y}
}

func (this Point2) Mult(val Point2) Point2 {
	return Point2{
		X: this.X * val.X,
		Y: this.Y * val.Y}
}

func (this Point2) Scale(factor float64) Point2 {
	return Point2{
		X: this.X * factor,
		Y: this.Y * factor}
}
