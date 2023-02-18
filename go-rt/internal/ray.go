package rt

import "github.com/wmbat/ray_tracer/internal/maths"

type Ray struct {
	Origin maths.Point3
	Direction maths.Vec3
}

func (this Ray) At(time float64) maths.Point3 {
	return this.Origin.Add(this.Direction.Scale(time).ToPoint3())
}
