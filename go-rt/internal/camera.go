package internal

import "github.com/wmbat/ray_tracer/internal/maths"

type Camera struct {
	Origin maths.Point3 
	Viewport maths.Size2
	FocalLength float64
}

func NewCamera() Camera {
	return Camera{}
}

func CalculateLowerLeftCorner(origin *maths.Point3, horizontal *maths.Vec3, vertical *maths.Vec3, focalLength float64) maths.Point3 {
	horizontalHalf := maths.Point3FromVec3(horizontal).Scale(0.5)
	verticalHalf := maths.Point3FromVec3(vertical).Scale(0.5)
	focal := maths.Point3{X: 0, Y: 0, Z: focalLength}

	return origin.Sub(&horizontalHalf).Sub(&verticalHalf).Sub(&focal)
}
