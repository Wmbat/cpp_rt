package internal

import "github.com/wmbat/ray_tracer/internal/maths"

type Ray struct {
	Origin    maths.Point3
	Direction maths.Vec3
}

func (this Ray) PointAlong(time float64) maths.Point3 {
	return this.Direction.Scale(time).ToPoint3().Add(&this.Origin)
}
