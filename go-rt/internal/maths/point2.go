package maths

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Point2[Type constraints.Integer | constraints.Float] struct {
	X, Y Type
}

func (this Point2[Type]) Clone() Point2[Type] {
	return Point2[Type]{this.X, this.Y}
}

func (this Point2[Type]) String() string {
	return fmt.Sprintf("<%v, %v>", this.X, this.Y)
}

func (this Point2[Type]) Add(val Point2[Type]) Point2[Type] {
	return Point2[Type]{
		X: this.X + val.X,
		Y: this.Y + val.Y}
}

func (this Point2[Type]) Sub(val Point2[Type]) Point2[Type] {
	return Point2[Type]{
		X: this.X - val.X,
		Y: this.Y - val.Y}
}

func (this Point2[Type]) Mult(val Point2[Type]) Point2[Type] {
	return Point2[Type]{
		X: this.X * val.X,
		Y: this.Y * val.Y}
}

func (this Point2[Type]) Scale(factor Type) Point2[Type] {
	return Point2[Type]{
		X: this.X * factor,
		Y: this.Y * factor}
}
